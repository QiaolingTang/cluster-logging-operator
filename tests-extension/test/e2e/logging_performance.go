package logging

import (
	"context"
	"fmt"
	"math"
	"path/filepath"
	"strconv"
	"time"

	"github.com/openshift/cluster-logging-operator-tests-extension/test/e2e/testdata"

	g "github.com/onsi/ginkgo/v2"
	o "github.com/onsi/gomega"
	"github.com/openshift/origin/test/extended/util/compat_otp"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	e2e "k8s.io/kubernetes/test/e2e/framework"
)

var _ = g.Describe("[OTP][sig-openshift-logging] Logging NonPreRelease - LokiStack Performance Test", func() {
	defer g.GinkgoRecover()

	var (
		oc                 = compat_otp.NewCLI("lokistack-performace-test", compat_otp.KubeConfigPath())
		loggingBaseDir, sc string
		nodes              *corev1.NodeList
		workerNodeCount    int
	)

	g.BeforeEach(func() {

		if !validateInfraAndResourcesForLoki(oc, "38Gi", "18") {
			g.Skip("Current platform not supported/resources not available for this test!")
		}

		// Check worker nodes count
		var err error
		nodes, err = oc.AdminKubeClient().CoreV1().Nodes().List(context.Background(), metav1.ListOptions{
			LabelSelector: "node-role.kubernetes.io/worker=,kubernetes.io/os=linux",
		})
		o.Expect(err).NotTo(o.HaveOccurred())
		workerNodeCount = len(nodes.Items)
		if workerNodeCount == 0 {
			g.Skip("Skipping test: No worker nodes available in the cluster")
		}

		sc, _ = getStorageClassName(oc)
		if len(sc) == 0 {
			g.Skip("The cluster doesn't have a storage class for this test!")
		}

		loggingBaseDir = testdata.FixturePath("logging")
		subTemplate := filepath.Join(loggingBaseDir, "subscription", "sub-template.yaml")
		CLO := SubscriptionObjects{
			OperatorName:  "cluster-logging-operator",
			Namespace:     cloNS,
			PackageName:   "cluster-logging",
			Subscription:  subTemplate,
			OperatorGroup: filepath.Join(loggingBaseDir, "subscription", "allnamespace-og.yaml"),
		}
		LO := SubscriptionObjects{
			OperatorName:  "loki-operator-controller-manager",
			Namespace:     loNS,
			PackageName:   "loki-operator",
			Subscription:  subTemplate,
			OperatorGroup: filepath.Join(loggingBaseDir, "subscription", "allnamespace-og.yaml"),
		}
		compat_otp.By("deploy CLO and LO")
		CLO.SubscribeOperator(oc)
		LO.SubscribeOperator(oc)
	})

	g.It("Author:kbharti-Longduration-CPaasrunOnly-High-84898-Performance-Vector-LokiStack-Performance test from vector to LokiStack using 1x.extra-small tsize and ViaQ datamodel with log loss measurement[Serial][Slow]", func() {
		// Performance test to measure throughput and log loss from vector to LokiStack with 1x.extra-small configuration
		lokiStackTemplate := filepath.Join(loggingBaseDir, "lokistack", "lokistack-simple.yaml")

		ls := lokiStack{
			name:          "lokistack-performance-84898",
			namespace:     loggingNS,
			tSize:         "1x.extra-small",
			storageType:   getStorageType(oc),
			storageSecret: "storage-secret-performance-84898",
			storageClass:  sc,
			bucketName:    "logging-loki-performance-" + getRandomString(),
			template:      lokiStackTemplate,
		}

		defer ls.removeObjectStorage(oc)
		err := ls.prepareResourcesForLokiStack(oc)
		o.Expect(err).NotTo(o.HaveOccurred())
		defer ls.removeLokiStack(oc)
		err = ls.deployLokiStack(oc)
		o.Expect(err).NotTo(o.HaveOccurred())
		ls.waitForLokiStackToBeReady(oc)

		compat_otp.By("Patch LokiStack with custom ingestion limits for performance testing")
		patchData := `{
				"spec": {
					"limits": {
						"global": {
							"ingestion": {
								"ingestionBurstSize": 50,
								"ingestionRate": 16,
								"maxGlobalStreamsPerTenant": 5000
							}
						}
					}
				}
			}`
		err = oc.AsAdmin().WithoutNamespace().Run("patch").Args("lokistack", ls.name, "-n", ls.namespace,
			"-p", patchData, "--type=merge").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		e2e.Logf("LokiStack patched with ingestionRate: 16, ingestionBurstSize: 50, maxGlobalStreamsPerTenant: 5000")
		ls.waitForLokiStackToBeReady(oc)

		compat_otp.By("Create ClusterLogForwarder for performance testing using ViaQ datamodel")
		clf := clusterlogforwarder{
			name:                      "clf-performance-84898",
			namespace:                 loggingNS,
			templateFile:              filepath.Join(loggingBaseDir, "observability.openshift.io_clusterlogforwarder", "lokistack.yaml"),
			serviceAccountName:        "logcollector-performance-84898",
			secretName:                "lokistack-performance-84898",
			collectApplicationLogs:    true,
			collectAuditLogs:          true,
			collectInfrastructureLogs: true,
			waitForPodReady:           true,
		}
		clf.createServiceAccount(oc)
		defer removeClusterRoleFromServiceAccount(oc, clf.namespace, clf.serviceAccountName, "logging-collector-logs-writer")
		err = addClusterRoleToServiceAccount(oc, clf.namespace, clf.serviceAccountName, "logging-collector-logs-writer")
		o.Expect(err).NotTo(o.HaveOccurred())
		defer resource{"secret", clf.secretName, clf.namespace}.clear(oc)
		ls.createSecretFromGateway(oc, clf.secretName, clf.namespace, "")
		defer clf.delete(oc)
		clf.create(oc, "LOKISTACK_NAME="+ls.name, "LOKISTACK_NAMESPACE="+ls.namespace, "INPUT_REFS=[\"application\"]")

		var (
			jsonLogFile        = filepath.Join(loggingBaseDir, "generatelog", "logging-performance-app-generator.json")
			targetTotalLogs    = int64(3600000)   // NUM_LINES: 3.6M total logs
			totalRatePerMinute = 120000.0         // RATE: 120K logs/minute total across all pods
			testDuration       = 30 * time.Minute // 30 minutes duration
		)

		compat_otp.By("Deploy application pods on all worker nodes")
		// Calculate distribution across pods
		podsPerNode := 8 // 8 pods per worker node for higher log volume
		totalPods := workerNodeCount * podsPerNode
		ratePerPodPerMinute := totalRatePerMinute / float64(totalPods) // Rate per pod per minute
		numLinesPerPod := targetTotalLogs / int64(totalPods)           // Total logs per pod
		totalExpectedLogs := targetTotalLogs
		e2e.Logf("Performance test configuration:")
		e2e.Logf("  - Worker nodes: %d", workerNodeCount)
		e2e.Logf("  - Pods per node: %d", podsPerNode)
		e2e.Logf("  - Total pods: %d", totalPods)
		e2e.Logf("  - Rate per pod: %.1f logs/minute", ratePerPodPerMinute)
		e2e.Logf("  - NUM_LINES per pod: %d", numLinesPerPod)
		e2e.Logf("  - Total rate: %.0f logs/minute", totalRatePerMinute)
		e2e.Logf("  - Target total logs: %d (%.1fM)", targetTotalLogs, float64(targetTotalLogs)/1000000)
		e2e.Logf("  - Test duration: %v", testDuration)

		g.By("Create application generator project with pods distributed across all worker nodes")
		oc.SetupProject()
		appProj := oc.Namespace()

		if workerNodeCount == 1 {
			// For single worker node, create one ReplicationController with all pods
			e2e.Logf("Single worker node detected, deploying all %d pods on node: %s", totalPods, nodes.Items[0].Name)
			err = oc.WithoutNamespace().Run("new-app").Args("-n", appProj, "-f", jsonLogFile,
				"-p", fmt.Sprintf("RATE=%.1f", ratePerPodPerMinute),
				"-p", fmt.Sprintf("NUM_LINES=%d", numLinesPerPod),
				"-p", fmt.Sprintf("REPLICAS=%d", totalPods)).Execute()
			o.Expect(err).NotTo(o.HaveOccurred())
		} else {
			// Deploy pods per worker node to ensure distribution
			for i, node := range nodes.Items {
				// Create a separate ReplicationController for each worker node with nodeSelector
				rcName := fmt.Sprintf("logging-centos-logtest-node-%d", i)
				configMapName := fmt.Sprintf("logtest-config-node-%d", i)

				// Deploy pods for this specific node - nodeSelector will use default worker node selector from template
				err = oc.WithoutNamespace().Run("new-app").Args("-n", appProj, "-f", jsonLogFile,
					"-p", fmt.Sprintf("RATE=%.1f", ratePerPodPerMinute),
					"-p", fmt.Sprintf("NUM_LINES=%d", numLinesPerPod),
					"-p", fmt.Sprintf("REPLICAS=%d", podsPerNode),
					"-p", fmt.Sprintf("REPLICATIONCONTROLLER=%s", rcName),
					"-p", fmt.Sprintf("CONFIGMAP=%s", configMapName)).Execute()
				o.Expect(err).NotTo(o.HaveOccurred())

				e2e.Logf("Deployed %d pods on worker node: %s", podsPerNode, node.Name)
			}
		}

		g.By("Start log loss measurement period")
		appPods, err := oc.AdminKubeClient().CoreV1().Pods(appProj).List(context.Background(), metav1.ListOptions{LabelSelector: "run=centos-logtest"})
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(len(appPods.Items)).Should(o.Equal(totalPods), fmt.Sprintf("Should have %d pods, got %d", totalPods, len(appPods.Items)))

		// Use the calculated expected logs from pod distribution
		performanceMetrics.LogsSent = totalExpectedLogs

		e2e.Logf("Starting log loss measurement: Expected %d logs over %v from %d pods", totalExpectedLogs, testDuration, totalPods)

		// Record start time for log loss measurement
		measurementStartTime := time.Now()

		g.By("Measure application log performance and count metrics during test period")

		// Wait for the test duration to allow log generation
		time.Sleep(testDuration)
		measurementEndTime := time.Now()

		// Use expected logs from calculation as baseline
		performanceMetrics.LogsSent = totalExpectedLogs

		g.By("Query Loki to count ingested logs within measurement time window")

		defer removeClusterRoleFromServiceAccount(oc, oc.Namespace(), "default", "cluster-admin")
		err = addClusterRoleToServiceAccount(oc, oc.Namespace(), "default", "cluster-admin")
		o.Expect(err).NotTo(o.HaveOccurred())
		bearerToken := getSAToken(oc, "default", oc.Namespace())
		route := "https://" + getRouteAddress(oc, ls.namespace, ls.name)
		lc := newLokiClient(route).withToken(bearerToken).retry(5)

		// Build a window equal to the test duration + buffer for late ingestion.
		testDurationMinutes := int(math.Ceil(testDuration.Minutes())) + 5 // 5m buffer
		timeRangeQuery := fmt.Sprintf(
			`sum(count_over_time({log_type="application",kubernetes_namespace_name="%s"}[%dm]))`,
			appProj, testDurationMinutes,
		)
		e2e.Logf("Loki query: %s", timeRangeQuery)

		// Use FORWARD direction so the last sample is the newest, with a reasonable limit
		appLogs, err := lc.queryRange("application", timeRangeQuery, 1000, measurementStartTime, measurementEndTime, false)
		o.Expect(err).NotTo(o.HaveOccurred())

		var logsInLoki int64
		e2e.Logf("Loki query response status: %s", appLogs.Status)

		if appLogs.Status == "success" && len(appLogs.Data.Result) > 0 {
			// sum(...) should return a single time series
			result := appLogs.Data.Result[0]
			e2e.Logf("Samples in result: %d", len(result.Values))
			if n := len(result.Values); n > 0 {
				last := result.Values[n-1] // newest sample (because FORWARD)
				if valueSlice, ok := last.([]any); ok && len(valueSlice) >= 2 {
					if countStr, ok := valueSlice[1].(string); ok {
						if count, err := strconv.ParseInt(countStr, 10, 64); err == nil {
							logsInLoki = count
						} else {
							e2e.Logf("failed to parse count: %v", err)
						}
					}
				}
			}
		} else {
			e2e.Logf("Query failed or returned no results. Status: %s", appLogs.Status)
		}

		performanceMetrics.LogsReceived = logsInLoki
		e2e.Logf("Total logs received (window=%dm): %d", testDurationMinutes, logsInLoki)

		e2e.Logf("Loki query results:")
		e2e.Logf("  - Logs found in Loki: %d", logsInLoki)
		e2e.Logf("  - Expected logs: %d", performanceMetrics.LogsSent)

		// Calculate log loss percentage
		if performanceMetrics.LogsSent > 0 {
			performanceMetrics.LogLossPercentage = float64(performanceMetrics.LogsSent-performanceMetrics.LogsReceived) / float64(performanceMetrics.LogsSent) * 100
		}

		// Log loss validation

		e2e.Logf("Expected Logs: %d", performanceMetrics.LogsSent)
		e2e.Logf("Received Logs: %d", performanceMetrics.LogsReceived)
		e2e.Logf("Log Loss: %.2f%%", performanceMetrics.LogLossPercentage)

		o.Expect(performanceMetrics.LogsReceived).Should(o.BeNumerically(">", 0),
			"LogsReceived should be greater than 0")
		o.Expect(performanceMetrics.LogLossPercentage).Should(o.BeNumerically("<", 10.0),
			"Log loss percentage should be less than 10% for 1x.extra-small LokiStack")
	})

	g.It("Author:kbharti-Longduration-CPaasrunOnly-High-84975-Performance-Vector-LokiStack-Performance test from vector to LokiStack using 1x.extra-small tsize and Otel datamodel with log loss measurement[Serial][Slow]", func() {
		// Performance test to measure throughput and log loss from vector to LokiStack with 1x.extra-small configuration
		lokiStackTemplate := filepath.Join(loggingBaseDir, "lokistack", "lokistack-simple.yaml")

		ls := lokiStack{
			name:          "lokistack-performance-84975",
			namespace:     loggingNS,
			tSize:         "1x.extra-small",
			storageType:   getStorageType(oc),
			storageSecret: "storage-secret-performance-84975",
			storageClass:  sc,
			bucketName:    "logging-loki-performance-" + getRandomString(),
			template:      lokiStackTemplate,
		}

		defer ls.removeObjectStorage(oc)
		err := ls.prepareResourcesForLokiStack(oc)
		o.Expect(err).NotTo(o.HaveOccurred())
		defer ls.removeLokiStack(oc)
		err = ls.deployLokiStack(oc)
		o.Expect(err).NotTo(o.HaveOccurred())
		ls.waitForLokiStackToBeReady(oc)

		compat_otp.By("Patch LokiStack with custom ingestion limits for performance testing")
		patchData := `{
				"spec": {
					"limits": {
						"global": {
							"ingestion": {
								"ingestionBurstSize": 50,
								"ingestionRate": 16,
								"maxGlobalStreamsPerTenant": 5000
							}
						}
					}
				}
			}`
		err = oc.AsAdmin().WithoutNamespace().Run("patch").Args("lokistack", ls.name, "-n", ls.namespace,
			"-p", patchData, "--type=merge").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		e2e.Logf("LokiStack patched with ingestionRate: 16, ingestionBurstSize: 50, maxGlobalStreamsPerTenant: 5000")
		ls.waitForLokiStackToBeReady(oc)

		compat_otp.By("Create ClusterLogForwarder for performance testing using Otel datamodel")
		clf := clusterlogforwarder{
			name:                      "clf-performance-84975",
			namespace:                 loggingNS,
			templateFile:              filepath.Join(loggingBaseDir, "observability.openshift.io_clusterlogforwarder", "lokistack.yaml"),
			serviceAccountName:        "logcollector-performance-84975",
			secretName:                "lokistack-performance-84975",
			collectApplicationLogs:    true,
			collectAuditLogs:          true,
			collectInfrastructureLogs: true,
			waitForPodReady:           true,
		}
		clf.createServiceAccount(oc)
		defer removeClusterRoleFromServiceAccount(oc, clf.namespace, clf.serviceAccountName, "logging-collector-logs-writer")
		err = addClusterRoleToServiceAccount(oc, clf.namespace, clf.serviceAccountName, "logging-collector-logs-writer")
		o.Expect(err).NotTo(o.HaveOccurred())
		defer resource{"secret", clf.secretName, clf.namespace}.clear(oc)
		ls.createSecretFromGateway(oc, clf.secretName, clf.namespace, "")
		defer clf.delete(oc)
		clf.create(oc, "LOKISTACK_NAME="+ls.name, "LOKISTACK_NAMESPACE="+ls.namespace, "DATAMODEL=Otel", "INPUT_REFS=[\"application\"]")

		var (
			jsonLogFile        = filepath.Join(loggingBaseDir, "generatelog", "logging-performance-app-generator.json")
			targetTotalLogs    = int64(3600000)   // NUM_LINES: 3.6M total logs
			totalRatePerMinute = 120000.0         // RATE: 120K logs/minute total across all pods
			testDuration       = 30 * time.Minute // 30 minutes duration
		)

		compat_otp.By("Deploy application pods on all worker nodes")
		// Calculate distribution across pods
		podsPerNode := 8 // 8 pods per worker node for higher log volume
		totalPods := workerNodeCount * podsPerNode
		ratePerPodPerMinute := totalRatePerMinute / float64(totalPods) // Rate per pod per minute
		numLinesPerPod := targetTotalLogs / int64(totalPods)           // Total logs per pod
		totalExpectedLogs := targetTotalLogs
		e2e.Logf("Performance test configuration:")
		e2e.Logf("  - Worker nodes: %d", workerNodeCount)
		e2e.Logf("  - Pods per node: %d", podsPerNode)
		e2e.Logf("  - Total pods: %d", totalPods)
		e2e.Logf("  - Rate per pod: %.1f logs/minute", ratePerPodPerMinute)
		e2e.Logf("  - NUM_LINES per pod: %d", numLinesPerPod)
		e2e.Logf("  - Total rate: %.0f logs/minute", totalRatePerMinute)
		e2e.Logf("  - Target total logs: %d (%.1fM)", targetTotalLogs, float64(targetTotalLogs)/1000000)
		e2e.Logf("  - Test duration: %v", testDuration)

		g.By("Create application generator project with pods distributed across all worker nodes")
		oc.SetupProject()
		appProj := oc.Namespace()

		if workerNodeCount == 1 {
			// For single worker node, create one ReplicationController with all pods
			e2e.Logf("Single worker node detected, deploying all %d pods on node: %s", totalPods, nodes.Items[0].Name)
			err = oc.WithoutNamespace().Run("new-app").Args("-n", appProj, "-f", jsonLogFile,
				"-p", fmt.Sprintf("RATE=%.1f", ratePerPodPerMinute),
				"-p", fmt.Sprintf("NUM_LINES=%d", numLinesPerPod),
				"-p", fmt.Sprintf("REPLICAS=%d", totalPods)).Execute()
			o.Expect(err).NotTo(o.HaveOccurred())
		} else {
			// Deploy pods per worker node to ensure distribution
			for i, node := range nodes.Items {
				// Create a separate ReplicationController for each worker node with nodeSelector
				rcName := fmt.Sprintf("logging-centos-logtest-node-%d", i)
				configMapName := fmt.Sprintf("logtest-config-node-%d", i)

				// Deploy pods for this specific node - nodeSelector will use default worker node selector from template
				err = oc.WithoutNamespace().Run("new-app").Args("-n", appProj, "-f", jsonLogFile,
					"-p", fmt.Sprintf("RATE=%.1f", ratePerPodPerMinute),
					"-p", fmt.Sprintf("NUM_LINES=%d", numLinesPerPod),
					"-p", fmt.Sprintf("REPLICAS=%d", podsPerNode),
					"-p", fmt.Sprintf("REPLICATIONCONTROLLER=%s", rcName),
					"-p", fmt.Sprintf("CONFIGMAP=%s", configMapName)).Execute()
				o.Expect(err).NotTo(o.HaveOccurred())

				e2e.Logf("Deployed %d pods on worker node: %s", podsPerNode, node.Name)
			}
		}

		g.By("Start log loss measurement period")
		appPods, err := oc.AdminKubeClient().CoreV1().Pods(appProj).List(context.Background(), metav1.ListOptions{LabelSelector: "run=centos-logtest"})
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(len(appPods.Items)).Should(o.Equal(totalPods), fmt.Sprintf("Should have %d pods, got %d", totalPods, len(appPods.Items)))

		// Use the calculated expected logs from pod distribution
		performanceMetrics.LogsSent = totalExpectedLogs

		e2e.Logf("Starting log loss measurement: Expected %d logs over %v from %d pods", totalExpectedLogs, testDuration, totalPods)

		// Record start time for log loss measurement
		measurementStartTime := time.Now()

		g.By("Measure application log performance and count metrics during test period")

		// Wait for the test duration to allow log generation
		time.Sleep(testDuration)
		measurementEndTime := time.Now()

		// Use expected logs from calculation as baseline
		performanceMetrics.LogsSent = totalExpectedLogs

		g.By("Query Loki to count ingested logs within measurement time window")

		defer removeClusterRoleFromServiceAccount(oc, oc.Namespace(), "default", "cluster-admin")
		err = addClusterRoleToServiceAccount(oc, oc.Namespace(), "default", "cluster-admin")
		o.Expect(err).NotTo(o.HaveOccurred())
		bearerToken := getSAToken(oc, "default", oc.Namespace())
		route := "https://" + getRouteAddress(oc, ls.namespace, ls.name)
		lc := newLokiClient(route).withToken(bearerToken).retry(5)

		// Build a window equal to the test duration + buffer for late ingestion.
		testDurationMinutes := int(math.Ceil(testDuration.Minutes())) + 5 // 5m buffer
		timeRangeQuery := fmt.Sprintf(
			`sum(count_over_time({log_type="application",kubernetes_namespace_name="%s"}[%dm]))`,
			appProj, testDurationMinutes,
		)
		e2e.Logf("Loki query: %s", timeRangeQuery)

		// Use FORWARD direction so the last sample is the newest, with a reasonable limit
		appLogs, err := lc.queryRange("application", timeRangeQuery, 1000, measurementStartTime, measurementEndTime, false)
		o.Expect(err).NotTo(o.HaveOccurred())

		var logsInLoki int64
		e2e.Logf("Loki query response status: %s", appLogs.Status)

		if appLogs.Status == "success" && len(appLogs.Data.Result) > 0 {
			// sum(...) should return a single time series
			result := appLogs.Data.Result[0]
			e2e.Logf("Samples in result: %d", len(result.Values))
			if n := len(result.Values); n > 0 {
				last := result.Values[n-1] // newest sample (because FORWARD)
				if valueSlice, ok := last.([]any); ok && len(valueSlice) >= 2 {
					if countStr, ok := valueSlice[1].(string); ok {
						if count, err := strconv.ParseInt(countStr, 10, 64); err == nil {
							logsInLoki = count
						} else {
							e2e.Logf("failed to parse count: %v", err)
						}
					}
				}
			}
		} else {
			e2e.Logf("Query failed or returned no results. Status: %s", appLogs.Status)
		}

		performanceMetrics.LogsReceived = logsInLoki
		e2e.Logf("Total logs received (window=%dm): %d", testDurationMinutes, logsInLoki)

		e2e.Logf("Loki query results:")
		e2e.Logf("  - Logs found in Loki: %d", logsInLoki)
		e2e.Logf("  - Expected logs: %d", performanceMetrics.LogsSent)

		// Calculate log loss percentage
		if performanceMetrics.LogsSent > 0 {
			performanceMetrics.LogLossPercentage = float64(performanceMetrics.LogsSent-performanceMetrics.LogsReceived) / float64(performanceMetrics.LogsSent) * 100
		}

		// Log loss validation

		e2e.Logf("Expected Logs: %d", performanceMetrics.LogsSent)
		e2e.Logf("Received Logs: %d", performanceMetrics.LogsReceived)
		e2e.Logf("Log Loss: %.2f%%", performanceMetrics.LogLossPercentage)

		o.Expect(performanceMetrics.LogsReceived).Should(o.BeNumerically(">", 0),
			"LogsReceived should be greater than 0")
		o.Expect(performanceMetrics.LogLossPercentage).Should(o.BeNumerically("<", 10.0),
			"Log loss percentage should be less than 10% for 1x.extra-small LokiStack")
	})
})
