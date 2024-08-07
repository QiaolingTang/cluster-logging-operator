package elasticsearch_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	obs "github.com/openshift/cluster-logging-operator/api/observability/v1"
	"github.com/openshift/cluster-logging-operator/internal/constants"
	"github.com/openshift/cluster-logging-operator/internal/generator/framework"
	. "github.com/openshift/cluster-logging-operator/internal/generator/vector/output/elasticsearch"
	. "github.com/openshift/cluster-logging-operator/test/matchers"
	corev1 "k8s.io/api/core/v1"
)

var _ = Describe("Generate Vector config", func() {
	const (
		secretName = "es-1"
		aUserName  = "testuser"
		aPassword  = "testpass"
	)
	var (
		tlsSpec = &obs.OutputTLSSpec{
			TLSSpec: obs.TLSSpec{
				CA: &obs.ValueReference{
					Key:        constants.TrustedCABundleKey,
					SecretName: secretName,
				},
				Certificate: &obs.ValueReference{
					Key:        constants.ClientCertKey,
					SecretName: secretName,
				},
				Key: &obs.SecretReference{
					Key:        constants.ClientPrivateKey,
					SecretName: secretName,
				},
			},
		}

		initOutput = func() obs.OutputSpec {
			return obs.OutputSpec{
				Type: obs.OutputTypeElasticsearch,
				Name: "es_1",
				Elasticsearch: &obs.Elasticsearch{
					URLSpec: obs.URLSpec{
						URL: "https://es.svc.infra.cluster:9200",
					},
					Index: `{.log_type||"none"}`,
					Authentication: &obs.HTTPAuthentication{
						Username: &obs.SecretReference{
							Key:        constants.ClientUsername,
							SecretName: secretName,
						},
						Password: &obs.SecretReference{
							Key:        constants.ClientPassword,
							SecretName: secretName,
						},
					},
					Version: 8,
				},
			}
		}

		secrets = map[string]*corev1.Secret{
			secretName: {
				Data: map[string][]byte{
					constants.ClientUsername: []byte(aUserName),
					constants.ClientPassword: []byte(aPassword),
				},
			},
		}
	)
	DescribeTable("For Elasticsearch output", func(visit func(spec *obs.OutputSpec), op framework.Options, expFile string) {
		exp, err := tomlContent.ReadFile(expFile)
		if err != nil {
			Fail(fmt.Sprintf("Error reading the file %q with exp config: %v", expFile, err))
		}
		outputSpec := initOutput()
		if visit != nil {
			visit(&outputSpec)
		}
		conf := New(outputSpec.Name, outputSpec, []string{"application"}, secrets, nil, op)
		Expect(string(exp)).To(EqualConfigFrom(conf))
	},
		Entry("with username,password", nil, framework.NoOptions, "es_with_auth_username_password.toml"),
		Entry("with tls key,cert,ca-bundle", func(spec *obs.OutputSpec) {
			spec.Elasticsearch.Authentication = nil
			spec.TLS = tlsSpec
			spec.Elasticsearch.Version = 6
		}, framework.NoOptions, "es_with_tls.toml"),
		Entry("without security", func(spec *obs.OutputSpec) {
			spec.Elasticsearch.Authentication = nil
			spec.Elasticsearch.Index = "foo"
		}, framework.NoOptions, "es_without_security.toml"),
		Entry("without secret and TLS.insecureSkipVerify=true", func(spec *obs.OutputSpec) {
			spec.Elasticsearch.Authentication = nil
			spec.TLS = &obs.OutputTLSSpec{
				InsecureSkipVerify: true,
			}
		}, framework.NoOptions, "es_with_tls_skip_verify.toml"),
		Entry("with custom index with static and dynamic values", func(spec *obs.OutputSpec) {
			spec.Elasticsearch.Authentication = nil
			spec.Elasticsearch.Index = `foo-{.kubernetes.namespace||"none"}`
		}, framework.NoOptions, "es_with_custom_index.toml"),
	)
})
