// Code generated for package testdata by go-bindata DO NOT EDIT. (@generated)
// sources:
// testdata/fixtures.go
// testdata/logging/OWNERS
// testdata/logging/UIPlugin/UIPlugin.yaml
// testdata/logging/eventrouter/eventrouter.yaml
// testdata/logging/external-log-stores/cert_generation.sh
// testdata/logging/external-log-stores/elasticsearch/6/http/no_user/configmap.yaml
// testdata/logging/external-log-stores/elasticsearch/6/http/no_user/deployment.yaml
// testdata/logging/external-log-stores/elasticsearch/6/http/user_auth/configmap.yaml
// testdata/logging/external-log-stores/elasticsearch/6/http/user_auth/deployment.yaml
// testdata/logging/external-log-stores/elasticsearch/6/https/no_user/configmap.yaml
// testdata/logging/external-log-stores/elasticsearch/6/https/no_user/deployment.yaml
// testdata/logging/external-log-stores/elasticsearch/6/https/user_auth/configmap.yaml
// testdata/logging/external-log-stores/elasticsearch/6/https/user_auth/deployment.yaml
// testdata/logging/external-log-stores/elasticsearch/7/http/no_user/configmap.yaml
// testdata/logging/external-log-stores/elasticsearch/7/http/no_user/deployment.yaml
// testdata/logging/external-log-stores/elasticsearch/7/http/user_auth/configmap.yaml
// testdata/logging/external-log-stores/elasticsearch/7/http/user_auth/deployment.yaml
// testdata/logging/external-log-stores/elasticsearch/7/https/no_user/configmap.yaml
// testdata/logging/external-log-stores/elasticsearch/7/https/no_user/deployment.yaml
// testdata/logging/external-log-stores/elasticsearch/7/https/user_auth/configmap.yaml
// testdata/logging/external-log-stores/elasticsearch/7/https/user_auth/deployment.yaml
// testdata/logging/external-log-stores/elasticsearch/8/http/no_user/configmap.yaml
// testdata/logging/external-log-stores/elasticsearch/8/http/no_user/deployment.yaml
// testdata/logging/external-log-stores/elasticsearch/8/http/user_auth/configmap.yaml
// testdata/logging/external-log-stores/elasticsearch/8/http/user_auth/deployment.yaml
// testdata/logging/external-log-stores/elasticsearch/8/https/no_user/configmap.yaml
// testdata/logging/external-log-stores/elasticsearch/8/https/no_user/deployment.yaml
// testdata/logging/external-log-stores/elasticsearch/8/https/user_auth/configmap.yaml
// testdata/logging/external-log-stores/elasticsearch/8/https/user_auth/deployment.yaml
// testdata/logging/external-log-stores/fluentd/insecure/configmap.yaml
// testdata/logging/external-log-stores/fluentd/insecure/deployment.yaml
// testdata/logging/external-log-stores/fluentd/insecure/http-configmap.yaml
// testdata/logging/external-log-stores/fluentd/secure/cm-mtls-share.yaml
// testdata/logging/external-log-stores/fluentd/secure/cm-mtls.yaml
// testdata/logging/external-log-stores/fluentd/secure/cm-serverauth-share.yaml
// testdata/logging/external-log-stores/fluentd/secure/cm-serverauth.yaml
// testdata/logging/external-log-stores/fluentd/secure/deployment.yaml
// testdata/logging/external-log-stores/fluentd/secure/http-cm-mtls.yaml
// testdata/logging/external-log-stores/fluentd/secure/http-cm-serverauth.yaml
// testdata/logging/external-log-stores/kafka/amqstreams/kafka-no-auth-cluster.yaml
// testdata/logging/external-log-stores/kafka/amqstreams/kafka-no-auth-consumer-job.yaml
// testdata/logging/external-log-stores/kafka/amqstreams/kafka-sasl-cluster.yaml
// testdata/logging/external-log-stores/kafka/amqstreams/kafka-sasl-consumer-job.yaml
// testdata/logging/external-log-stores/kafka/amqstreams/kafka-sasl-consumers-config.yaml
// testdata/logging/external-log-stores/kafka/amqstreams/kafka-sasl-user.yaml
// testdata/logging/external-log-stores/kafka/amqstreams/kafka-topic.yaml
// testdata/logging/external-log-stores/kafka/cert_generation.sh
// testdata/logging/external-log-stores/kafka/kafka-rbac.yaml
// testdata/logging/external-log-stores/kafka/kafka-svc.yaml
// testdata/logging/external-log-stores/kafka/plaintext-ssl/consumer-configmap.yaml
// testdata/logging/external-log-stores/kafka/plaintext-ssl/kafka-configmap.yaml
// testdata/logging/external-log-stores/kafka/plaintext-ssl/kafka-consumer-deployment.yaml
// testdata/logging/external-log-stores/kafka/plaintext-ssl/kafka-statefulset.yaml
// testdata/logging/external-log-stores/kafka/sasl-plaintext/consumer-configmap.yaml
// testdata/logging/external-log-stores/kafka/sasl-plaintext/kafka-configmap.yaml
// testdata/logging/external-log-stores/kafka/sasl-plaintext/kafka-consumer-deployment.yaml
// testdata/logging/external-log-stores/kafka/sasl-plaintext/kafka-statefulset.yaml
// testdata/logging/external-log-stores/kafka/sasl-ssl/consumer-configmap.yaml
// testdata/logging/external-log-stores/kafka/sasl-ssl/kafka-configmap.yaml
// testdata/logging/external-log-stores/kafka/sasl-ssl/kafka-consumer-deployment.yaml
// testdata/logging/external-log-stores/kafka/sasl-ssl/kafka-statefulset.yaml
// testdata/logging/external-log-stores/kafka/zookeeper/configmap-ssl.yaml
// testdata/logging/external-log-stores/kafka/zookeeper/configmap.yaml
// testdata/logging/external-log-stores/kafka/zookeeper/zookeeper-statefulset.yaml
// testdata/logging/external-log-stores/kafka/zookeeper/zookeeper-svc.yaml
// testdata/logging/external-log-stores/loki/loki-configmap.yaml
// testdata/logging/external-log-stores/loki/loki-deployment.yaml
// testdata/logging/external-log-stores/otel/otel-collector.yaml
// testdata/logging/external-log-stores/rsyslog/insecure/configmap.yaml
// testdata/logging/external-log-stores/rsyslog/insecure/deployment.yaml
// testdata/logging/external-log-stores/rsyslog/insecure/svc.yaml
// testdata/logging/external-log-stores/rsyslog/secure/configmap.yaml
// testdata/logging/external-log-stores/rsyslog/secure/deployment.yaml
// testdata/logging/external-log-stores/rsyslog/secure/svc.yaml
// testdata/logging/external-log-stores/splunk/route-edge_splunk_template.yaml
// testdata/logging/external-log-stores/splunk/route-passthrough_splunk_template.yaml
// testdata/logging/external-log-stores/splunk/secret_splunk_template.yaml
// testdata/logging/external-log-stores/splunk/secret_tls_passphrase_splunk_template.yaml
// testdata/logging/external-log-stores/splunk/secret_tls_splunk_template.yaml
// testdata/logging/external-log-stores/splunk/statefulset_splunk-8.2_template.yaml
// testdata/logging/external-log-stores/splunk/statefulset_splunk-9.0_template.yaml
// testdata/logging/generatelog/42981.yaml
// testdata/logging/generatelog/container_json_log_template.json
// testdata/logging/generatelog/container_json_log_template_unannoted.json
// testdata/logging/generatelog/container_non_json_log_template.json
// testdata/logging/generatelog/logging-performance-app-generator.json
// testdata/logging/generatelog/multi_container_json_log_template.yaml
// testdata/logging/generatelog/multiline-error-log.yaml
// testdata/logging/logfilemetricexporter/lfme.yaml
// testdata/logging/loki-log-alerts/cluster-monitoring-config.yaml
// testdata/logging/loki-log-alerts/loki-app-alerting-rule-template.yaml
// testdata/logging/loki-log-alerts/loki-app-recording-rule-template.yaml
// testdata/logging/loki-log-alerts/loki-infra-alerting-rule-template.yaml
// testdata/logging/loki-log-alerts/loki-infra-recording-rule-template.yaml
// testdata/logging/loki-log-alerts/user-workload-monitoring-config.yaml
// testdata/logging/lokistack/lokistack-simple-ipv6-tls.yaml
// testdata/logging/lokistack/lokistack-simple-ipv6.yaml
// testdata/logging/lokistack/lokistack-simple-tls.yaml
package testdata

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _fixturesGo = []byte(`package testdata

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var (
	fixtureDir string
)

func init() {
	var err error
	fixtureDir, err = ioutil.TempDir("", "testdata-fixtures-")
	if err != nil {
		panic(fmt.Sprintf("failed to create fixture directory: %v", err))
	}
}

func FixturePath(elem ...string) string {
	relativePath := filepath.Join(elem...)
	targetPath := filepath.Join(fixtureDir, relativePath)

	if _, err := os.Stat(targetPath); err == nil {
		return targetPath
	}

	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		panic(fmt.Sprintf("failed to create directory for %s: %v", relativePath, err))
	}

	bindataPath := relativePath
	tempDir, err := os.MkdirTemp("", "bindata-extract-")
	if err != nil {
		panic(fmt.Sprintf("failed to create temp directory: %v", err))
	}
	defer os.RemoveAll(tempDir)

	if err := RestoreAsset(tempDir, bindataPath); err != nil {
		if err := RestoreAssets(tempDir, bindataPath); err != nil {
			panic(fmt.Sprintf("failed to restore fixture %s: %v", relativePath, err))
		}
	}

	extractedPath := filepath.Join(tempDir, bindataPath)
	if err := os.Rename(extractedPath, targetPath); err != nil {
		panic(fmt.Sprintf("failed to move extracted files: %v", err))
	}

	return targetPath
}

func CleanupFixtures() error {
	if fixtureDir != "" {
		return os.RemoveAll(fixtureDir)
	}
	return nil
}

func GetFixtureData(elem ...string) ([]byte, error) {
	relativePath := filepath.Join(elem...)
	cleanPath := relativePath
	if len(cleanPath) > 0 && cleanPath[0] == '/' {
		cleanPath = cleanPath[1:]
	}
	return Asset(cleanPath)
}

func MustGetFixtureData(elem ...string) []byte {
	data, err := GetFixtureData(elem...)
	if err != nil {
		panic(fmt.Sprintf("failed to get fixture data: %v", err))
	}
	return data
}

func FixtureExists(elem ...string) bool {
	relativePath := filepath.Join(elem...)
	cleanPath := relativePath
	if len(cleanPath) > 0 && cleanPath[0] == '/' {
		cleanPath = cleanPath[1:]
	}
	_, err := Asset(cleanPath)
	return err == nil
}

func ListFixtures() []string {
	names := AssetNames()
	fixtures := make([]string, 0, len(names))
	for _, name := range names {
		if strings.HasPrefix(name, "testdata/") {
			fixtures = append(fixtures, strings.TrimPrefix(name, "testdata/"))
		}
	}
	sort.Strings(fixtures)
	return fixtures
}
`)

func fixturesGoBytes() ([]byte, error) {
	return _fixturesGo, nil
}

func fixturesGo() (*asset, error) {
	bytes, err := fixturesGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fixtures.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingOwners = []byte(`approvers:
  - QiaolingTang
`)

func loggingOwnersBytes() ([]byte, error) {
	return _loggingOwners, nil
}

func loggingOwners() (*asset, error) {
	bytes, err := loggingOwnersBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/OWNERS", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingUipluginUipluginYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: subscription-template
objects:
- apiVersion: observability.openshift.io/v1alpha1
  kind: UIPlugin
  metadata:
    name: logging
  spec:
    logging:
      logsLimit: 50
      lokiStack:
        name: ${LOKISTACK_NAME}
    type: Logging
parameters:
  - name: LOKISTACK_NAME
    value: logging-loki
`)

func loggingUipluginUipluginYamlBytes() ([]byte, error) {
	return _loggingUipluginUipluginYaml, nil
}

func loggingUipluginUipluginYaml() (*asset, error) {
	bytes, err := loggingUipluginUipluginYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/UIPlugin/UIPlugin.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingEventrouterEventrouterYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: eventrouter-template
  annotations:
    description: "A pod forwarding kubernetes events to OpenShift Logging stack."
    tags: "events,EFK,logging,cluster-logging"
objects:
  - kind: ServiceAccount
    apiVersion: v1
    metadata:
      name: ${NAME}
      namespace: ${NAMESPACE}
  - kind: ClusterRole
    apiVersion: v1
    metadata:
      name: ${NAME}-reader
    rules:
    - apiGroups: [""]
      resources: ["events"]
      verbs: ["get", "watch", "list"]
  - kind: ClusterRoleBinding
    apiVersion: v1
    metadata:
      name: ${NAME}-reader-binding
    subjects:
    - kind: ServiceAccount
      name: ${NAME}
      namespace: ${NAMESPACE}
    roleRef:
      kind: ClusterRole
      name: ${NAME}-reader
  - kind: ConfigMap
    apiVersion: v1
    metadata:
      name: ${NAME}
      namespace: ${NAMESPACE}
    data:
      config.json: |-
        {
          "sink": "stdout"
        }
  - kind: Deployment
    apiVersion: apps/v1
    metadata:
      name: ${NAME}
      namespace: ${NAMESPACE}
      labels:
        component: "eventrouter"
        logging-infra: "eventrouter"
        provider: "openshift"
    spec:
      selector:
        matchLabels:
          component: "eventrouter"
          logging-infra: "eventrouter"
          provider: "openshift"
      replicas: 1
      template:
        metadata:
          labels:
            component: "eventrouter"
            logging-infra: "eventrouter"
            provider: "openshift"
          name: ${NAME}
        spec:
          serviceAccount: ${NAME}
          containers:
            - name: kube-eventrouter
              image: ${IMAGE}
              imagePullPolicy: IfNotPresent
              resources:
                requests:
                  cpu: ${CPU}
                  memory: ${MEMORY}
              volumeMounts:
              - name: config-volume
                mountPath: /etc/eventrouter
          volumes:
            - name: config-volume
              configMap:
                name: ${NAME}
parameters:
  - name: IMAGE
    displayName: Image
    value: "brew.registry.redhat.io/rh-osbs/openshift-logging-eventrouter-rhel9:v0.4.0"
  - name: CPU
    displayName: CPU
    value: "100m"
  - name: MEMORY
    displayName: Memory
    value: "128Mi"
  - name: NAMESPACE
    displayName: Namespace
    value: "openshift-logging"
  - name: NAME
    value: eventrouter
    displayName: Event Router name
`)

func loggingEventrouterEventrouterYamlBytes() ([]byte, error) {
	return _loggingEventrouterEventrouterYaml, nil
}

func loggingEventrouterEventrouterYaml() (*asset, error) {
	bytes, err := loggingEventrouterEventrouterYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/eventrouter/eventrouter.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresCert_generationSh = []byte(`#! /bin/bash

WORKING_DIR=${1:-/tmp/_working_dir}
NAMESPACE=${2:-openshift-logging}
CA_PATH=${CA_PATH:-$WORKING_DIR/ca.crt}
LOG_STORE=${3:-elasticsearch}
BASE_DOMAIN=${4:-}
PASS_PHRASE=${5:-}
REGENERATE_NEEDED=0

function init_cert_files() {

  if [ ! -f ${WORKING_DIR}/ca.db ]; then
    touch ${WORKING_DIR}/ca.db
  fi

  if [ ! -f ${WORKING_DIR}/ca.serial.txt ]; then
    echo 00 >${WORKING_DIR}/ca.serial.txt
  fi
}

function generate_signing_ca() {
  if [ ! -f ${WORKING_DIR}/ca.crt ] || [ ! -f ${WORKING_DIR}/ca.key ] || ! openssl x509 -checkend 0 -noout -in ${WORKING_DIR}/ca.crt; then
    openssl req -x509 \
      -new \
      -newkey rsa:4096 \
      -keyout ${WORKING_DIR}/ca.key \
      -nodes \
      -days 1825 \
      -out ${WORKING_DIR}/ca.crt \
      -subj "/CN=openshift-cluster-logging-signer"

    REGENERATE_NEEDED=1
  fi
}

function create_signing_conf() {
  cat <<EOF >"${WORKING_DIR}/signing.conf"
# Simple Signing CA

# The [default] section contains global constants that can be referred to from
# the entire configuration file. It may also hold settings pertaining to more
# than one openssl command.

[ default ]
dir                     = ${WORKING_DIR}               # Top dir

# The next part of the configuration file is used by the openssl req command.
# It defines the CA's key pair, its DN, and the desired extensions for the CA
# certificate.

[ req ]
default_bits            = 4096                  # RSA key size
encrypt_key             = yes                   # Protect private key
default_md              = sha512                # MD to use
utf8                    = yes                   # Input is UTF-8
string_mask             = utf8only              # Emit UTF-8 strings
prompt                  = no                    # Don't prompt for DN
distinguished_name      = ca_dn                 # DN section
req_extensions          = ca_reqext             # Desired extensions

[ ca_dn ]
0.domainComponent       = "io"
1.domainComponent       = "openshift"
organizationName        = "OpenShift Origin"
organizationalUnitName  = "Logging Signing CA"
commonName              = "Logging Signing CA"

[ ca_reqext ]
keyUsage                = critical,keyCertSign,cRLSign
basicConstraints        = critical,CA:true,pathlen:0
subjectKeyIdentifier    = hash

# The remainder of the configuration file is used by the openssl ca command.
# The CA section defines the locations of CA assets, as well as the policies
# applying to the CA.

[ ca ]
default_ca              = signing_ca            # The default CA section

[ signing_ca ]
certificate             = \$dir/ca.crt       # The CA cert
private_key             = \$dir/ca.key # CA private key
new_certs_dir           = \$dir/           # Certificate archive
serial                  = \$dir/ca.serial.txt # Serial number file
crlnumber               = \$dir/ca.crl.srl # CRL number file
database                = \$dir/ca.db # Index file
unique_subject          = no                    # Require unique subject
default_days            = 730                   # How long to certify for
default_md              = sha512                # MD to use
policy                  = any_pol             # Default naming policy
email_in_dn             = no                    # Add email to cert DN
preserve                = no                    # Keep passed DN ordering
name_opt                = ca_default            # Subject DN display options
cert_opt                = ca_default            # Certificate display options
copy_extensions         = copy                  # Copy extensions from CSR
x509_extensions         = client_ext             # Default cert extensions
default_crl_days        = 7                     # How long before next CRL
crl_extensions          = crl_ext               # CRL extensions

# Naming policies control which parts of a DN end up in the certificate and
# under what circumstances certification should be denied.

[ match_pol ]
domainComponent         = match                 # Must match 'simple.org'
organizationName        = match                 # Must match 'Simple Inc'
organizationalUnitName  = optional              # Included if present
commonName              = supplied              # Must be present

[ any_pol ]
domainComponent         = optional
countryName             = optional
stateOrProvinceName     = optional
localityName            = optional
organizationName        = optional
organizationalUnitName  = optional
commonName              = optional
emailAddress            = optional

# Certificate extensions define what types of certificates the CA is able to
# create.

[ client_ext ]
keyUsage                = critical,digitalSignature,keyEncipherment
basicConstraints        = CA:false
extendedKeyUsage        = clientAuth
subjectKeyIdentifier    = hash
authorityKeyIdentifier  = keyid

[ server_ext ]
keyUsage                = critical,digitalSignature,keyEncipherment
basicConstraints        = CA:false
extendedKeyUsage        = serverAuth,clientAuth
subjectKeyIdentifier    = hash
authorityKeyIdentifier  = keyid

# CRL extensions exist solely to point to the CA certificate that has issued
# the CRL.

[ crl_ext ]
authorityKeyIdentifier  = keyid
EOF
}

function sign_cert() {
  local component=$1

  openssl ca \
    -in ${WORKING_DIR}/${component}.csr \
    -notext \
    -out ${WORKING_DIR}/${component}.crt \
    -config ${WORKING_DIR}/signing.conf \
    -extensions v3_req \
    -batch \
    -extensions server_ext
}

function generate_cert_config() {
  local component=$1
  local extensions=${2:-}

  if [ "$extensions" != "" ]; then
    cat <<EOF >"${WORKING_DIR}/${component}.conf"
[ req ]
default_bits = 4096
prompt = no
encrypt_key = yes
default_md = sha512
distinguished_name = dn
req_extensions = req_ext
[ dn ]
CN = ${component}
OU = OpenShift
O = Logging
[ req_ext ]
subjectAltName = ${extensions}
EOF
  else
    cat <<EOF >"${WORKING_DIR}/${component}.conf"
[ req ]
default_bits = 4096
prompt = no
encrypt_key = yes
default_md = sha512
distinguished_name = dn
[ dn ]
CN = ${component}
OU = OpenShift
O = Logging
EOF
  fi
}

function generate_request() {
  local component=$1

  if [[ "$component" == "server" ]] || [[ -z "$PASS_PHRASE" ]]; then
    openssl req -new \
      -out ${WORKING_DIR}/${component}.csr \
      -newkey rsa:4096 \
      -keyout ${WORKING_DIR}/${component}.key \
      -config ${WORKING_DIR}/${component}.conf \
      -days 712 \
      -nodes
  else
    openssl req -new \
      -passout pass:"$PASS_PHRASE" \
      -out ${WORKING_DIR}/${component}.csr \
      -newkey rsa:4096 \
      -keyout ${WORKING_DIR}/${component}.key.pem \
      -config ${WORKING_DIR}/${component}.conf \
      -days 712
    # use pkcs8 for client key to avoid htting issue in FIPS cluster
    openssl pkcs8 -passin pass:"$PASS_PHRASE" -in ${WORKING_DIR}/${component}.key.pem -topk8 -nocrypt -passout pass:"$PASS_PHRASE" -out ${WORKING_DIR}/${component}.key
  fi

}

function generate_certs() {
  local component=$1
  local extensions=${2:-}

  if [ $REGENERATE_NEEDED = 1 ] || [ ! -f ${WORKING_DIR}/${component}.crt ] || ! openssl x509 -checkend 0 -noout -in ${WORKING_DIR}/${component}.crt; then
    generate_cert_config $component $extensions
    generate_request $component
    sign_cert $component
  fi
}

function generate_extensions() {
  local add_oid=$1
  local add_localhost=$2
  shift
  shift
  local cert_names=$@

  extension_names=""
  extension_index=1
  local use_comma=0

  if [ "$add_localhost" == "true" ]; then
    extension_names="IP.1:127.0.0.1,DNS.1:localhost"
    extension_index=2
    use_comma=1
  fi

  for name in ${cert_names//,/}; do
    if [ $use_comma = 1 ]; then
      extension_names="${extension_names},DNS.${extension_index}:${name}"
    else
      extension_names="DNS.${extension_index}:${name}"
      use_comma=1
    fi
    extension_index=$((extension_index + 1))
  done

  if [ "$add_oid" == "true" ]; then
    extension_names="${extension_names},RID.1:1.2.3.4.5.5"
  fi

  if [ ! -z "$BASE_DOMAIN" ]; then
    extension_names="${extension_names},DNS.${extension_index}:${LOG_STORE}-${NAMESPACE}.${BASE_DOMAIN}"
  fi

  echo "$extension_names"
}

if [ ! -d $WORKING_DIR ]; then
  mkdir -p $WORKING_DIR
fi

generate_signing_ca
init_cert_files
create_signing_conf

generate_certs 'server' "$(generate_extensions false true $LOG_STORE{,-cluster}{,.${NAMESPACE}.svc}{,.cluster.local})"
generate_certs 'client' "$(generate_extensions false false $LOG_STORE{,.${NAMESPACE}.svc}{,.cluster.local})"
`)

func loggingExternalLogStoresCert_generationShBytes() ([]byte, error) {
	return _loggingExternalLogStoresCert_generationSh, nil
}

func loggingExternalLogStoresCert_generationSh() (*asset, error) {
	bytes, err := loggingExternalLogStoresCert_generationShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/cert_generation.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch6HttpNo_userConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: v1
  data:
    elasticsearch.yml: |
      node.name:  ${NAME}
      cluster.name: ${NAME}
      discovery.zen.minimum_master_nodes: 1
      network.host: 0.0.0.0
      http.port: 9200
      http.host: 0.0.0.0
      transport.host: 127.0.0.1
      discovery.type: single-node
      xpack.security.enabled: false
      xpack.security.authc.api_key.enabled: false
      xpack.monitoring.enabled : false
      xpack.license.self_generated.type: basic
      xpack.security.http.ssl.enabled: false
      xpack.ml.enabled: ${MACHINE_LEARNING}
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
- name: MACHINE_LEARNING
  value: "true"
`)

func loggingExternalLogStoresElasticsearch6HttpNo_userConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch6HttpNo_userConfigmapYaml, nil
}

func loggingExternalLogStoresElasticsearch6HttpNo_userConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch6HttpNo_userConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/6/http/no_user/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch6HttpNo_userDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    labels:
      app: ${NAME}
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        app: ${NAME}
    strategy:
      activeDeadlineSeconds: 21600
      resources: {}
      rollingParams:
        intervalSeconds: 1
        maxSurge: 25%
        maxUnavailable: 25%
        timeoutSeconds: 600
        updatePeriodSeconds: 1
      type: Recreate
    template:
      metadata:
        labels:
          app: ${NAME}
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - image: quay.io/openshifttest/elasticsearch@sha256:206dea14c8a2c8a4d408808a08e2b4dc932218b45aae6147ba000fa08cc7251a
          imagePullPolicy: IfNotPresent
          name: ${NAME}
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          ports:
          - containerPort: 9300
            protocol: TCP
          - containerPort: 9200
            protocol: TCP
          volumeMounts:
          - mountPath: /usr/share/elasticsearch/config/elasticsearch.yml
            subPath: elasticsearch.yml
            name: elasticsearch-config
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        volumes:
        - configMap:
            defaultMode: 420
            name: ${NAME}
          name: elasticsearch-config
        serviceAccount: ${NAME}
        serviceAccountName: ${NAME}
        dnsPolicy: ClusterFirst
        restartPolicy: Always
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresElasticsearch6HttpNo_userDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch6HttpNo_userDeploymentYaml, nil
}

func loggingExternalLogStoresElasticsearch6HttpNo_userDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch6HttpNo_userDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/6/http/no_user/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch6HttpUser_authConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: es6-cm-template
objects:
- apiVersion: v1
  data:
    add-user.sh: |+
      output = $(/usr/share/elasticsearch/bin/elasticsearch-users useradd ${USERNAME} -p ${PASSWORD} -r superuser)
      if [[ $output =~ 'already exists' ]]
      then
      return 0
      fi

    elasticsearch.yml: |
      node.name:  ${NAME}
      cluster.name: ${NAME}
      discovery.zen.minimum_master_nodes: 1
      network.host: 0.0.0.0
      http.port: 9200
      http.host: 0.0.0.0
      transport.host: 127.0.0.1
      discovery.type: single-node
      xpack.security.authc.realms:
        native:
          type: file
          order: 0
          enabled: true
          authentication.enabled: true
      xpack.security.enabled: true
      xpack.security.authc.api_key.enabled: true
      xpack.monitoring.enabled : false
      xpack.license.self_generated.type: basic
      xpack.security.http.ssl.enabled: false
      xpack.ml.enabled: ${MACHINE_LEARNING}
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: USERNAME
  value: "fluentd"
- name: PASSWORD
  value: "redhat"
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
- name: MACHINE_LEARNING
  value: "true"
`)

func loggingExternalLogStoresElasticsearch6HttpUser_authConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch6HttpUser_authConfigmapYaml, nil
}

func loggingExternalLogStoresElasticsearch6HttpUser_authConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch6HttpUser_authConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/6/http/user_auth/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch6HttpUser_authDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    labels:
      app: ${NAME}
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        app: ${NAME}
    strategy:
      activeDeadlineSeconds: 21600
      resources: {}
      rollingParams:
        intervalSeconds: 1
        maxSurge: 25%
        maxUnavailable: 25%
        timeoutSeconds: 600
        updatePeriodSeconds: 1
      type: Recreate
    template:
      metadata:
        labels:
          app: ${NAME}
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - image: quay.io/openshifttest/elasticsearch@sha256:206dea14c8a2c8a4d408808a08e2b4dc932218b45aae6147ba000fa08cc7251a
          imagePullPolicy: IfNotPresent
          name: ${NAME}
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          ports:
          - containerPort: 9300
            protocol: TCP
          - containerPort: 9200
            protocol: TCP
          readinessProbe:
            exec:
              command:
              - sh
              - /usr/share/elasticsearch/add-user.sh
            failureThreshold: 3
            initialDelaySeconds: 5
            periodSeconds: 5
            successThreshold: 1
            timeoutSeconds: 10
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /usr/share/elasticsearch/config/elasticsearch.yml
            name: elasticsearch-config
            subPath: elasticsearch.yml
          - mountPath: /usr/share/elasticsearch/add-user.sh
            name: elasticsearch-config
            subPath: add-user.sh
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        serviceAccount: ${NAME}
        serviceAccountName: ${NAME}
        terminationGracePeriodSeconds: 30
        volumes:
        - configMap:
            defaultMode: 420
            name: ${NAME}
          name: elasticsearch-config
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresElasticsearch6HttpUser_authDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch6HttpUser_authDeploymentYaml, nil
}

func loggingExternalLogStoresElasticsearch6HttpUser_authDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch6HttpUser_authDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/6/http/user_auth/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch6HttpsNo_userConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: v1
  data:
    elasticsearch.yml: |
      node.name: ${NAME}
      cluster.name: ${NAME}
      discovery.zen.minimum_master_nodes: 1
      network.host: 0.0.0.0
      http.port: 9200
      http.host: 0.0.0.0
      transport.host: 127.0.0.1
      discovery.type: single-node
      xpack.security.enabled: true
      xpack.security.authc:
        anonymous:
          username: anonymous_user
          roles: superuser
          authz_exception: true
      xpack.security.authc.api_key.enabled: true
      xpack.monitoring.enabled : false
      xpack.license.self_generated.type: basic
      xpack.security.http.ssl.enabled: true
      xpack.security.http.ssl.key:  /usr/share/elasticsearch/config/secret/elasticsearch.key
      xpack.security.http.ssl.certificate: /usr/share/elasticsearch/config/secret/elasticsearch.crt
      xpack.security.http.ssl.certificate_authorities: [ "/usr/share/elasticsearch/config/secret/admin-ca" ]
      xpack.security.http.ssl.verification_mode: full
      xpack.security.http.ssl.client_authentication: ${CLIENT_AUTH}
      xpack.ml.enabled: ${MACHINE_LEARNING}
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
- name: CLIENT_AUTH
  value: none
- name: MACHINE_LEARNING
  value: "true"
`)

func loggingExternalLogStoresElasticsearch6HttpsNo_userConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch6HttpsNo_userConfigmapYaml, nil
}

func loggingExternalLogStoresElasticsearch6HttpsNo_userConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch6HttpsNo_userConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/6/https/no_user/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch6HttpsNo_userDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    labels:
      app: ${NAME}
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        app: ${NAME}
    strategy:
      type: Recreate
    template:
      metadata:
        creationTimestamp: null
        labels:
          app: ${NAME}
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - image: quay.io/openshifttest/elasticsearch@sha256:206dea14c8a2c8a4d408808a08e2b4dc932218b45aae6147ba000fa08cc7251a
          imagePullPolicy: IfNotPresent
          name: ${NAME}
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          ports:
          - containerPort: 9300
            protocol: TCP
          - containerPort: 9200
            protocol: TCP
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /usr/share/elasticsearch/config/elasticsearch.yml
            name: elasticsearch-config
            subPath: elasticsearch.yml
          - mountPath: /usr/share/elasticsearch/config/secret
            name: certificates
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        serviceAccount: ${NAME}
        serviceAccountName: ${NAME}
        terminationGracePeriodSeconds: 30
        volumes:
        - configMap:
            defaultMode: 420
            name: ${NAME}
          name: elasticsearch-config
        - name: certificates
          secret:
            defaultMode: 420
            secretName: ${NAME}
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresElasticsearch6HttpsNo_userDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch6HttpsNo_userDeploymentYaml, nil
}

func loggingExternalLogStoresElasticsearch6HttpsNo_userDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch6HttpsNo_userDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/6/https/no_user/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch6HttpsUser_authConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: v1
  data:
    add-user.sh: |+
      output = $(/usr/share/elasticsearch/bin/elasticsearch-users useradd ${USERNAME} -p ${PASSWORD} -r superuser)
      if [[ $output =~ 'already exists' ]]
      then
      return 0
      fi

    elasticsearch.yml: |
      node.name:  ${NAME}
      cluster.name: ${NAME}
      discovery.zen.minimum_master_nodes: 1
      network.host: 0.0.0.0
      http.port: 9200
      http.host: 0.0.0.0
      transport.host: 127.0.0.1
      discovery.type: single-node
      xpack.security.authc:
        realms:
          native:
            type: file
            order: 0
            enabled: true
            authentication.enabled: true
      xpack.security.enabled: true
      xpack.security.authc.api_key.enabled: true
      xpack.monitoring.enabled : false
      xpack.license.self_generated.type: basic
      xpack.security.http.ssl.enabled: true
      xpack.security.http.ssl.key:  /usr/share/elasticsearch/config/secret/elasticsearch.key
      xpack.security.http.ssl.certificate: /usr/share/elasticsearch/config/secret/elasticsearch.crt
      xpack.security.http.ssl.certificate_authorities: [ "/usr/share/elasticsearch/config/secret/admin-ca" ]
      xpack.security.http.ssl.verification_mode: full
      xpack.security.http.ssl.client_authentication: ${CLIENT_AUTH}
      xpack.ml.enabled: ${MACHINE_LEARNING}
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: USERNAME
  value: "fluentd"
- name: PASSWORD
  value: "redhat"
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
- name: CLIENT_AUTH
  value: none
- name: MACHINE_LEARNING
  value: "true"
`)

func loggingExternalLogStoresElasticsearch6HttpsUser_authConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch6HttpsUser_authConfigmapYaml, nil
}

func loggingExternalLogStoresElasticsearch6HttpsUser_authConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch6HttpsUser_authConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/6/https/user_auth/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch6HttpsUser_authDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    labels:
      app: ${NAME}
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        app: ${NAME}
    strategy:
      activeDeadlineSeconds: 21600
      resources: {}
      rollingParams:
        intervalSeconds: 1
        maxSurge: 25%
        maxUnavailable: 25%
        timeoutSeconds: 600
        updatePeriodSeconds: 1
      type: Recreate
    template:
      metadata:
        labels:
          app: ${NAME}
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - image: quay.io/openshifttest/elasticsearch@sha256:206dea14c8a2c8a4d408808a08e2b4dc932218b45aae6147ba000fa08cc7251a
          imagePullPolicy: IfNotPresent
          name: ${NAME}
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          ports:
          - containerPort: 9300
            protocol: TCP
          - containerPort: 9200
            protocol: TCP
          readinessProbe:
            exec:
              command:
              - sh
              - /usr/share/elasticsearch/add-user.sh
            failureThreshold: 3
            initialDelaySeconds: 5
            periodSeconds: 5
            successThreshold: 1
            timeoutSeconds: 10
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /usr/share/elasticsearch/config/elasticsearch.yml
            name: elasticsearch-config
            subPath: elasticsearch.yml
          - mountPath: /usr/share/elasticsearch/add-user.sh
            name: elasticsearch-config
            subPath: add-user.sh
          - mountPath: /usr/share/elasticsearch/config/secret
            name: certificates
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        serviceAccount: ${NAME}
        serviceAccountName: ${NAME}
        terminationGracePeriodSeconds: 30
        volumes:
        - configMap:
            defaultMode: 420
            name: ${NAME}
          name: elasticsearch-config
        - name: certificates
          secret:
            defaultMode: 420
            secretName: ${NAME}
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresElasticsearch6HttpsUser_authDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch6HttpsUser_authDeploymentYaml, nil
}

func loggingExternalLogStoresElasticsearch6HttpsUser_authDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch6HttpsUser_authDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/6/https/user_auth/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch7HttpNo_userConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: v1
  data:
    elasticsearch.yml: |
      node.name:  ${NAME}
      cluster.name: ${NAME}
      discovery.zen.minimum_master_nodes: 1
      network.host: 0.0.0.0
      http.port: 9200
      http.host: 0.0.0.0
      transport.host: 127.0.0.1
      discovery.type: single-node
      xpack.ml.enabled: ${MACHINE_LEARNING}
      xpack.security.enabled: false
      xpack.security.authc.api_key.enabled: false
      xpack.monitoring.enabled : false
      xpack.license.self_generated.type: basic
      xpack.security.http.ssl.enabled: false
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
- name: MACHINE_LEARNING
  value: "true"
`)

func loggingExternalLogStoresElasticsearch7HttpNo_userConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch7HttpNo_userConfigmapYaml, nil
}

func loggingExternalLogStoresElasticsearch7HttpNo_userConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch7HttpNo_userConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/7/http/no_user/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch7HttpNo_userDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    labels:
      app: ${NAME}
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        app: ${NAME}
    strategy:
      activeDeadlineSeconds: 21600
      resources: {}
      rollingParams:
        intervalSeconds: 1
        maxSurge: 25%
        maxUnavailable: 25%
        timeoutSeconds: 600
        updatePeriodSeconds: 1
      type: Recreate
    template:
      metadata:
        labels:
          app: ${NAME}
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - image: quay.io/openshifttest/elasticsearch@sha256:420ab335838747b1d350ed39f4d88cf075479fc915cda8f91373c7e00de65887
          imagePullPolicy: IfNotPresent
          name: ${NAME}
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          ports:
          - containerPort: 9300
            protocol: TCP
          - containerPort: 9200
            protocol: TCP
          volumeMounts:
          - mountPath: /usr/share/elasticsearch/config/elasticsearch.yml
            subPath: elasticsearch.yml
            name: elasticsearch-config
          resources:
            limits:
              cpu: 1
              memory: 2Gi
            requests:
              cpu: 1
              memory: 2Gi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        volumes:
        - configMap:
            defaultMode: 420
            name: ${NAME}
          name: elasticsearch-config
        serviceAccount: ${NAME}
        serviceAccountName: ${NAME}
        dnsPolicy: ClusterFirst
        restartPolicy: Always
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresElasticsearch7HttpNo_userDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch7HttpNo_userDeploymentYaml, nil
}

func loggingExternalLogStoresElasticsearch7HttpNo_userDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch7HttpNo_userDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/7/http/no_user/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch7HttpUser_authConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: v1
  data:
    add-user.sh: |+
      output = $(/usr/share/elasticsearch/bin/elasticsearch-users useradd ${USERNAME} -p ${PASSWORD} -r superuser)
      if [[ $output =~ 'already exists' ]]
      then
      return 0
      fi

    elasticsearch.yml: |
      node.name:  ${NAME}
      cluster.name: ${NAME}
      discovery.zen.minimum_master_nodes: 1
      network.host: 0.0.0.0
      http.port: 9200
      http.host: 0.0.0.0
      transport.host: 127.0.0.1
      discovery.type: single-node
      xpack.security.authc.realms:
        file:
          test:
            order: 0
            enabled: true
            authentication.enabled: true
      xpack.security.enabled: true
      xpack.ml.enabled: ${MACHINE_LEARNING}
      xpack.security.authc.api_key.enabled: true
      xpack.monitoring.enabled : false
      xpack.license.self_generated.type: basic
      xpack.security.http.ssl.enabled: false
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
- name: USERNAME
  value: "fluentd"
- name: PASSWORD
  value: "redhat"
- name: MACHINE_LEARNING
  value: "true"
`)

func loggingExternalLogStoresElasticsearch7HttpUser_authConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch7HttpUser_authConfigmapYaml, nil
}

func loggingExternalLogStoresElasticsearch7HttpUser_authConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch7HttpUser_authConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/7/http/user_auth/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch7HttpUser_authDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    labels:
      app: ${NAME}
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        app: ${NAME}
    strategy:
      activeDeadlineSeconds: 21600
      resources: {}
      rollingParams:
        intervalSeconds: 1
        maxSurge: 25%
        maxUnavailable: 25%
        timeoutSeconds: 600
        updatePeriodSeconds: 1
      type: Recreate
    template:
      metadata:
        labels:
          app: ${NAME}
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - image: quay.io/openshifttest/elasticsearch@sha256:420ab335838747b1d350ed39f4d88cf075479fc915cda8f91373c7e00de65887
          imagePullPolicy: IfNotPresent
          name: ${NAME}
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          ports:
          - containerPort: 9300
            protocol: TCP
          - containerPort: 9200
            protocol: TCP
          readinessProbe:
            exec:
              command:
              - sh
              - /usr/share/elasticsearch/add-user.sh
            failureThreshold: 3
            initialDelaySeconds: 5
            periodSeconds: 5
            successThreshold: 1
            timeoutSeconds: 10
          resources:
            limits:
              cpu: 1
              memory: 2Gi
            requests:
              cpu: 1
              memory: 2Gi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /usr/share/elasticsearch/config/elasticsearch.yml
            name: elasticsearch-config
            subPath: elasticsearch.yml
          - mountPath: /usr/share/elasticsearch/add-user.sh
            name: elasticsearch-config
            subPath: add-user.sh
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        serviceAccount: ${NAME}
        serviceAccountName: ${NAME}
        terminationGracePeriodSeconds: 30
        volumes:
        - configMap:
            defaultMode: 420
            name: ${NAME}
          name: elasticsearch-config
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresElasticsearch7HttpUser_authDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch7HttpUser_authDeploymentYaml, nil
}

func loggingExternalLogStoresElasticsearch7HttpUser_authDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch7HttpUser_authDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/7/http/user_auth/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch7HttpsNo_userConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: v1
  data:
    elasticsearch.yml: |
      node.name: ${NAME}
      cluster.name: ${NAME}
      discovery.zen.minimum_master_nodes: 1
      network.host: 0.0.0.0
      http.port: 9200
      http.host: 0.0.0.0
      transport.host: 127.0.0.1
      discovery.type: single-node
      xpack.security.enabled: true
      xpack.security.authc:
        anonymous:
          username: anonymous_user
          roles: superuser
          authz_exception: true
      xpack.security.authc.api_key.enabled: true
      xpack.monitoring.enabled : false
      xpack.license.self_generated.type: basic
      xpack.security.http.ssl.enabled: true
      xpack.ml.enabled: ${MACHINE_LEARNING}
      xpack.security.http.ssl.key:  /usr/share/elasticsearch/config/secret/elasticsearch.key
      xpack.security.http.ssl.certificate: /usr/share/elasticsearch/config/secret/elasticsearch.crt
      xpack.security.http.ssl.certificate_authorities: [ "/usr/share/elasticsearch/config/secret/admin-ca" ]
      xpack.security.http.ssl.verification_mode: full
      xpack.security.http.ssl.client_authentication: ${CLIENT_AUTH}
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
- name: CLIENT_AUTH
  value: required
- name: MACHINE_LEARNING
  value: "true"
`)

func loggingExternalLogStoresElasticsearch7HttpsNo_userConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch7HttpsNo_userConfigmapYaml, nil
}

func loggingExternalLogStoresElasticsearch7HttpsNo_userConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch7HttpsNo_userConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/7/https/no_user/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch7HttpsNo_userDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    labels:
      app: ${NAME}
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        app: ${NAME}
    strategy:
      activeDeadlineSeconds: 21600
      resources: {}
      rollingParams:
        intervalSeconds: 1
        maxSurge: 25%
        maxUnavailable: 25%
        timeoutSeconds: 600
        updatePeriodSeconds: 1
      type: Recreate
    template:
      metadata:
        labels:
          app: ${NAME}
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - image: quay.io/openshifttest/elasticsearch@sha256:420ab335838747b1d350ed39f4d88cf075479fc915cda8f91373c7e00de65887
          imagePullPolicy: IfNotPresent
          name: ${NAME}
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          ports:
          - containerPort: 9300
            protocol: TCP
          - containerPort: 9200
            protocol: TCP
          resources:
            limits:
              cpu: 1
              memory: 2Gi
            requests:
              cpu: 1
              memory: 2Gi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /usr/share/elasticsearch/config/elasticsearch.yml
            name: elasticsearch-config
            subPath: elasticsearch.yml
          - mountPath: /usr/share/elasticsearch/config/secret
            name: certificates
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        serviceAccount: ${NAME}
        serviceAccountName: ${NAME}
        terminationGracePeriodSeconds: 30
        volumes:
        - configMap:
            defaultMode: 420
            name: ${NAME}
          name: elasticsearch-config
        - name: certificates
          secret:
            defaultMode: 420
            secretName: ${NAME}
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresElasticsearch7HttpsNo_userDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch7HttpsNo_userDeploymentYaml, nil
}

func loggingExternalLogStoresElasticsearch7HttpsNo_userDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch7HttpsNo_userDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/7/https/no_user/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch7HttpsUser_authConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: v1
  data:
    add-user.sh: |+
      output = $(/usr/share/elasticsearch/bin/elasticsearch-users useradd ${USERNAME} -p ${PASSWORD} -r superuser)
      if [[ $output =~ 'already exists' ]]
      then
      return 0
      fi

    elasticsearch.yml: |
      node.name:  ${NAME}
      cluster.name: ${NAME}
      discovery.zen.minimum_master_nodes: 1
      network.host: 0.0.0.0
      http.port: 9200
      http.host: 0.0.0.0
      transport.host: 127.0.0.1
      discovery.type: single-node
      xpack.security.authc.realms.file:
        test:
          order: 0
          enabled: true
          authentication.enabled: true
      xpack.security.enabled: true
      xpack.security.authc.api_key.enabled: true
      xpack.monitoring.enabled : false
      xpack.license.self_generated.type: basic
      xpack.ml.enabled: ${MACHINE_LEARNING}
      xpack.security.http.ssl.enabled: true
      xpack.security.http.ssl.key:  /usr/share/elasticsearch/config/secret/elasticsearch.key
      xpack.security.http.ssl.certificate: /usr/share/elasticsearch/config/secret/elasticsearch.crt
      xpack.security.http.ssl.certificate_authorities: [ "/usr/share/elasticsearch/config/secret/admin-ca" ]
      xpack.security.http.ssl.verification_mode: full
      xpack.security.http.ssl.client_authentication: ${CLIENT_AUTH}
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
- name: USERNAME
  value: "fluentd"
- name: PASSWORD
  value: "redhat"
- name: CLIENT_AUTH
  value: none
- name: MACHINE_LEARNING
  value: "true"
`)

func loggingExternalLogStoresElasticsearch7HttpsUser_authConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch7HttpsUser_authConfigmapYaml, nil
}

func loggingExternalLogStoresElasticsearch7HttpsUser_authConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch7HttpsUser_authConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/7/https/user_auth/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch7HttpsUser_authDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    labels:
      app: ${NAME}
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        app: ${NAME}
    strategy:
      activeDeadlineSeconds: 21600
      resources: {}
      rollingParams:
        intervalSeconds: 1
        maxSurge: 25%
        maxUnavailable: 25%
        timeoutSeconds: 600
        updatePeriodSeconds: 1
      type: Recreate
    template:
      metadata:
        labels:
          app: ${NAME}
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - image: quay.io/openshifttest/elasticsearch@sha256:420ab335838747b1d350ed39f4d88cf075479fc915cda8f91373c7e00de65887
          imagePullPolicy: IfNotPresent
          name: ${NAME}
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          ports:
          - containerPort: 9300
            protocol: TCP
          - containerPort: 9200
            protocol: TCP
          readinessProbe:
            exec:
              command:
              - sh
              - /usr/share/elasticsearch/add-user.sh
            failureThreshold: 3
            initialDelaySeconds: 5
            periodSeconds: 5
            successThreshold: 1
            timeoutSeconds: 10
          resources:
            limits:
              cpu: 1
              memory: 2Gi
            requests:
              cpu: 1
              memory: 2Gi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /usr/share/elasticsearch/config/elasticsearch.yml
            name: elasticsearch-config
            subPath: elasticsearch.yml
          - mountPath: /usr/share/elasticsearch/add-user.sh
            name: elasticsearch-config
            subPath: add-user.sh
          - mountPath: /usr/share/elasticsearch/config/secret
            name: certificates
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        serviceAccount: ${NAME}
        serviceAccountName: ${NAME}
        terminationGracePeriodSeconds: 30
        volumes:
        - configMap:
            defaultMode: 420
            name: ${NAME}
          name: elasticsearch-config
        - name: certificates
          secret:
            defaultMode: 420
            secretName: ${NAME}
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresElasticsearch7HttpsUser_authDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch7HttpsUser_authDeploymentYaml, nil
}

func loggingExternalLogStoresElasticsearch7HttpsUser_authDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch7HttpsUser_authDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/7/https/user_auth/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch8HttpNo_userConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: v1
  data:
    elasticsearch.yml: |
      node.name:  ${NAME}
      cluster.name: ${NAME}
      network.host: 0.0.0.0
      http.port: 9200
      http.host: 0.0.0.0
      transport.host: 127.0.0.1
      discovery.type: single-node
      xpack.security.enabled: false
      xpack.security.authc.api_key.enabled: false
      xpack.monitoring.collection.enabled: false
      xpack.license.self_generated.type: basic
      xpack.security.http.ssl.enabled: false
      xpack.ml.enabled: ${MACHINE_LEARNING}
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
- name: MACHINE_LEARNING
  value: "true"
`)

func loggingExternalLogStoresElasticsearch8HttpNo_userConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch8HttpNo_userConfigmapYaml, nil
}

func loggingExternalLogStoresElasticsearch8HttpNo_userConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch8HttpNo_userConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/8/http/no_user/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch8HttpNo_userDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    labels:
      app: ${NAME}
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        app: ${NAME}
    strategy:
      activeDeadlineSeconds: 21600
      resources: {}
      rollingParams:
        intervalSeconds: 1
        maxSurge: 25%
        maxUnavailable: 25%
        timeoutSeconds: 600
        updatePeriodSeconds: 1
      type: Recreate
    template:
      metadata:
        labels:
          app: ${NAME}
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - image: quay.io/openshifttest/elasticsearch@sha256:234e8ecfb6c1bdafffeb190c4a48e8e3c4b74b69d1b541b913dfbca29e952f63
          imagePullPolicy: IfNotPresent
          name: ${NAME}
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          ports:
          - containerPort: 9300
            protocol: TCP
          - containerPort: 9200
            protocol: TCP
          volumeMounts:
          - mountPath: /usr/share/elasticsearch/config/elasticsearch.yml
            subPath: elasticsearch.yml
            name: elasticsearch-config
          resources:
            limits:
              cpu: 1
              memory: 2Gi
            requests:
              cpu: 1
              memory: 2Gi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        volumes:
        - configMap:
            defaultMode: 420
            name: ${NAME}
          name: elasticsearch-config
        serviceAccount: ${NAME}
        serviceAccountName: ${NAME}
        dnsPolicy: ClusterFirst
        restartPolicy: Always
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresElasticsearch8HttpNo_userDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch8HttpNo_userDeploymentYaml, nil
}

func loggingExternalLogStoresElasticsearch8HttpNo_userDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch8HttpNo_userDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/8/http/no_user/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch8HttpUser_authConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: v1
  data:
    add-user.sh: |+
      output = $(/usr/share/elasticsearch/bin/elasticsearch-users useradd ${USERNAME} -p ${PASSWORD} -r superuser)
      if [[ $output =~ 'already exists' ]]
      then
      return 0
      fi

    elasticsearch.yml: |
      node.name:  ${NAME}
      cluster.name: ${NAME}
      network.host: 0.0.0.0
      http.port: 9200
      http.host: 0.0.0.0
      transport.host: 127.0.0.1
      discovery.type: single-node
      xpack.security.authc.realms:
        file:
          test:
            order: 0
            enabled: true
            authentication.enabled: true
      xpack.security.enabled: true
      xpack.security.authc.api_key.enabled: true
      xpack.monitoring.collection.enabled: false
      xpack.license.self_generated.type: basic
      xpack.security.http.ssl.enabled: false
      xpack.ml.enabled: ${MACHINE_LEARNING}
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
- name: USERNAME
  value: "fluentd"
- name: PASSWORD
  value: "redhat"
- name: MACHINE_LEARNING
  value: "true"
`)

func loggingExternalLogStoresElasticsearch8HttpUser_authConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch8HttpUser_authConfigmapYaml, nil
}

func loggingExternalLogStoresElasticsearch8HttpUser_authConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch8HttpUser_authConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/8/http/user_auth/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch8HttpUser_authDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    labels:
      app: ${NAME}
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        app: ${NAME}
    strategy:
      activeDeadlineSeconds: 21600
      resources: {}
      rollingParams:
        intervalSeconds: 1
        maxSurge: 25%
        maxUnavailable: 25%
        timeoutSeconds: 600
        updatePeriodSeconds: 1
      type: Recreate
    template:
      metadata:
        labels:
          app: ${NAME}
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - image: quay.io/openshifttest/elasticsearch@sha256:234e8ecfb6c1bdafffeb190c4a48e8e3c4b74b69d1b541b913dfbca29e952f63
          imagePullPolicy: IfNotPresent
          name: ${NAME}
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          ports:
          - containerPort: 9300
            protocol: TCP
          - containerPort: 9200
            protocol: TCP
          readinessProbe:
            exec:
              command:
              - sh
              - /usr/share/elasticsearch/add-user.sh
            failureThreshold: 3
            initialDelaySeconds: 5
            periodSeconds: 5
            successThreshold: 1
            timeoutSeconds: 10
          resources:
            limits:
              cpu: 1
              memory: 2Gi
            requests:
              cpu: 1
              memory: 2Gi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /usr/share/elasticsearch/config/elasticsearch.yml
            name: elasticsearch-config
            subPath: elasticsearch.yml
          - mountPath: /usr/share/elasticsearch/add-user.sh
            name: elasticsearch-config
            subPath: add-user.sh
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        serviceAccount: ${NAME}
        serviceAccountName: ${NAME}
        terminationGracePeriodSeconds: 30
        volumes:
        - configMap:
            defaultMode: 420
            name: ${NAME}
          name: elasticsearch-config
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresElasticsearch8HttpUser_authDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch8HttpUser_authDeploymentYaml, nil
}

func loggingExternalLogStoresElasticsearch8HttpUser_authDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch8HttpUser_authDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/8/http/user_auth/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch8HttpsNo_userConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: v1
  data:
    elasticsearch.yml: |
      node.name: ${NAME}
      cluster.name: ${NAME}
      network.host: 0.0.0.0
      http.port: 9200
      http.host: 0.0.0.0
      transport.host: 127.0.0.1
      discovery.type: single-node
      xpack.security.enabled: true
      xpack.security.authc:
        anonymous:
          username: anonymous_user
          roles: superuser
          authz_exception: true
      xpack.security.authc.api_key.enabled: true
      xpack.monitoring.collection.enabled: false
      xpack.license.self_generated.type: basic
      xpack.security.http.ssl.enabled: true
      xpack.security.http.ssl.key:  /usr/share/elasticsearch/config/secret/elasticsearch.key
      xpack.security.http.ssl.certificate: /usr/share/elasticsearch/config/secret/elasticsearch.crt
      xpack.security.http.ssl.certificate_authorities: [ "/usr/share/elasticsearch/config/secret/admin-ca" ]
      xpack.security.http.ssl.verification_mode: full
      xpack.security.http.ssl.client_authentication: ${CLIENT_AUTH}
      xpack.ml.enabled: ${MACHINE_LEARNING}
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
- name: CLIENT_AUTH
  value: required
- name: MACHINE_LEARNING
  value: "true"
`)

func loggingExternalLogStoresElasticsearch8HttpsNo_userConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch8HttpsNo_userConfigmapYaml, nil
}

func loggingExternalLogStoresElasticsearch8HttpsNo_userConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch8HttpsNo_userConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/8/https/no_user/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch8HttpsNo_userDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    labels:
      app: ${NAME}
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        app: ${NAME}
    strategy:
      activeDeadlineSeconds: 21600
      resources: {}
      rollingParams:
        intervalSeconds: 1
        maxSurge: 25%
        maxUnavailable: 25%
        timeoutSeconds: 600
        updatePeriodSeconds: 1
      type: Recreate
    template:
      metadata:
        labels:
          app: ${NAME}
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - image: quay.io/openshifttest/elasticsearch@sha256:234e8ecfb6c1bdafffeb190c4a48e8e3c4b74b69d1b541b913dfbca29e952f63
          imagePullPolicy: IfNotPresent
          name: ${NAME}
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          ports:
          - containerPort: 9300
            protocol: TCP
          - containerPort: 9200
            protocol: TCP
          resources:
            limits:
              cpu: 1
              memory: 2Gi
            requests:
              cpu: 1
              memory: 2Gi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /usr/share/elasticsearch/config/elasticsearch.yml
            name: elasticsearch-config
            subPath: elasticsearch.yml
          - mountPath: /usr/share/elasticsearch/config/secret
            name: certificates
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        serviceAccount: ${NAME}
        serviceAccountName: ${NAME}
        terminationGracePeriodSeconds: 30
        volumes:
        - configMap:
            defaultMode: 420
            name: ${NAME}
          name: elasticsearch-config
        - name: certificates
          secret:
            defaultMode: 420
            secretName: ${NAME}
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresElasticsearch8HttpsNo_userDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch8HttpsNo_userDeploymentYaml, nil
}

func loggingExternalLogStoresElasticsearch8HttpsNo_userDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch8HttpsNo_userDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/8/https/no_user/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch8HttpsUser_authConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: v1
  data:
    add-user.sh: |+
      output = $(/usr/share/elasticsearch/bin/elasticsearch-users useradd ${USERNAME} -p ${PASSWORD} -r superuser)
      if [[ $output =~ 'already exists' ]]
      then
      return 0
      fi

    elasticsearch.yml: |
      node.name:  ${NAME}
      cluster.name: ${NAME}
      network.host: 0.0.0.0
      http.port: 9200
      http.host: 0.0.0.0
      transport.host: 127.0.0.1
      discovery.type: single-node
      xpack.security.authc.realms.file:
        test:
          order: 0
          enabled: true
          authentication.enabled: true
      xpack.security.enabled: true
      xpack.security.authc.api_key.enabled: true
      xpack.monitoring.collection.enabled: false
      xpack.license.self_generated.type: basic
      xpack.security.http.ssl.enabled: true
      xpack.security.http.ssl.key:  /usr/share/elasticsearch/config/secret/elasticsearch.key
      xpack.security.http.ssl.certificate: /usr/share/elasticsearch/config/secret/elasticsearch.crt
      xpack.security.http.ssl.certificate_authorities: [ "/usr/share/elasticsearch/config/secret/admin-ca" ]
      xpack.security.http.ssl.verification_mode: full
      xpack.security.http.ssl.client_authentication: ${CLIENT_AUTH}
      xpack.ml.enabled: ${MACHINE_LEARNING}
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
- name: USERNAME
  value: "fluentd"
- name: PASSWORD
  value: "redhat"
- name: CLIENT_AUTH
  value: none
- name: MACHINE_LEARNING
  value: "true"
`)

func loggingExternalLogStoresElasticsearch8HttpsUser_authConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch8HttpsUser_authConfigmapYaml, nil
}

func loggingExternalLogStoresElasticsearch8HttpsUser_authConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch8HttpsUser_authConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/8/https/user_auth/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresElasticsearch8HttpsUser_authDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    labels:
      app: ${NAME}
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        app: ${NAME}
    strategy:
      activeDeadlineSeconds: 21600
      resources: {}
      rollingParams:
        intervalSeconds: 1
        maxSurge: 25%
        maxUnavailable: 25%
        timeoutSeconds: 600
        updatePeriodSeconds: 1
      type: Recreate
    template:
      metadata:
        labels:
          app: ${NAME}
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - image: quay.io/openshifttest/elasticsearch@sha256:234e8ecfb6c1bdafffeb190c4a48e8e3c4b74b69d1b541b913dfbca29e952f63
          imagePullPolicy: IfNotPresent
          name: ${NAME}
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          ports:
          - containerPort: 9300
            protocol: TCP
          - containerPort: 9200
            protocol: TCP
          readinessProbe:
            exec:
              command:
              - sh
              - /usr/share/elasticsearch/add-user.sh
            failureThreshold: 3
            initialDelaySeconds: 5
            periodSeconds: 5
            successThreshold: 1
            timeoutSeconds: 10
          resources:
            limits:
              cpu: 1
              memory: 2Gi
            requests:
              cpu: 1
              memory: 2Gi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /usr/share/elasticsearch/config/elasticsearch.yml
            name: elasticsearch-config
            subPath: elasticsearch.yml
          - mountPath: /usr/share/elasticsearch/add-user.sh
            name: elasticsearch-config
            subPath: add-user.sh
          - mountPath: /usr/share/elasticsearch/config/secret
            name: certificates
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        serviceAccount: ${NAME}
        serviceAccountName: ${NAME}
        terminationGracePeriodSeconds: 30
        volumes:
        - configMap:
            defaultMode: 420
            name: ${NAME}
          name: elasticsearch-config
        - name: certificates
          secret:
            defaultMode: 420
            secretName: ${NAME}
parameters:
- name: NAME
  value: elasticsearch-server
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresElasticsearch8HttpsUser_authDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresElasticsearch8HttpsUser_authDeploymentYaml, nil
}

func loggingExternalLogStoresElasticsearch8HttpsUser_authDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresElasticsearch8HttpsUser_authDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/elasticsearch/8/https/user_auth/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresFluentdInsecureConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: v1
  data:
    fluent.conf: |
      <source>
        @type forward
        port  24224
      </source>

      <match kubernetes.var.log.pods.openshift-*_** kubernetes.var.log.pods.default_** kubernetes.var.log.pods.kube-*_**>
        @type file
        append true
        path /fluentd/log/infra-container.*.log
        symlink_path /fluentd/log/infra-container.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match journal.** system.var.log**>
        @type file
        append true
        path /fluentd/log/infra.*.log
        symlink_path /fluentd/log/infra.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match kubernetes.**>
        @type file
        append true
        path /fluentd/log/app.*.log
        symlink_path /fluentd/log/app.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match linux-audit.log** k8s-audit.log** openshift-audit.log** ovn-audit.log**>
        @type file
        append true
        path /fluentd/log/audit.*.log
        symlink_path /fluentd/log/audit.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match **>
        @type stdout
      </match>
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: fluentdserver
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresFluentdInsecureConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresFluentdInsecureConfigmapYaml, nil
}

func loggingExternalLogStoresFluentdInsecureConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresFluentdInsecureConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/fluentd/insecure/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresFluentdInsecureDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
    labels:
      provider: aosqe
      component: ${NAME}
      logging-infra: ${NAME}
  spec:
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        provider: aosqe
        component: ${NAME}
        logging-infra: ${NAME}
    strategy:
      type: Recreate
    template:
      metadata:
        labels:
          logging-infra: ${NAME}
          provider: aosqe
          component: ${NAME}
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - name: "fluentdserver"
          image: "quay.io/openshifttest/fluentd:1.2.2"
          imagePullPolicy: "IfNotPresent"
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          ports:
          - containerPort: 24224
            name: fluentdserver
          volumeMounts:
          - mountPath: /fluentd/etc
            name: config
            readOnly: true
        volumes:
        - configMap:
            defaultMode: 420
            name: ${NAME}
          name: config
parameters:
- name: NAME
  value: fluentdserver
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresFluentdInsecureDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresFluentdInsecureDeploymentYaml, nil
}

func loggingExternalLogStoresFluentdInsecureDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresFluentdInsecureDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/fluentd/insecure/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresFluentdInsecureHttpConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: fluentd-http-template
objects:
- apiVersion: v1
  data:
    fluent.conf: |
      <system>
        log_level info
      </system>
      <source>
        @type http
        port 24224
        bind 0.0.0.0
        body_size_limit 32m
        keepalive_timeout 10s
        add_http_headers true
        add_remote_addr true
        <parse>
          @type json
        </parse>
        @label @collector_logs
      </source>
      <source>
        @type http
        port 24224
        bind "::"
        body_size_limit 32m
        keepalive_timeout 10s
        add_http_headers true
        add_remote_addr true
        <parse>
          @type json
        </parse>
        @label @collector_logs
      </source>
      <label @collector_logs>
        <match logs.app>
          @type file
          append true
          path /fluentd/log/app.*.log
          symlink_path /fluentd/log/app.log
        </match>
        <match logs.infra>
          @type file
          append true
          path /fluentd/log/infra.*.log
          symlink_path /fluentd/log/infra.log
        </match>
        <match logs.audit>
          @type file
          append true
          path /fluentd/log/audit.*.log
          symlink_path /fluentd/log/audit.log
        </match>
      </label>
      <label @FLUENT_LOG>
        <match **>
      	  @type stdout
        </match>
      </label>
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: fluentdserver
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresFluentdInsecureHttpConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresFluentdInsecureHttpConfigmapYaml, nil
}

func loggingExternalLogStoresFluentdInsecureHttpConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresFluentdInsecureHttpConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/fluentd/insecure/http-configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresFluentdSecureCmMtlsShareYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: v1
  data:
    fluent.conf: |
      <source>
        @type forward
        port  24224
        <transport tls>
          ca_path /etc/fluentd/secrets/ca-bundle.crt
          cert_path /etc/fluentd/secrets/tls.crt
          private_key_path /etc/fluentd/secrets/tls.key
          client_cert_auth true
        </transport>
        <security>
          shared_key ${SHARED_KEY}
          self_hostname ${NAME}
        </security>
      </source>

      <match kubernetes.var.log.pods.openshift-*_** kubernetes.var.log.pods.default_** kubernetes.var.log.pods.kube-*_**>
        @type file
        append true
        path /fluentd/log/infra-container.*.log
        symlink_path /fluentd/log/infra-container.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match journal.** system.var.log**>
        @type file
        append true
        path /fluentd/log/infra.*.log
        symlink_path /fluentd/log/infra.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match kubernetes.**>
        @type file
        append true
        path /fluentd/log/app.*.log
        symlink_path /fluentd/log/app.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
       time_format       %Y%m%dT%H%M%S%z
      </match>
      <match linux-audit.log** k8s-audit.log** openshift-audit.log** ovn-audit.log**>
        @type file
        append true
        path /fluentd/log/audit.*.log
        symlink_path /fluentd/log/audit.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match **>
        @type stdout
      </match>
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: fluentdserver
- name: NAMESPACE
  value: openshift-logging
- name: SHARED_KEY
`)

func loggingExternalLogStoresFluentdSecureCmMtlsShareYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresFluentdSecureCmMtlsShareYaml, nil
}

func loggingExternalLogStoresFluentdSecureCmMtlsShareYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresFluentdSecureCmMtlsShareYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/fluentd/secure/cm-mtls-share.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresFluentdSecureCmMtlsYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: v1
  data:
    fluent.conf: |
      <source>
        @type forward
        port  24224
        <transport tls>
          ca_path /etc/fluentd/secrets/ca-bundle.crt
          cert_path /etc/fluentd/secrets/tls.crt
          private_key_path /etc/fluentd/secrets/tls.key
          client_cert_auth true
        </transport>
      </source>

      <match kubernetes.var.log.pods.openshift-*_** kubernetes.var.log.pods.default_** kubernetes.var.log.pods.kube-*_**>
        @type file
        append true
        path /fluentd/log/infra-container.*.log
        symlink_path /fluentd/log/infra-container.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match journal.** system.var.log**>
        @type file
        append true
        path /fluentd/log/infra.*.log
        symlink_path /fluentd/log/infra.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match kubernetes.**>
        @type file
        append true
        path /fluentd/log/app.*.log
        symlink_path /fluentd/log/app.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match linux-audit.log** k8s-audit.log** openshift-audit.log** ovn-audit.log**>
        @type file
        append true
        path /fluentd/log/audit.*.log
        symlink_path /fluentd/log/audit.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match **>
        @type stdout
      </match>
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: fluentdserver
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresFluentdSecureCmMtlsYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresFluentdSecureCmMtlsYaml, nil
}

func loggingExternalLogStoresFluentdSecureCmMtlsYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresFluentdSecureCmMtlsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/fluentd/secure/cm-mtls.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresFluentdSecureCmServerauthShareYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: v1
  data:
    fluent.conf: |
      <source>
        @type forward
        port  24224
        <transport tls>
          ca_cert_path /etc/fluentd/secrets/ca-bundle.crt
          ca_private_key_path /etc/fluentd/secrets/ca.key
        </transport>
        <security>
          shared_key ${SHARED_KEY}
          self_hostname ${NAME}
        </security>
      </source>

      <match kubernetes.var.log.pods.openshift-*_** kubernetes.var.log.pods.default_** kubernetes.var.log.pods.kube-*_**>
        @type file
        append true
        path /fluentd/log/infra-container.*.log
        symlink_path /fluentd/log/infra-container.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match journal.** system.var.log**>
        @type file
        append true
        path /fluentd/log/infra.*.log
        symlink_path /fluentd/log/infra.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match kubernetes.**>
        @type file
        append true
        path /fluentd/log/app.*.log
        symlink_path /fluentd/log/app.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match linux-audit.log** k8s-audit.log** openshift-audit.log** ovn-audit.log**>
        @type file
        append true
        path /fluentd/log/audit.*.log
        symlink_path /fluentd/log/audit.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match **>
        @type stdout
      </match>
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: fluentdserver
- name: NAMESPACE
  value: openshift-logging
- name: SHARED_KEY
`)

func loggingExternalLogStoresFluentdSecureCmServerauthShareYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresFluentdSecureCmServerauthShareYaml, nil
}

func loggingExternalLogStoresFluentdSecureCmServerauthShareYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresFluentdSecureCmServerauthShareYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/fluentd/secure/cm-serverauth-share.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresFluentdSecureCmServerauthYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: v1
  data:
    fluent.conf: |
      <source>
        @type forward
        port  24224
        <transport tls>
          ca_cert_path /etc/fluentd/secrets/ca-bundle.crt
          ca_private_key_path /etc/fluentd/secrets/ca.key
        </transport>
      </source>

      <match kubernetes.var.log.pods.openshift-*_** kubernetes.var.log.pods.default_** kubernetes.var.log.pods.kube-*_**>
        @type file
        append true
        path /fluentd/log/infra-container.*.log
        symlink_path /fluentd/log/infra-container.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match journal.** system.var.log**>
        @type file
        append true
        path /fluentd/log/infra.*.log
        symlink_path /fluentd/log/infra.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match kubernetes.**>
        @type file
        append true
        path /fluentd/log/app.*.log
        symlink_path /fluentd/log/app.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match linux-audit.log** k8s-audit.log** openshift-audit.log** ovn-audit.log**>
        @type file
        append true
        path /fluentd/log/audit.*.log
        symlink_path /fluentd/log/audit.log
        time_slice_format %Y%m%d
        time_slice_wait   1m
        time_format       %Y%m%dT%H%M%S%z
      </match>
      <match **>
        @type stdout
      </match>
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: fluentdserver
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresFluentdSecureCmServerauthYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresFluentdSecureCmServerauthYaml, nil
}

func loggingExternalLogStoresFluentdSecureCmServerauthYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresFluentdSecureCmServerauthYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/fluentd/secure/cm-serverauth.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresFluentdSecureDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: external-log-store-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
    labels:
      provider: aosqe
      component: ${NAME}
      logging-infra: ${NAME}
  spec:
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        provider: aosqe
        component: ${NAME}
        logging-infra: ${NAME}
    strategy:
      type: Recreate
    template:
      metadata:
        labels:
          logging-infra: ${NAME}
          provider: aosqe
          component: ${NAME}
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - name: "fluentdserver"
          image: "quay.io/openshifttest/fluentd:1.2.2"
          imagePullPolicy: "IfNotPresent"
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          ports:
          - containerPort: 24224
            name: fluentdserver
          volumeMounts:
          - mountPath: /fluentd/etc
            name: config
            readOnly: true
          - mountPath: /etc/fluentd/secrets
            name: certs
            readOnly: true
        volumes:
        - configMap:
            defaultMode: 420
            name: ${NAME}
          name: config
        - name: certs
          secret:
            defaultMode: 420
            secretName: ${NAME}
parameters:
- name: NAME
  value: fluentdserver
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresFluentdSecureDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresFluentdSecureDeploymentYaml, nil
}

func loggingExternalLogStoresFluentdSecureDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresFluentdSecureDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/fluentd/secure/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresFluentdSecureHttpCmMtlsYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: fluentd-http-template
objects:
- apiVersion: v1
  data:
    fluent.conf: |
      <system>
        log_level info
      </system>
      <source>
        @type http
        port 24224
        bind 0.0.0.0
        body_size_limit 32m
        keepalive_timeout 10s
        add_http_headers true
        add_remote_addr true
        <transport tls>
          ca_path /etc/fluentd/secrets/ca-bundle.crt
          cert_path /etc/fluentd/secrets/tls.crt
          private_key_path /etc/fluentd/secrets/tls.key
          client_cert_auth true
        </transport>
        <parse>
          @type json
        </parse>
        @label @collector_logs
      </source>
      <source>
        @type http
        port 24224
        bind "::"
        body_size_limit 32m
        keepalive_timeout 10s
        add_http_headers true
        add_remote_addr true
        <transport tls>
          ca_path /etc/fluentd/secrets/ca-bundle.crt
          cert_path /etc/fluentd/secrets/tls.crt
          private_key_path /etc/fluentd/secrets/tls.key
          client_cert_auth true
        </transport>
        <parse>
          @type json
        </parse>
        @label @collector_logs
      </source>
      <label @collector_logs>
        <match logs.app>
          @type file
          append true
          path /fluentd/log/app.*.log
          symlink_path /fluentd/log/app.log
        </match>
        <match logs.infra>
          @type file
          append true
          path /fluentd/log/infra.*.log
          symlink_path /fluentd/log/infra.log
        </match>
        <match logs.audit>
          @type file
          append true
          path /fluentd/log/audit.*.log
          symlink_path /fluentd/log/audit.log
        </match>
      </label>
      <label @FLUENT_LOG>
        <match **>
        	@type stdout
        </match>
      </label>
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: fluentdserver
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresFluentdSecureHttpCmMtlsYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresFluentdSecureHttpCmMtlsYaml, nil
}

func loggingExternalLogStoresFluentdSecureHttpCmMtlsYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresFluentdSecureHttpCmMtlsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/fluentd/secure/http-cm-mtls.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresFluentdSecureHttpCmServerauthYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: fluentd-http-template
objects:
- apiVersion: v1
  data:
    fluent.conf: |
      <system>
        log_level info
      </system>
      <source>
        @type http
        port 24224
        bind 0.0.0.0
        body_size_limit 32m
        keepalive_timeout 10s
        add_http_headers true
        add_remote_addr true
        <transport tls>
          ca_path /etc/fluentd/secrets/ca-bundle.crt
          cert_path /etc/fluentd/secrets/tls.crt
          private_key_path /etc/fluentd/secrets/tls.key
          client_cert_auth false
        </transport>
        <parse>
          @type json
        </parse>
        @label @collector_logs
      </source>
      <source>
        @type http
        port 24224
        bind "::"
        body_size_limit 32m
        keepalive_timeout 10s
        add_http_headers true
        add_remote_addr true
        <transport tls>
          ca_path /etc/fluentd/secrets/ca-bundle.crt
          cert_path /etc/fluentd/secrets/tls.crt
          private_key_path /etc/fluentd/secrets/tls.key
          client_cert_auth false
        </transport>
        <parse>
          @type json
        </parse>
        @label @collector_logs
      </source>
      <label @collector_logs>
        <match logs.app>
          @type file
          append true
          path /fluentd/log/app.*.log
          symlink_path /fluentd/log/app.log
        </match>
        <match logs.infra>
          @type file
          append true
          path /fluentd/log/infra.*.log
          symlink_path /fluentd/log/infra.log
        </match>
        <match logs.audit>
          @type file
          append true
          path /fluentd/log/audit.*.log
          symlink_path /fluentd/log/audit.log
        </match>
      </label>
      <label @FLUENT_LOG>
        <match **>
        	@type stdout
        </match>
      </label>
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: fluentdserver
- name: NAMESPACE
  value: openshift-logging
`)

func loggingExternalLogStoresFluentdSecureHttpCmServerauthYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresFluentdSecureHttpCmServerauthYaml, nil
}

func loggingExternalLogStoresFluentdSecureHttpCmServerauthYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresFluentdSecureHttpCmServerauthYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/fluentd/secure/http-cm-serverauth.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaAmqstreamsKafkaNoAuthClusterYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: kafka-template
objects:
- apiVersion: kafka.strimzi.io/v1beta2
  kind: Kafka
  metadata:
    name: ${NAME}
  spec:
    kafka:
      replicas: 1
      version: ${VERSION}
      resources:
        requests:
          cpu: '64m'
          memory: 256Mi
        limits:
          memory: 4Gi
          cpu: "1"
      jvmOptions:
        '-Xms': 256m
        '-Xmx': 256m
      config:
        log.cleaner.enable: true
        log.segment.bytes: 268435456
        log.cleanup.policy: delete
        transaction.state.log.replication.factor: 1
        log.retention.bytes: 1073741824
        transaction.state.log.min.isr: 1
        log.retention.hours: 1
        auto.create.topics.enable: false
        offsets.topic.replication.factor: 1
      listeners:
      - name: plain
        port: 9092
        type: internal
        tls: false
        configuration:
          useServiceDnsDomain: true
      - name: tls
        port: 9093
        type: internal
        tls: true
        authentication:
          type: tls
      storage:
        type: ephemeral
    zookeeper:
      replicas: 1
      resources:
        limits:
          cpu: '1'
          memory: 2Gi
        requests:
          cpu: '64m'
          memory: 256Mi
      storage:
      storage:
        type: ephemeral
    entityOperator:
      topicOperator:
        reconciliationIntervalSeconds: 90
      userOperator:
        reconciliationIntervalSeconds: 120
parameters:
- name: NAME
  value: "my-cluster"
- name: VERSION
  value: "3.9.0"
`)

func loggingExternalLogStoresKafkaAmqstreamsKafkaNoAuthClusterYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaAmqstreamsKafkaNoAuthClusterYaml, nil
}

func loggingExternalLogStoresKafkaAmqstreamsKafkaNoAuthClusterYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaAmqstreamsKafkaNoAuthClusterYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/amqstreams/kafka-no-auth-cluster.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaAmqstreamsKafkaNoAuthConsumerJobYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: kafka-template
objects:
- apiVersion: batch/v1
  kind: Job
  metadata:
    name: ${NAME}
  spec:
    ttlSecondsAfterFinished: 60
    template:
      spec:
        containers:
        - name: kakfa-consumer
          image: "registry.redhat.io/amq7/amq-streams-kafka-25-rhel7@sha256:e719f662bd4d6b8c54b1ee2e47c51f8d75a27a238a51d9ee38007187b3a627a4"
          command: ["bin/kafka-console-consumer.sh","--bootstrap-server", "${CLUSTER_NAME}-kafka-bootstrap:9092", "--topic", "${TOPIC_NAME}", "--from-beginning"]
        restartPolicy: Never
        nodeSelector:
          kubernetes.io/arch: amd64
          kubernetes.io/os: linux
parameters:
- name: NAME
  value: "topic-logging-app-consumer"
- name: CLUSTER_NAME
  value: "my-cluster"
- name: TOPIC_NAME
  value: "topic-logging-app"
`)

func loggingExternalLogStoresKafkaAmqstreamsKafkaNoAuthConsumerJobYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaAmqstreamsKafkaNoAuthConsumerJobYaml, nil
}

func loggingExternalLogStoresKafkaAmqstreamsKafkaNoAuthConsumerJobYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaAmqstreamsKafkaNoAuthConsumerJobYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/amqstreams/kafka-no-auth-consumer-job.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaAmqstreamsKafkaSaslClusterYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: kafka-template
objects:
- apiVersion: kafka.strimzi.io/v1beta2
  kind: Kafka
  metadata:
    name: ${NAME}
  spec:
    entityOperator:
      topicOperator:
        reconciliationIntervalSeconds: 90
      userOperator:
        reconciliationIntervalSeconds: 120
    kafka:
      authorization:
        type: simple
      config:
        log.cleaner.enable: true
        log.segment.bytes: 268435456
        log.cleanup.policy: delete
        transaction.state.log.replication.factor: 1
        log.retention.bytes: 1073741824
        transaction.state.log.min.isr: 1
        log.retention.hours: 1
        auto.create.topics.enable: false
        offsets.topic.replication.factor: 1
      jvmOptions:
        '-Xms': 256m
        '-Xmx': 256m
      listeners:
        - authentication:
            type: scram-sha-512
          name: external
          port: 9093
          tls: true
          type: route
        - authentication:
            type: scram-sha-512
          configuration:
            useServiceDnsDomain: true
          name: plain
          port: 9092
          tls: false
          type: internal
      replicas: 1
      resources:
        limits:
          cpu: '2'
          memory: 4Gi
        requests:
          cpu: '64m'
          memory: 256Mi
      storage:
        type: ephemeral
      version: ${VERSION}
    zookeeper:
      jvmOptions:
        '-Xms': 256m
        '-Xmx': 256m
      replicas: 1
      resources:
        limits:
          cpu: '1'
          memory: 2Gi
        requests:
          cpu: '64m'
          memory: 256Mi
      storage:
        type: ephemeral
parameters:
- name: NAME
  value: "my-cluster"
- name: VERSION
  value: "3.9.0"
`)

func loggingExternalLogStoresKafkaAmqstreamsKafkaSaslClusterYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaAmqstreamsKafkaSaslClusterYaml, nil
}

func loggingExternalLogStoresKafkaAmqstreamsKafkaSaslClusterYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaAmqstreamsKafkaSaslClusterYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/amqstreams/kafka-sasl-cluster.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaAmqstreamsKafkaSaslConsumerJobYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: kafka-consumer-job-template
objects:
- apiVersion: batch/v1
  kind: Job
  metadata:
    name: ${NAME}
  spec:
    ttlSecondsAfterFinished: 60
    template:
      spec:
        containers:
        - name: kafka-consumer
          image: registry.redhat.io/amq-streams/kafka-35-rhel8@sha256:fc900527fa19b35ec909c2a44e9e22ff0119934dcdf6e5da3665631d724a1bf4
          command: ["bin/kafka-console-consumer.sh","--bootstrap-server=${CLUSTER_NAME}-kafka-bootstrap:9092","--topic=${TOPIC_NAME}","--consumer.config=/opt/kafka/qeclient/client.property", "--from-beginning"]
          volumeMounts:
          - mountPath: /opt/kafka/qeclient
            name: kafka-config
          - mountPath: /opt/kafka/qep12
            name: cluster-ca
        restartPolicy: Never
        volumes:
        - configMap:
            defaultMode: 420
            name: ${CLIENT_CONFIGMAP_NAME}
          name: kafka-config
        - name: cluster-ca
          secret:
            defaultMode: 288
            secretName: ${CA_SECRET_NAME}
parameters:
- name: NAME
  value: "topic-logging-consumer"
- name: CLUSTER_NAME
  value: "my-cluster"
- name: TOPIC_NAME
  value: "topic-logging"
- name: CLIENT_CONFIGMAP_NAME
  value: ""
- name: CA_SECRET_NAME
  value: ""
`)

func loggingExternalLogStoresKafkaAmqstreamsKafkaSaslConsumerJobYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaAmqstreamsKafkaSaslConsumerJobYaml, nil
}

func loggingExternalLogStoresKafkaAmqstreamsKafkaSaslConsumerJobYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaAmqstreamsKafkaSaslConsumerJobYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/amqstreams/kafka-sasl-consumer-job.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaAmqstreamsKafkaSaslConsumersConfigYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: kafka-client-config-template
objects:
- apiVersion: v1
  data:
    client.property: |-
      security.protocol=SASL_PLAINTEXT
      sasl.mechanism=SCRAM-SHA-512
      sasl.jaas.config=org.apache.kafka.common.security.scram.ScramLoginModule required username="${USER}" password="${PASSWORD}";
      bootstrap.servers=${KAFKA_NAME}-kafka-bootstrap:9092
      ssl.truststore.location=/opt/kafka/qep12/ca.p12
      ssl.truststore.password=${TRUSTSTORE_PASSWORD}
      ssl.truststore.type=PKCS12 
      group.id=my-group
  kind: ConfigMap
  metadata:
    name: ${NAME}
parameters:
- name: NAME
  value: "client-property"
- name: USER
  value: "my-user"
- name: PASSWORD
  value: ""
- name: TRUSTSTORE_PASSWORD
  value: ""
- name: KAFKA_NAME
  value: "my-cluster"
`)

func loggingExternalLogStoresKafkaAmqstreamsKafkaSaslConsumersConfigYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaAmqstreamsKafkaSaslConsumersConfigYaml, nil
}

func loggingExternalLogStoresKafkaAmqstreamsKafkaSaslConsumersConfigYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaAmqstreamsKafkaSaslConsumersConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/amqstreams/kafka-sasl-consumers-config.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaAmqstreamsKafkaSaslUserYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: kafka-user-template
objects:
- apiVersion: kafka.strimzi.io/v1beta2
  kind: KafkaUser
  metadata:
    name: ${NAME}
    labels:
      strimzi.io/cluster: ${KAFKA_NAME}
  spec:
    authentication:
      type: scram-sha-512
    authorization:
      acls:
        - host: '*'
          operations:
            - Read
            - Describe
            - Write
            - Create
          resource:
            name: ${TOPIC_PREFIX}
            patternType: prefix
            type: topic
        - host: '*'
          operations:
            - Read
          resource:
            name: my-group
            patternType: literal
            type: group
      type: simple
parameters:
- name: NAME
  value: "my-user"
- name:  KAFKA_NAME
  value: "my-cluster"
- name: TOPIC_PREFIX
  value: "topic-logging"
`)

func loggingExternalLogStoresKafkaAmqstreamsKafkaSaslUserYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaAmqstreamsKafkaSaslUserYaml, nil
}

func loggingExternalLogStoresKafkaAmqstreamsKafkaSaslUserYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaAmqstreamsKafkaSaslUserYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/amqstreams/kafka-sasl-user.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaAmqstreamsKafkaTopicYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: kafka-template
objects:
- apiVersion: kafka.strimzi.io/v1beta1
  kind: KafkaTopic
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
    labels:
      strimzi.io/cluster: ${CLUSTER_NAME}
  spec:
    partitions: 1
    replicas: 1
    config:
      retention.ms: 300000
      segment.bytes: 1073741824
parameters:
- name: CLUSTER_NAME
  value: "my-cluster"
- name: NAME
  value: "logging-topic-all"
- name: NAMESPACE
  value: "amq-aosqe"
`)

func loggingExternalLogStoresKafkaAmqstreamsKafkaTopicYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaAmqstreamsKafkaTopicYaml, nil
}

func loggingExternalLogStoresKafkaAmqstreamsKafkaTopicYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaAmqstreamsKafkaTopicYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/amqstreams/kafka-topic.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaCert_generationSh = []byte(`#!/usr/bin/env bash
set -eou pipefail

WORKING_DIR=${1:-/tmp/_working_dir}
NAMESPACE=${2:-openshift-logging}
CA_PATH=$WORKING_DIR/ca
cn_name="aosqeca"
CLIENT_CA_PATH="${WORKING_DIR}/client"
CLUSTER_CA_PATH="${WORKING_DIR}/cluster"
REGENERATE_NEEDED=0

function init_cert_files() {
if [ ! -d $CA_PATH ]; then
    mkdir $CA_PATH
fi
if [ ! -d $CLIENT_CA_PATH ]; then
    mkdir ${CLIENT_CA_PATH}
fi
if [ ! -d $CLUSTER_CA_PATH ]; then
    mkdir ${CLUSTER_CA_PATH}
fi
}

function create_root_ca() {
#fqdn=my-cluster-kafka-bootstrap.amq-aosqe.svc
#1)  generate root ca key and and self cert
openssl req -x509 -new -newkey rsa:4096 -keyout $CA_PATH/root_ca.key -nodes -days 1825 -out $CA_PATH/ca_bundle.crt -subj "/CN=aosqeroot"   -passin pass:aosqe2021 -passout pass:aosqe2021

# Create trustore jks
/usr/lib/jvm/jre-openjdk/bin/keytool -import -file $CA_PATH/ca_bundle.crt -keystore $CA_PATH/ca_bundle.jks  --srcstorepass aosqe2021 --deststorepass aosqe2021 -noprompt || exit 1
}

# Create Client CSR
function create_client_sign() {
cat <<EOF >$CLIENT_CA_PATH/client_csr.conf
[ req ]
default_bits = 4096
prompt = no
encrypt_key = yes
default_md = sha512
distinguished_name = dn
req_extensions = server_ext

[ dn ]
CN = "aosqeclient"

[ server_ext ]
basicConstraints        = CA:false
extendedKeyUsage        = serverAuth,clientAuth
subjectAltName = DNS.1:*.cluster.local,DNS.2:*.svc,DNS.3:*.pod
EOF
openssl req -new -out $CLIENT_CA_PATH/client.csr -newkey rsa:4096 -keyout $CLIENT_CA_PATH/client.key -config $CLIENT_CA_PATH/client_csr.conf -nodes


# Sign client ca by intermediate ca
cat <<EOF >$CLIENT_CA_PATH/client_sign.conf
# Simple Signing CA

# The [default] section contains global constants that can be referred to from
# the entire configuration file. It may also hold settings pertaining to more
# than one openssl command.

[ default ]
dir                     = $CLIENT_CA_PATH             # Top dir

# The remainder of the configuration file is used by the openssl ca command.
# The CA section defines the locations of CA assets, as well as the policies
# applying to the CA.

[ ca ]
default_ca              = signing_ca            # The default CA section

[ signing_ca ]
certificate             = $CA_PATH/ca_bundle.crt       # The CA cert
private_key             = $CA_PATH/root_ca.key # CA private key
new_certs_dir           = $CA_PATH/           # Certificate archive
serial                  = $CA_PATH/root_ca.serial.txt # Serial number file
crlnumber               = $CA_PATH/root_ca.crl.srl # CRL number file
database                = $CA_PATH/root_ca.db # Index file
unique_subject          = no                    # Require unique subject
default_days            = 730                   # How long to certify for
default_md              = sha512                # MD to use
policy                  = any_pol             # Default naming policy
email_in_dn             = no                    # Add email to cert DN
preserve                = no                    # Keep passed DN ordering
name_opt                = ca_default            # Subject DN display options
cert_opt                = ca_default            # Certificate display options
copy_extensions         = copy                  # Copy extensions from CSR
x509_extensions         = server_ext             # Default cert extensions
default_crl_days        = 7                     # How long before next CRL

# Naming policies control which parts of a DN end up in the certificate and
# under what circumstances certification should be denied.

[ any_pol ]
domainComponent         = optional
countryName             = optional
stateOrProvinceName     = optional
localityName            = optional
organizationName        = optional
organizationalUnitName  = optional
commonName              = optional
emailAddress            = optional

# Certificate extensions define what types of certificates the CA is able to
# create.

[ server_ext ]
basicConstraints        = CA:false
extendedKeyUsage        = serverAuth,clientAuth

[ ca_reqext ]
basicConstraints        = CA:false


# CRL extensions exist solely to point to the CA certificate that has issued
# the CRL.
EOF

touch $CA_PATH/root_ca.db
if [ ! -f $CA_PATH/root_ca.serial.txt ] ; then
    echo "01">$CA_PATH/root_ca.serial.txt
fi
openssl ca -in $CLIENT_CA_PATH/client.csr -notext -out $CLIENT_CA_PATH/client.crt -config $CLIENT_CA_PATH/client_sign.conf -batch

# Create Client keystone
openssl pkcs12 -export -in $CLIENT_CA_PATH/client.crt -inkey $CLIENT_CA_PATH/client.key -out $CLIENT_CA_PATH/client.pkcs12  -passin pass:aosqe2021 -passout pass:aosqe2021
/usr/lib/jvm/jre-openjdk/bin/keytool -importkeystore -srckeystore $CLIENT_CA_PATH/client.pkcs12 -srcstoretype PKCS12 -destkeystore $CLIENT_CA_PATH/client.jks -deststoretype JKS --srcstorepass aosqe2021 --deststorepass aosqe2021 -noprompt
}

# Create cluster csr
# https://support.dnsimple.com/articles/ssl-certificate-names/
function create_cluster_sign() {
cat <<EOF >$CLUSTER_CA_PATH/cluster_csr.conf
[ req ]
default_bits = 4096
prompt = no
encrypt_key = yes
default_md = sha512
distinguished_name = dn
req_extensions = server_ext

[ dn ]
CN = "aosqecluster"

[ server_ext ]
basicConstraints        = CA:false
extendedKeyUsage        = serverAuth,clientAuth
subjectAltName = DNS.1:kafka.${NAMESPACE}.svc.cluster.local,DNS.2:kafka.${NAMESPACE}.svc,DNS.3:kafka-0.kafka.${NAMESPACE}.svc.cluster.local,DNS.4: kafka-0.kafka.${NAMESPACE}.svc, DNS.5: kafka, DNS.6: kakfa-0
EOF
openssl req -new -out $CLUSTER_CA_PATH/cluster.csr -newkey rsa:4096 -keyout $CLUSTER_CA_PATH/cluster.key -config $CLUSTER_CA_PATH/cluster_csr.conf -nodes


cat <<EOF >$CLUSTER_CA_PATH/cluster_sign.conf
# Simple Signing CA

# The [default] section contains global constants that can be referred to from
# the entire configuration file. It may also hold settings pertaining to more
# than one openssl command.

[ default ]
dir                     = $CA_PATH             # Top dir

# The next part of the configuration file is used by the openssl req command.
# It defines the CA's key pair, its DN, and the desired extensions for the CA
# certificate.

[ req ]
default_bits            = 4096                  # RSA key size
encrypt_key             = yes                   # Protect private key
default_md              = sha512                # MD to use
utf8                    = yes                   # Input is UTF-8
string_mask             = utf8only              # Emit UTF-8 strings
prompt                  = no                    # Don't prompt for DN
distinguished_name      = ca_dn                 # DN section
req_extensions          = ca_reqext             # Desired extensions

[ ca_dn ]
commonName              = "aosqeintermediate"

[ ca_reqext ]
basicConstraints        = CA:false

# The remainder of the configuration file is used by the openssl ca command.
# The CA section defines the locations of CA assets, as well as the policies
# applying to the CA.

[ ca ]
default_ca              = signing_ca            # The default CA section

[ signing_ca ]
certificate             = $CA_PATH/ca_bundle.crt       # The CA cert
private_key             = $CA_PATH/root_ca.key # CA private key
new_certs_dir           = $CA_PATH/           # Certificate archive
serial                  = $CA_PATH/root_ca.serial.txt # Serial number file
crlnumber               = $CA_PATH/root_ca.crl.srl # CRL number file
database                = $CA_PATH/root_ca.db # Index file
unique_subject          = no                    # Require unique subject
default_days            = 730                   # How long to certify for
default_md              = sha512                # MD to use
policy                  = any_pol             # Default naming policy
email_in_dn             = no                    # Add email to cert DN
preserve                = no                    # Keep passed DN ordering
name_opt                = ca_default            # Subject DN display options
cert_opt                = ca_default            # Certificate display options
copy_extensions         = copy                  # Copy extensions from CSR
#x509_extensions         = server_ext             # Default cert extensions
default_crl_days        = 7                     # How long before next CRL

# Naming policies control which parts of a DN end up in the certificate and
# under what circumstances certification should be denied.

[ match_pol ]
domainComponent         = match                 # Must match 'simple.org'
organizationName        = match                 # Must match 'Simple Inc'
organizationalUnitName  = optional              # Included if present
commonName              = supplied              # Must be present

[ any_pol ]
domainComponent         = optional
countryName             = optional
stateOrProvinceName     = optional
localityName            = optional
organizationName        = optional
organizationalUnitName  = optional
commonName              = optional
emailAddress            = optional

# Certificate extensions define what types of certificates the CA is able to
# create.

[ server_ext ]
basicConstraints        = CA:false
extendedKeyUsage        = serverAuth,clientAuth
subjectAltName = DNS.1:kafka.${NAMESPACE}.svc.cluster.local,DNS.2:kafka.${NAMESPACE}.svc,DNS.3:kafka-0.kafka.${NAMESPACE}.svc.cluster.local,DNS.4: kafka-0.kafka.${NAMESPACE}.svc, DNS.5: kafka, DNS.6: kakfa-0
EOF

touch $CA_PATH/root_ca.db
if [ ! -f $CA_PATH/root_ca.serial.txt ] ; then
    echo "01">$CA_PATH/root_ca.serial.txt
fi
openssl ca -in $CLUSTER_CA_PATH/cluster.csr -notext -out $CLUSTER_CA_PATH/cluster.crt -config $CLUSTER_CA_PATH/cluster_sign.conf -batch

#Create keystone
openssl pkcs12 -export -in $CLUSTER_CA_PATH/cluster.crt -inkey $CLUSTER_CA_PATH/cluster.key -out $CLUSTER_CA_PATH/cluster.pkcs12  -passin pass:aosqe2021 -passout pass:aosqe2021
/usr/lib/jvm/jre-openjdk/bin/keytool -importkeystore -srckeystore $CLUSTER_CA_PATH/cluster.pkcs12 -srcstoretype PKCS12 -destkeystore $CLUSTER_CA_PATH/cluster.jks -deststoretype JKS  --srcstorepass aosqe2021 --deststorepass aosqe2021 -noprompt
}

init_cert_files
create_root_ca
create_client_sign
create_cluster_sign
`)

func loggingExternalLogStoresKafkaCert_generationShBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaCert_generationSh, nil
}

func loggingExternalLogStoresKafkaCert_generationSh() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaCert_generationShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/cert_generation.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaKafkaRbacYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: kafkaserver-template
objects:
- apiVersion: rbac.authorization.k8s.io/v1
  kind: ClusterRole
  metadata:
    name: ${NAME}
  rules:
  - apiGroups:
    - ""
    resources:
    - nodes
    verbs:
    - get
  - apiGroups:
    - ""
    resources:
    - pods
    verbs:
    - get
    - create
    - update
    - patch
    - delete
- apiVersion: rbac.authorization.k8s.io/v1
  kind: ClusterRoleBinding
  metadata:
    name: ${NAME}
  roleRef:
    apiGroup: rbac.authorization.k8s.io
    kind: ClusterRole
    name: ${NAME}
  subjects:
  - kind: ServiceAccount
    name: default
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: "kafka-node-reader"
- name: NAMESPACE
  value: "openshift-logging"
`)

func loggingExternalLogStoresKafkaKafkaRbacYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaKafkaRbacYaml, nil
}

func loggingExternalLogStoresKafkaKafkaRbacYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaKafkaRbacYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/kafka-rbac.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaKafkaSvcYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: kafkaserver-template
objects:
- apiVersion: v1
  kind: Service
  metadata:
    labels:
      logging-infra: support
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    ports:
    - name: plaintext
      port: 9092
      protocol: TCP
      targetPort: 9092
    - name: saslplaintext
      port: 9093
      protocol: TCP
      targetPort: 9093
    - name: slasssl
      port: 9094
      protocol: TCP
      targetPort: 9093
    selector:
      component: kafka
      provider: openshift
    sessionAffinity: None
    type: ClusterIP
parameters:
- name: NAME
  value: "kafka"
- name: NAMESPACE
  value: "openshift-logging"
`)

func loggingExternalLogStoresKafkaKafkaSvcYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaKafkaSvcYaml, nil
}

func loggingExternalLogStoresKafkaKafkaSvcYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaKafkaSvcYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/kafka-svc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaPlaintextSslConsumerConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: consumer-configmap-template
objects:
- apiVersion: v1
  data:
    client.properties: |
      bootstrap.servers=kafka:9093
      #group.id=test-consumer-group
      security.protocol=SSL
      ssl.truststore.location=/etc/kafkacert/ca-bundle.jks
      ssl.truststore.password=aosqe2021
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: "kafka-client"
- name: NAMESPACE
  value: "openshift-logging"
`)

func loggingExternalLogStoresKafkaPlaintextSslConsumerConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaPlaintextSslConsumerConfigmapYaml, nil
}

func loggingExternalLogStoresKafkaPlaintextSslConsumerConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaPlaintextSslConsumerConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/plaintext-ssl/consumer-configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaPlaintextSslKafkaConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: consumer-configmap-template
objects:
- apiVersion: v1
  data:
    init.sh: |
      #!/bin/bash
      set -e
      cp /etc/kafka-configmap/log4j.properties /etc/kafka/

      KAFKA_BROKER_ID=${HOSTNAME##*-}
      SEDS=("s/#init#broker.id=#init#/broker.id=$KAFKA_BROKER_ID/")
      LABELS="kafka-broker-id=$KAFKA_BROKER_ID"
      ANNOTATIONS=""

      hash kubectl 2>/dev/null || {
        SEDS+=("s/#init#broker.rack=#init#/#init#broker.rack=# kubectl not found in path/")
      } && {
        ZONE=$(kubectl get node "$NODE_NAME" -o=go-template='{{index .metadata.labels "failure-domain.beta.kubernetes.io/zone"}}')
        if [ "x$ZONE" == "x<no value>" ]; then
          SEDS+=("s/#init#broker.rack=#init#/#init#broker.rack=# zone label not found for node $NODE_NAME/")
        else
          SEDS+=("s/#init#broker.rack=#init#/broker.rack=$ZONE/")
          LABELS="$LABELS kafka-broker-rack=$ZONE"
        fi

        [ -z "$ADVERTISE_ADDR" ] && echo "ADVERTISE_ADDR is empty, will advertise detected DNS name"
        OUTSIDE_HOST=$(kubectl get node "$NODE_NAME" -o jsonpath='{.status.addresses[?(@.type=="InternalIP")].address}')
        OUTSIDE_PORT=$((32400 + ${KAFKA_BROKER_ID}))
        SEDS+=("s|#init#advertised.listeners=PLAINTEXT://#init#|advertised.listeners=PLAINTEXT://${ADVERTISE_ADDR}:9092,OUTSIDE://${OUTSIDE_HOST}:${OUTSIDE_PORT}|")
        ANNOTATIONS="$ANNOTATIONS kafka-listener-outside-host=$OUTSIDE_HOST kafka-listener-outside-port=$OUTSIDE_PORT"

        if [ ! -z "$LABELS" ]; then
          kubectl -n $POD_NAMESPACE label pod $POD_NAME $LABELS || echo "Failed to label $POD_NAMESPACE.$POD_NAME - RBAC issue?"
        fi
        if [ ! -z "$ANNOTATIONS" ]; then
          kubectl -n $POD_NAMESPACE annotate pod $POD_NAME $ANNOTATIONS || echo "Failed to annotate $POD_NAMESPACE.$POD_NAME - RBAC issue?"
        fi
      }
      printf '%s\n' "${SEDS[@]}" | sed -f - /etc/kafka-configmap/server.properties > /etc/kafka/server.properties.tmp
      [ $? -eq 0 ] && mv /etc/kafka/server.properties.tmp /etc/kafka/server.properties
    log4j.properties: |
      # Unspecified loggers and loggers with additivity=true output to server.log and stdout
      # Note that INFO only applies to unspecified loggers, the log level of the child logger is used otherwise
      log4j.rootLogger=INFO, stdout

      log4j.appender.stdout=org.apache.log4j.ConsoleAppender
      log4j.appender.stdout.layout=org.apache.log4j.PatternLayout
      log4j.appender.stdout.layout.ConversionPattern=[%d] %p %m (%c)%n

      log4j.appender.kafkaAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.kafkaAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.kafkaAppender.File=${kafka.logs.dir}/server.log
      log4j.appender.kafkaAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.kafkaAppender.layout.ConversionPattern=[%d] %p %m (%c)%n

      log4j.appender.stateChangeAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.stateChangeAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.stateChangeAppender.File=${kafka.logs.dir}/state-change.log
      log4j.appender.stateChangeAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.stateChangeAppender.layout.ConversionPattern=[%d] %p %m (%c)%n

      log4j.appender.requestAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.requestAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.requestAppender.File=${kafka.logs.dir}/kafka-request.log
      log4j.appender.requestAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.requestAppender.layout.ConversionPattern=[%d] %p %m (%c)%n

      log4j.appender.cleanerAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.cleanerAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.cleanerAppender.File=${kafka.logs.dir}/log-cleaner.log
      log4j.appender.cleanerAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.cleanerAppender.layout.ConversionPattern=[%d] %p %m (%c)%n

      log4j.appender.controllerAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.controllerAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.controllerAppender.File=${kafka.logs.dir}/controller.log
      log4j.appender.controllerAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.controllerAppender.layout.ConversionPattern=[%d] %p %m (%c)%n

      log4j.appender.authorizerAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.authorizerAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.authorizerAppender.File=${kafka.logs.dir}/kafka-authorizer.log
      log4j.appender.authorizerAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.authorizerAppender.layout.ConversionPattern=[%d] %p %m (%c)%n

      # Change the two lines below to adjust ZK client logging
      log4j.logger.org.I0Itec.zkclient.ZkClient=INFO
      log4j.logger.org.apache.zookeeper=INFO

      # Change the two lines below to adjust the general broker logging level (output to server.log and stdout)
      log4j.logger.kafka=INFO
      log4j.logger.org.apache.kafka=INFO

      # Change to DEBUG or TRACE to enable request logging
      log4j.logger.kafka.request.logger=WARN, requestAppender
      log4j.additivity.kafka.request.logger=false

      # Uncomment the lines below and change log4j.logger.kafka.network.RequestChannel$ to TRACE for additional output
      # related to the handling of requests
      #log4j.logger.kafka.network.Processor=TRACE, requestAppender
      #log4j.logger.kafka.server.KafkaApis=TRACE, requestAppender
      #log4j.additivity.kafka.server.KafkaApis=false
      log4j.logger.kafka.network.RequestChannel$=WARN, requestAppender
      log4j.additivity.kafka.network.RequestChannel$=false

      log4j.logger.kafka.controller=TRACE, controllerAppender
      log4j.additivity.kafka.controller=false

      log4j.logger.kafka.log.LogCleaner=INFO, cleanerAppender
      log4j.additivity.kafka.log.LogCleaner=false

      log4j.logger.state.change.logger=TRACE, stateChangeAppender
      log4j.additivity.state.change.logger=false

      # Change to DEBUG to enable audit log for the authorizer
      log4j.logger.kafka.authorizer.logger=WARN, authorizerAppender
      log4j.additivity.kafka.authorizer.logger=false
    server.properties: |
      #init#broker.id=#init#
      listeners=PLAINTEXT://:9092,SSL://:9093
      ssl.client.auth=requested
      ssl.keystore.location=/etc/kafkacert/cluster.jks
      ssl.keystore.password=aosqe2021
      ssl.truststore.location=/etc/kafkacert/ca_bundle.jks
      ssl.truststore.password=aosqe2021
      security.inter.broker.protocol=PLAINTEXT
      num.network.threads=3
      num.io.threads=8
      message.max.bytes=314572800
      socket.send.buffer.bytes=102400
      socket.receive.buffer.bytes=102400
      socket.request.max.bytes=104857600
      socket.request.max.bytes=314572800
      log.dirs=/tmp/kafka-logs
      num.partitions=1
      num.recovery.threads.per.data.dir=1
      offsets.topic.replication.factor=1
      transaction.state.log.replication.factor=1
      transaction.state.log.min.isr=1
      log.retention.hours=2
      log.segment.bytes=1073741824
      log.retention.check.interval.ms=300000
      zookeeper.connect=zookeeper:2181
      zookeeper.connection.timeout.ms=18000
      group.initial.rebalance.delay.ms=0
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: "kafka"
- name: NAMESPACE
  value: "openshift-logging"
`)

func loggingExternalLogStoresKafkaPlaintextSslKafkaConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaPlaintextSslKafkaConfigmapYaml, nil
}

func loggingExternalLogStoresKafkaPlaintextSslKafkaConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaPlaintextSslKafkaConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/plaintext-ssl/kafka-configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaPlaintextSslKafkaConsumerDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: kafkaserver-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "1"
    labels:
      component: kafka-consumer
      logging-infra: kafka
      provider: openshift
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        component: kafka-consumer
        logging-infra: kafka
        provider: openshift
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        creationTimestamp: null
        labels:
          component: kafka-consumer
          logging-infra: kafka
          provider: openshift
        name: kafka-consumer
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - command:
          - /bin/bash
          - -ce
          - /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9093 --topic clo-topic --from-beginning  --consumer.config /etc/kafka-config/client.properties
          image: quay.io/openshifttest/kafka@sha256:2411662d89dd5700e1fe49aa8be1219843948da90cfe51a1c7a49bcef9d22dab
          imagePullPolicy: IfNotPresent
          name: kafka-consumer
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /shared
            name: shared
          - mountPath: /etc/kafka-config
            name: kafka-client
          - mountPath: /etc/kafkacert
            name: kafkacert
          env:
          - name: KAFKA_OPTS
            value: -Djava.security.auth.login.config=/etc/kafka-configmap/kafka_client_jaas.conf
        dnsPolicy: ClusterFirst
        nodeSelector:
          kubernetes.io/arch: amd64
          kubernetes.io/os: linux
        restartPolicy: Always
        schedulerName: default-scheduler
        terminationGracePeriodSeconds: 30
        volumes:
        - emptyDir: {}
          name: shared
        - configMap:
            defaultMode: 420
            name: ${CM_NAME}
          name: kafka-client
        - secret:
            defaultMode: 420
            secretName: ${SECRETNAME}
          name: kafkacert
parameters:
- name: NAME
  value: "kafka-consumer-plaintext-ssl"
- name: NAMESPACE
  value: "openshift-logging"
- name: CM_NAME
  value: "kafka-client"
- name: SECRETNAME
  value: "kafka-client-cert"
`)

func loggingExternalLogStoresKafkaPlaintextSslKafkaConsumerDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaPlaintextSslKafkaConsumerDeploymentYaml, nil
}

func loggingExternalLogStoresKafkaPlaintextSslKafkaConsumerDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaPlaintextSslKafkaConsumerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/plaintext-ssl/kafka-consumer-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaPlaintextSslKafkaStatefulsetYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: kafkaserver-template
objects:
- apiVersion: apps/v1
  kind: StatefulSet
  metadata:
    labels:
      app: kafka
      component: kafka
      provider: openshift
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    podManagementPolicy: Parallel
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        app: kafka
    serviceName: ${SERVICENAME}
    template:
      metadata:
        creationTimestamp: null
        labels:
          app: kafka
          component: kafka
          provider: openshift
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - command:
          - /opt/kafka/bin/kafka-server-start.sh
          - /etc/kafka/server.properties
          env:
          - name: CLASSPATH
            value: /opt/kafka/libs/extensions/*
          - name: KAFKA_LOG4J_OPTS
            value: -Dlog4j.configuration=file:/etc/kafka/log4j.properties
          - name: JMX_PORT
            value: "5555"
          image: quay.io/openshifttest/kafka@sha256:2411662d89dd5700e1fe49aa8be1219843948da90cfe51a1c7a49bcef9d22dab
          imagePullPolicy: IfNotPresent
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          lifecycle:
            preStop:
              exec:
                command:
                - sh
                - -ce
                - kill -s TERM 1; while $(kill -0 1 2>/dev/null); do sleep 1; done
          name: broker
          ports:
          - containerPort: 9092
            name: inside
            protocol: TCP
          - containerPort: 9093
            name: ssl
            protocol: TCP
          - containerPort: 9094
            name: outide
            protocol: TCP
          - containerPort: 5555
            name: jmx
            protocol: TCP
          readinessProbe:
            failureThreshold: 3
            periodSeconds: 10
            successThreshold: 1
            tcpSocket:
              port: 9092
            timeoutSeconds: 1
          resources:
            limits:
              memory: 1Gi
            requests:
              cpu: 250m
              memory: 500Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /etc/kafka-configmap
            name: brokerconfig
          - mountPath: /etc/kafka
            name: config
          - mountPath: /etc/kafkacert
            name: kafkacert
          - mountPath: /opt/kafka/logs
            name: brokerlogs
          - mountPath: /opt/kafka/libs/extensions
            name: extensions
          - mountPath: /var/lib/kafka/data
            name: data
        dnsPolicy: ClusterFirst
        initContainers:
        - command:
          - /bin/bash
          - /etc/kafka-configmap/init.sh
          env:
          - name: NODE_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: spec.nodeName
          - name: POD_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.name
          - name: POD_NAMESPACE
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.namespace
          - name: ADVERTISE_ADDR
            value: kafka
          image: quay.io/openshifttest/kafka-initutils@sha256:e73ff7a44b43b85b53849c0459ba32e704540852b885a5c78af9753f86a49d68
          imagePullPolicy: IfNotPresent
          name: init-config
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /etc/kafka-configmap
            name: brokerconfig
          - mountPath: /etc/kafka
            name: config
          - mountPath: /opt/kafka/libs/extensions
            name: extensions
        restartPolicy: Always
        nodeSelector:
          kubernetes.io/arch: amd64
          kubernetes.io/os: linux
        schedulerName: default-scheduler
        terminationGracePeriodSeconds: 30
        volumes:
        - configMap:
            defaultMode: 420
            name: ${CM_NAME}
          name: brokerconfig
        - secret:
            defaultMode: 420
            secretName: ${SECRETNAME}
          name: kafkacert
        - emptyDir: {}
          name: brokerlogs
        - emptyDir: {}
          name: config
        - emptyDir: {}
          name: extensions
        - emptyDir: {}
          name: data
    updateStrategy:
      type: RollingUpdate
parameters:
- name: NAME
  value: "kafka"
- name: NAMESPACE
  value: "openshift-logging"
- name: SERVICENAME
  value: "kafka"
- name: CM_NAME
  value: "kafka"
- name: SECRETNAME
  value: "kafka-cluster-cert"
`)

func loggingExternalLogStoresKafkaPlaintextSslKafkaStatefulsetYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaPlaintextSslKafkaStatefulsetYaml, nil
}

func loggingExternalLogStoresKafkaPlaintextSslKafkaStatefulsetYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaPlaintextSslKafkaStatefulsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/plaintext-ssl/kafka-statefulset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaSaslPlaintextConsumerConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: consumer-configmap-template
objects:
- apiVersion: v1
  data:
    client.properties: |
      bootstrap.servers=kafka:9092
      #group.id=test-consumer-group
      sasl.mechanism=PLAIN
      security.protocol=SASL_PLAINTEXT
      sasl.jaas.config=org.apache.kafka.common.security.plain.PlainLoginModule required \
         username="admin" \
         password="admin-secret";
    kafka_client_jaas.conf: |
      KafkaClient {
         org.apache.kafka.common.security.plain.PlainLoginModule required
         username="admin"
         password="admin-secret";
      };
    sasl-consumer.properties: |
      #export KAFKA_OPTS="-Djava.security.auth.login.config=/etc/kafka-configmap/kafka_client_jaas.conf"
      #/opt/kafka/bin/kafka-console-producer.sh --broker-list kafka:9092 --producer.config=/etc/kafka-config/sasl-producer.properties  --topic  clo-topic
      bootstrap.servers=kafka:9092
      compression.type=none
      ### SECURITY ######
      security.protocol=SASL_PLANTEXT
      sasl.mechanism=PLAIN
      sasl.jaas.config=org.apache.kafka.common.security.plain.PlainLoginModule required username="admin" password="admin-secret";
      ssl.truststore.location=/etc/kafkacert/ca-bundle.jks
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: "kafka-client"
- name: NAMESPACE
  value: "openshift-logging"
`)

func loggingExternalLogStoresKafkaSaslPlaintextConsumerConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaSaslPlaintextConsumerConfigmapYaml, nil
}

func loggingExternalLogStoresKafkaSaslPlaintextConsumerConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaSaslPlaintextConsumerConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/sasl-plaintext/consumer-configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaSaslPlaintextKafkaConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: consumer-configmap-template
objects:
- apiVersion: v1
  data:
    init.sh: |
      #!/bin/bash
      set -e
      cp /etc/kafka-configmap/log4j.properties /etc/kafka/
      KAFKA_BROKER_ID=${HOSTNAME##*-}
      SEDS=("s/#init#broker.id=#init#/broker.id=$KAFKA_BROKER_ID/")
      LABELS="kafka-broker-id=$KAFKA_BROKER_ID"
      ANNOTATIONS=""

      hash kubectl 2>/dev/null || {
        SEDS+=("s/#init#broker.rack=#init#/#init#broker.rack=# kubectl not found in path/")
      } && {
        ZONE=$(kubectl get node "$NODE_NAME" -o=go-template='{{index .metadata.labels "failure-domain.beta.kubernetes.io/zone"}}')
        if [ "x$ZONE" == "x<no value>" ]; then
          SEDS+=("s/#init#broker.rack=#init#/#init#broker.rack=# zone label not found for node $NODE_NAME/")
        else
          SEDS+=("s/#init#broker.rack=#init#/broker.rack=$ZONE/")
          LABELS="$LABELS kafka-broker-rack=$ZONE"
        fi

        [ -z "$ADVERTISE_ADDR" ] && echo "ADVERTISE_ADDR is empty, will advertise detected DNS name"
        OUTSIDE_HOST=$(kubectl get node "$NODE_NAME" -o jsonpath='{.status.addresses[?(@.type=="InternalIP")].address}')
        OUTSIDE_PORT=$((32400 + ${KAFKA_BROKER_ID}))
        SEDS+=("s|#init#advertised.listeners=PLAINTEXT://#init#|advertised.listeners=PLAINTEXT://${ADVERTISE_ADDR}:9092,SASL_PLAINTEXT://${ADVERTISE_ADDR}:9093|")
        ANNOTATIONS="$ANNOTATIONS kafka-listener-outside-host=$OUTSIDE_HOST kafka-listener-outside-port=$OUTSIDE_PORT"

        if [ ! -z "$LABELS" ]; then
          kubectl -n $POD_NAMESPACE label pod $POD_NAME $LABELS || echo "Failed to label $POD_NAMESPACE.$POD_NAME - RBAC issue?"
        fi
        if [ ! -z "$ANNOTATIONS" ]; then
          kubectl -n $POD_NAMESPACE annotate pod $POD_NAME $ANNOTATIONS || echo "Failed to annotate $POD_NAMESPACE.$POD_NAME - RBAC issue?"
        fi
      }
      printf '%s\n' "${SEDS[@]}" | sed -f - /etc/kafka-configmap/server.properties > /etc/kafka/server.properties.tmp
      [ $? -eq 0 ] && mv /etc/kafka/server.properties.tmp /etc/kafka/server.properties
    kafka_server_jaas.conf: |
      KafkaServer {
         org.apache.kafka.common.security.plain.PlainLoginModule required
         serviceName="kafka"
         username="admin"
         password="admin-secret"
         user_admin="admin-secret"
         user_alice="alice-secret";
      };
    log4j.properties: |
      log4j.rootLogger=INFO, stdout
      log4j.appender.stdout=org.apache.log4j.ConsoleAppender
      log4j.appender.stdout.layout=org.apache.log4j.PatternLayout
      log4j.appender.stdout.layout.ConversionPattern=[%d] %p %m (%c)%n
      log4j.appender.kafkaAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.kafkaAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.kafkaAppender.File=${kafka.logs.dir}/server.log
      log4j.appender.kafkaAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.kafkaAppender.layout.ConversionPattern=[%d] %p %m (%c)%n
      log4j.appender.stateChangeAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.stateChangeAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.stateChangeAppender.File=${kafka.logs.dir}/state-change.log
      log4j.appender.stateChangeAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.stateChangeAppender.layout.ConversionPattern=[%d] %p %m (%c)%n
      log4j.appender.requestAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.requestAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.requestAppender.File=${kafka.logs.dir}/kafka-request.log
      log4j.appender.requestAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.requestAppender.layout.ConversionPattern=[%d] %p %m (%c)%n
      log4j.appender.cleanerAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.cleanerAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.cleanerAppender.File=${kafka.logs.dir}/log-cleaner.log
      log4j.appender.cleanerAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.cleanerAppender.layout.ConversionPattern=[%d] %p %m (%c)%n
      log4j.appender.controllerAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.controllerAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.controllerAppender.File=${kafka.logs.dir}/controller.log
      log4j.appender.controllerAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.controllerAppender.layout.ConversionPattern=[%d] %p %m (%c)%n
      log4j.appender.authorizerAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.authorizerAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.authorizerAppender.File=${kafka.logs.dir}/kafka-authorizer.log
      log4j.appender.authorizerAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.authorizerAppender.layout.ConversionPattern=[%d] %p %m (%c)%n
      log4j.logger.org.I0Itec.zkclient.ZkClient=INFO
      log4j.logger.org.apache.zookeeper=INFO
      log4j.logger.kafka=INFO
      log4j.logger.org.apache.kafka=INFO
      log4j.logger.kafka.request.logger=WARN, requestAppender
      log4j.additivity.kafka.request.logger=false
      log4j.logger.kafka.network.RequestChannel$=WARN, requestAppender
      log4j.additivity.kafka.network.RequestChannel$=false
      log4j.logger.kafka.controller=TRACE, controllerAppender
      log4j.additivity.kafka.controller=false
      log4j.logger.kafka.log.LogCleaner=INFO, cleanerAppender
      log4j.additivity.kafka.log.LogCleaner=false
      log4j.logger.state.change.logger=TRACE, stateChangeAppender
      log4j.additivity.state.change.logger=false
      log4j.logger.kafka.authorizer.logger=WARN, authorizerAppender
      log4j.additivity.kafka.authorizer.logger=false
    server.properties: |
      #https://docs.confluent.io/platform/current/kafka/authentication_sasl/authentication_sasl_plain.html
      #init#broker.id=#init#
      ssl.client.auth=none
      sasl.enabled.mechanisms=PLAIN
      sasl.mechanism.inter.broker.protocol=PLAIN
      security.inter.broker.protocol=SASL_PLAINTEXT
      listeners=SASL_PLAINTEXT://:9092
      security.protocol=SASL_PLAINTEXT
      authorizer.class.name=kafka.security.authorizer.AclAuthorizer
      super.users=User:admin
      allow.everyone.if.no.acl.found=true
      num.partitions=1
      num.network.threads=3
      num.io.threads=8
      num.recovery.threads.per.data.dir=1
      message.max.bytes=314572800
      socket.send.buffer.bytes=102400
      socket.receive.buffer.bytes=102400
      socket.request.max.bytes=104857600
      socket.request.max.bytes=314572800
      log.dirs=/tmp/kafka-logs
      offsets.topic.replication.factor=1
      transaction.state.log.replication.factor=1
      transaction.state.log.min.isr=1
      log.retention.hours=2
      log.segment.bytes=1073741824
      log.retention.check.interval.ms=300000
      zookeeper.connect=zookeeper:2181
      zookeeper.connection.timeout.ms=18000
      group.initial.rebalance.delay.ms=0
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: "kafka"
- name: NAMESPACE
  value: "openshift-logging"
`)

func loggingExternalLogStoresKafkaSaslPlaintextKafkaConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaSaslPlaintextKafkaConfigmapYaml, nil
}

func loggingExternalLogStoresKafkaSaslPlaintextKafkaConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaSaslPlaintextKafkaConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/sasl-plaintext/kafka-configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaSaslPlaintextKafkaConsumerDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: kafkaserver-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "1"
    labels:
      component: kafka-consumer
      logging-infra: kafka
      provider: openshift
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    progressDeadlineSeconds: 601
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        component: kafka-consumer
        logging-infra: kafka
        provider: openshift
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        creationTimestamp: null
        labels:
          component: kafka-consumer
          logging-infra: kafka
          provider: openshift
        name: kafka-consumer
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - command:
          - /bin/bash
          - -ce
          - /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic clo-topic --from-beginning  --consumer.config /etc/kafka-config/client.properties
          image: quay.io/openshifttest/kafka@sha256:2411662d89dd5700e1fe49aa8be1219843948da90cfe51a1c7a49bcef9d22dab
          imagePullPolicy: IfNotPresent
          name: kafka-consumer
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /shared
            name: shared
          - mountPath: /etc/kafka-config
            name: kafka-client
          env:
          - name: KAFKA_OPTS
            value: -Djava.security.auth.login.config=/etc/kafka-configmap/kafka_client_jaas.conf
        dnsPolicy: ClusterFirst
        nodeSelector:
          kubernetes.io/arch: amd64
          kubernetes.io/os: linux
        restartPolicy: Always
        schedulerName: default-scheduler
        terminationGracePeriodSeconds: 30
        volumes:
        - emptyDir: {}
          name: shared
        - configMap:
            defaultMode: 420
            name: ${CM_NAME}
          name: kafka-client
parameters:
- name: NAME
  value: "kafka-consumer-sasl-plaintext"
- name: NAMESPACE
  value: "openshift-logging"
- name: CM_NAME
  value: "kafka-client"
`)

func loggingExternalLogStoresKafkaSaslPlaintextKafkaConsumerDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaSaslPlaintextKafkaConsumerDeploymentYaml, nil
}

func loggingExternalLogStoresKafkaSaslPlaintextKafkaConsumerDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaSaslPlaintextKafkaConsumerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/sasl-plaintext/kafka-consumer-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaSaslPlaintextKafkaStatefulsetYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: kafkaserver-template
objects:
- apiVersion: apps/v1
  kind: StatefulSet
  metadata:
    labels:
      app: kafka
      component: kafka
      provider: openshift
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    podManagementPolicy: Parallel
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        app: kafka
    serviceName: ${SERVICENAME}
    template:
      metadata:
        creationTimestamp: null
        labels:
          app: kafka
          component: kafka
          provider: openshift
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - command:
          - /opt/kafka/bin/kafka-server-start.sh
          - /etc/kafka/server.properties
          env:
          - name: CLASSPATH
            value: /opt/kafka/libs/extensions/*
          - name: KAFKA_LOG4J_OPTS
            value: -Dlog4j.configuration=file:/etc/kafka/log4j.properties
          - name: KAFKA_OPTS
            value: -Djava.security.auth.login.config=/etc/kafka-configmap/kafka_server_jaas.conf
          - name: JMX_PORT
            value: "5555"
          image: quay.io/openshifttest/kafka@sha256:2411662d89dd5700e1fe49aa8be1219843948da90cfe51a1c7a49bcef9d22dab
          imagePullPolicy: IfNotPresent
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          lifecycle:
            preStop:
              exec:
                command:
                - sh
                - -ce
                - kill -s TERM 1; while $(kill -0 1 2>/dev/null); do sleep 1; done
          name: broker
          ports:
          - containerPort: 9092
            name: inside
            protocol: TCP
          - containerPort: 9093
            name: ssl
            protocol: TCP
          - containerPort: 9094
            name: outide
            protocol: TCP
          - containerPort: 5555
            name: jmx
            protocol: TCP
          readinessProbe:
            failureThreshold: 3
            periodSeconds: 10
            successThreshold: 1
            tcpSocket:
              port: 9092
            timeoutSeconds: 1
          resources:
            limits:
              memory: 1Gi
            requests:
              cpu: 250m
              memory: 500Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /etc/kafka-configmap
            name: brokerconfig
          - mountPath: /etc/kafka
            name: config
          - mountPath: /etc/kafkacert
            name: kafkacert
          - mountPath: /opt/kafka/logs
            name: brokerlogs
          - mountPath: /opt/kafka/libs/extensions
            name: extensions
          - mountPath: /var/lib/kafka/data
            name: data
        dnsPolicy: ClusterFirst
        initContainers:
        - command:
          - /bin/bash
          - /etc/kafka-configmap/init.sh
          env:
          - name: NODE_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: spec.nodeName
          - name: POD_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.name
          - name: POD_NAMESPACE
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.namespace
          - name: ADVERTISE_ADDR
            value: kafka
          image: quay.io/openshifttest/kafka-initutils@sha256:e73ff7a44b43b85b53849c0459ba32e704540852b885a5c78af9753f86a49d68
          imagePullPolicy: IfNotPresent
          name: init-config
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /etc/kafka-configmap
            name: brokerconfig
          - mountPath: /etc/kafka
            name: config
          - mountPath: /opt/kafka/libs/extensions
            name: extensions
        restartPolicy: Always
        nodeSelector:
          kubernetes.io/arch: amd64
          kubernetes.io/os: linux
        schedulerName: default-scheduler
        terminationGracePeriodSeconds: 30
        volumes:
        - configMap:
            defaultMode: 420
            name: ${CM_NAME}
          name: brokerconfig
        - emptyDir: {}
          name: kafkacert
        - emptyDir: {}
          name: brokerlogs
        - emptyDir: {}
          name: config
        - emptyDir: {}
          name: extensions
        - emptyDir: {}
          name: data
    updateStrategy:
      type: RollingUpdate
parameters:
- name: NAME
  value: "kafka"
- name: NAMESPACE
  value: "openshift-logging"
- name: SERVICENAME
  value: "kafka"
- name: CM_NAME
  value: "kafka"
`)

func loggingExternalLogStoresKafkaSaslPlaintextKafkaStatefulsetYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaSaslPlaintextKafkaStatefulsetYaml, nil
}

func loggingExternalLogStoresKafkaSaslPlaintextKafkaStatefulsetYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaSaslPlaintextKafkaStatefulsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/sasl-plaintext/kafka-statefulset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaSaslSslConsumerConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: consumer-configmap-template
objects:
- apiVersion: v1
  data:
    client.properties: |
      bootstrap.servers=kafka:9093
      #group.id=test-consumer-group
      sasl.mechanism=PLAIN
      security.protocol=SASL_SSL
      sasl.jaas.config=org.apache.kafka.common.security.plain.PlainLoginModule required \
         username="admin" \
         password="admin-secret";
      ssl.truststore.location=/etc/kafkacert/ca-bundle.jks
      ssl.truststore.password=aosqe2021
    kafka_client_jaas.conf: |
      KafkaClient {
         org.apache.kafka.common.security.plain.PlainLoginModule required
         username="admin"
         password="admin-secret";
      };
    ssl-consumer.properties: |
      #export KAFKA_OPTS="-Djava.security.auth.login.config=/etc/kafka-configmap/kafka_client_jaas.conf"
      #/opt/kafka/bin/kafka-console-producer.sh --broker-list kafka:9093 --producer.config=/etc/kafka-config/ssl-producer.properties  --topic  clo-topic
      bootstrap.servers=kafka:9093
      compression.type=none
      ### SECURITY ######
      security.protocol=SASL_SSL
      sasl.mechanism=PLAIN
      sasl.jaas.config=org.apache.kafka.common.security.plain.PlainLoginModule required username="admin" password="admin-secret";
      ssl.truststore.location=/etc/kafkacert/ca-bundle.jks
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: "kafka-client"
- name: NAMESPACE
  value: "openshift-logging"
`)

func loggingExternalLogStoresKafkaSaslSslConsumerConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaSaslSslConsumerConfigmapYaml, nil
}

func loggingExternalLogStoresKafkaSaslSslConsumerConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaSaslSslConsumerConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/sasl-ssl/consumer-configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaSaslSslKafkaConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: consumer-configmap-template
objects:
- apiVersion: v1
  data:
    init.sh: |
      #!/bin/bash
      set -e
      cp /etc/kafka-configmap/log4j.properties /etc/kafka/
      KAFKA_BROKER_ID=${HOSTNAME##*-}
      SEDS=("s/#init#broker.id=#init#/broker.id=$KAFKA_BROKER_ID/")
      LABELS="kafka-broker-id=$KAFKA_BROKER_ID"
      ANNOTATIONS=""

      hash kubectl 2>/dev/null || {
        SEDS+=("s/#init#broker.rack=#init#/#init#broker.rack=# kubectl not found in path/")
      } && {
        ZONE=$(kubectl get node "$NODE_NAME" -o=go-template='{{index .metadata.labels "failure-domain.beta.kubernetes.io/zone"}}')
        if [ "x$ZONE" == "x<no value>" ]; then
          SEDS+=("s/#init#broker.rack=#init#/#init#broker.rack=# zone label not found for node $NODE_NAME/")
        else
          SEDS+=("s/#init#broker.rack=#init#/broker.rack=$ZONE/")
          LABELS="$LABELS kafka-broker-rack=$ZONE"
        fi

        [ -z "$ADVERTISE_ADDR" ] && echo "ADVERTISE_ADDR is empty, will advertise detected DNS name"
        OUTSIDE_HOST=$(kubectl get node "$NODE_NAME" -o jsonpath='{.status.addresses[?(@.type=="InternalIP")].address}')
        OUTSIDE_PORT=$((32400 + ${KAFKA_BROKER_ID}))
        SEDS+=("s|#init#advertised.listeners=PLAINTEXT://#init#|advertised.listeners=PLAINTEXT://${ADVERTISE_ADDR}:9092,SASL_SSL://${ADVERTISE_ADDR}:9093|")
        ANNOTATIONS="$ANNOTATIONS kafka-listener-outside-host=$OUTSIDE_HOST kafka-listener-outside-port=$OUTSIDE_PORT"

        if [ ! -z "$LABELS" ]; then
          kubectl -n $POD_NAMESPACE label pod $POD_NAME $LABELS || echo "Failed to label $POD_NAMESPACE.$POD_NAME - RBAC issue?"
        fi
        if [ ! -z "$ANNOTATIONS" ]; then
          kubectl -n $POD_NAMESPACE annotate pod $POD_NAME $ANNOTATIONS || echo "Failed to annotate $POD_NAMESPACE.$POD_NAME - RBAC issue?"
        fi
      }
      printf '%s\n' "${SEDS[@]}" | sed -f - /etc/kafka-configmap/server.properties > /etc/kafka/server.properties.tmp
      [ $? -eq 0 ] && mv /etc/kafka/server.properties.tmp /etc/kafka/server.properties
    kafka_server_jaas.conf: |
      KafkaServer {
         org.apache.kafka.common.security.plain.PlainLoginModule required
         serviceName="kafka"
         username="admin"
         password="admin-secret"
         user_admin="admin-secret"
         user_alice="alice-secret";
      };
    log4j.properties: |
      log4j.rootLogger=INFO, stdout
      log4j.appender.stdout=org.apache.log4j.ConsoleAppender
      log4j.appender.stdout.layout=org.apache.log4j.PatternLayout
      log4j.appender.stdout.layout.ConversionPattern=[%d] %p %m (%c)%n
      log4j.appender.kafkaAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.kafkaAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.kafkaAppender.File=${kafka.logs.dir}/server.log
      log4j.appender.kafkaAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.kafkaAppender.layout.ConversionPattern=[%d] %p %m (%c)%n
      log4j.appender.stateChangeAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.stateChangeAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.stateChangeAppender.File=${kafka.logs.dir}/state-change.log
      log4j.appender.stateChangeAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.stateChangeAppender.layout.ConversionPattern=[%d] %p %m (%c)%n
      log4j.appender.requestAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.requestAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.requestAppender.File=${kafka.logs.dir}/kafka-request.log
      log4j.appender.requestAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.requestAppender.layout.ConversionPattern=[%d] %p %m (%c)%n
      log4j.appender.cleanerAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.cleanerAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.cleanerAppender.File=${kafka.logs.dir}/log-cleaner.log
      log4j.appender.cleanerAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.cleanerAppender.layout.ConversionPattern=[%d] %p %m (%c)%n
      log4j.appender.controllerAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.controllerAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.controllerAppender.File=${kafka.logs.dir}/controller.log
      log4j.appender.controllerAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.controllerAppender.layout.ConversionPattern=[%d] %p %m (%c)%n
      log4j.appender.authorizerAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.authorizerAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.authorizerAppender.File=${kafka.logs.dir}/kafka-authorizer.log
      log4j.appender.authorizerAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.authorizerAppender.layout.ConversionPattern=[%d] %p %m (%c)%n
      log4j.logger.org.I0Itec.zkclient.ZkClient=INFO
      log4j.logger.org.apache.zookeeper=INFO
      log4j.logger.kafka=INFO
      log4j.logger.org.apache.kafka=INFO
      log4j.logger.kafka.request.logger=WARN, requestAppender
      log4j.additivity.kafka.request.logger=false
      log4j.logger.kafka.network.RequestChannel$=WARN, requestAppender
      log4j.additivity.kafka.network.RequestChannel$=false
      log4j.logger.kafka.controller=TRACE, controllerAppender
      log4j.additivity.kafka.controller=false
      log4j.logger.kafka.log.LogCleaner=INFO, cleanerAppender
      log4j.additivity.kafka.log.LogCleaner=false
      log4j.logger.state.change.logger=TRACE, stateChangeAppender
      log4j.additivity.state.change.logger=false
      log4j.logger.kafka.authorizer.logger=WARN, authorizerAppender
      log4j.additivity.kafka.authorizer.logger=false
    server.properties: |
      #https://docs.confluent.io/platform/current/kafka/authentication_sasl/authentication_sasl_plain.html
      #init#broker.id=#init#
      ssl.client.auth=requested
      ssl.keystore.location=/etc/kafkacert/cluster.jks
      ssl.keystore.password=aosqe2021
      ssl.truststore.location=/etc/kafkacert/ca_bundle.jks
      ssl.truststore.password=aosqe2021
      sasl.enabled.mechanisms=PLAIN
      sasl.mechanism.inter.broker.protocol=PLAIN
      security.inter.broker.protocol=PLAINTEXT
      listeners=PLAINTEXT://:9092,SASL_SSL://:9093
      #init#advertised.listeners=PLAINTEXT://#init#
      security.protocol=SASL_SSL
      authorizer.class.name=kafka.security.authorizer.AclAuthorizer
      super.users=User:admin
      allow.everyone.if.no.acl.found=true
      num.partitions=1
      num.network.threads=3
      num.io.threads=8
      num.recovery.threads.per.data.dir=1
      message.max.bytes=314572800
      socket.send.buffer.bytes=102400
      socket.receive.buffer.bytes=102400
      socket.request.max.bytes=104857600
      socket.request.max.bytes=314572800
      log.dirs=/tmp/kafka-logs
      offsets.topic.replication.factor=1
      transaction.state.log.replication.factor=1
      transaction.state.log.min.isr=1
      log.retention.hours=2
      log.segment.bytes=1073741824
      log.retention.check.interval.ms=300000
      zookeeper.connect=zookeeper:2181
      zookeeper.connection.timeout.ms=18000
      group.initial.rebalance.delay.ms=0

  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: "kafka"
- name: NAMESPACE
  value: "openshift-logging"
`)

func loggingExternalLogStoresKafkaSaslSslKafkaConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaSaslSslKafkaConfigmapYaml, nil
}

func loggingExternalLogStoresKafkaSaslSslKafkaConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaSaslSslKafkaConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/sasl-ssl/kafka-configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaSaslSslKafkaConsumerDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: kafkaserver-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "1"
    labels:
      component: kafka-consumer
      logging-infra: kafka
      provider: openshift
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        component: kafka-consumer
        logging-infra: kafka
        provider: openshift
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        creationTimestamp: null
        labels:
          component: kafka-consumer
          logging-infra: kafka
          provider: openshift
        name: kafka-consumer
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - command:
          - /bin/bash
          - -ce
          - /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9093 --topic clo-topic --from-beginning  --consumer.config /etc/kafka-config/client.properties
          image: quay.io/openshifttest/kafka@sha256:2411662d89dd5700e1fe49aa8be1219843948da90cfe51a1c7a49bcef9d22dab
          imagePullPolicy: IfNotPresent
          name: kafka-consumer
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /shared
            name: shared
          - mountPath: /etc/kafka-config
            name: kafka-client
          - mountPath: /etc/kafkacert
            name: kafkacert
          env:
          - name: KAFKA_OPTS
            value: -Djava.security.auth.login.config=/etc/kafka-configmap/kafka_client_jaas.conf
        dnsPolicy: ClusterFirst
        nodeSelector:
          kubernetes.io/arch: amd64
          kubernetes.io/os: linux
        restartPolicy: Always
        schedulerName: default-scheduler
        terminationGracePeriodSeconds: 30
        volumes:
        - emptyDir: {}
          name: shared
        - configMap:
            defaultMode: 420
            name: ${CM_NAME}
          name: kafka-client
        - secret:
            defaultMode: 420
            secretName: ${SECRETNAME}
          name: kafkacert
parameters:
- name: NAME
  value: "kafka-consumer-sasl-ssl"
- name: NAMESPACE
  value: "openshift-logging"
- name: CM_NAME
  value: "kafka-client"
- name: SECRETNAME
  value: "kafka-client-cert"
`)

func loggingExternalLogStoresKafkaSaslSslKafkaConsumerDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaSaslSslKafkaConsumerDeploymentYaml, nil
}

func loggingExternalLogStoresKafkaSaslSslKafkaConsumerDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaSaslSslKafkaConsumerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/sasl-ssl/kafka-consumer-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaSaslSslKafkaStatefulsetYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: kafkaserver-template
objects:
- apiVersion: apps/v1
  kind: StatefulSet
  metadata:
    labels:
      app: kafka
      component: kafka
      provider: openshift
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    podManagementPolicy: Parallel
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        app: kafka
    serviceName: ${SERVICENAME}
    template:
      metadata:
        creationTimestamp: null
        labels:
          app: kafka
          component: kafka
          provider: openshift
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - command:
          - /opt/kafka/bin/kafka-server-start.sh
          - /etc/kafka/server.properties
          env:
          - name: CLASSPATH
            value: /opt/kafka/libs/extensions/*
          - name: KAFKA_LOG4J_OPTS
            value: -Dlog4j.configuration=file:/etc/kafka/log4j.properties
          - name: KAFKA_OPTS
            value: -Djava.security.auth.login.config=/etc/kafka-configmap/kafka_server_jaas.conf
          - name: JMX_PORT
            value: "5555"
          image: quay.io/openshifttest/kafka@sha256:2411662d89dd5700e1fe49aa8be1219843948da90cfe51a1c7a49bcef9d22dab
          imagePullPolicy: IfNotPresent
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          lifecycle:
            preStop:
              exec:
                command:
                - sh
                - -ce
                - kill -s TERM 1; while $(kill -0 1 2>/dev/null); do sleep 1; done
          name: broker
          ports:
          - containerPort: 9092
            name: inside
            protocol: TCP
          - containerPort: 9093
            name: ssl
            protocol: TCP
          - containerPort: 9094
            name: outide
            protocol: TCP
          - containerPort: 5555
            name: jmx
            protocol: TCP
          readinessProbe:
            failureThreshold: 3
            periodSeconds: 10
            successThreshold: 1
            tcpSocket:
              port: 9092
            timeoutSeconds: 1
          resources:
            limits:
              memory: 1Gi
            requests:
              cpu: 250m
              memory: 500Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /etc/kafka-configmap
            name: brokerconfig
          - mountPath: /etc/kafka
            name: config
          - mountPath: /etc/kafkacert
            name: kafkacert
          - mountPath: /opt/kafka/logs
            name: brokerlogs
          - mountPath: /opt/kafka/libs/extensions
            name: extensions
          - mountPath: /var/lib/kafka/data
            name: data
        dnsPolicy: ClusterFirst
        initContainers:
        - command:
          - /bin/bash
          - /etc/kafka-configmap/init.sh
          env:
          - name: NODE_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: spec.nodeName
          - name: POD_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.name
          - name: POD_NAMESPACE
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.namespace
          - name: ADVERTISE_ADDR
            value: kafka
          image: quay.io/openshifttest/kafka-initutils@sha256:e73ff7a44b43b85b53849c0459ba32e704540852b885a5c78af9753f86a49d68
          imagePullPolicy: IfNotPresent
          name: init-config
          nodeSelector:
            kubernetes.io/arch: amd64
            kubernetes.io/os: linux
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /etc/kafka-configmap
            name: brokerconfig
          - mountPath: /etc/kafka
            name: config
          - mountPath: /opt/kafka/libs/extensions
            name: extensions
        restartPolicy: Always
        schedulerName: default-scheduler
        terminationGracePeriodSeconds: 30
        volumes:
        - configMap:
            defaultMode: 420
            name: ${CM_NAME}
          name: brokerconfig
        - secret:
            defaultMode: 420
            secretName: ${SECRETNAME}
          name: kafkacert
        - emptyDir: {}
          name: brokerlogs
        - emptyDir: {}
          name: config
        - emptyDir: {}
          name: extensions
        - emptyDir: {}
          name: data
    updateStrategy:
      type: RollingUpdate
parameters:
- name: NAME
  value: "kafka"
- name: NAMESPACE
  value: "openshift-logging"
- name: SERVICENAME
  value: "kafka"
- name: CM_NAME
  value: "kafka"
- name: SECRETNAME
  value: "kafka-cluster-cert"
`)

func loggingExternalLogStoresKafkaSaslSslKafkaStatefulsetYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaSaslSslKafkaStatefulsetYaml, nil
}

func loggingExternalLogStoresKafkaSaslSslKafkaStatefulsetYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaSaslSslKafkaStatefulsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/sasl-ssl/kafka-statefulset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaZookeeperConfigmapSslYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: zookeeper-configmap-template
objects:
- apiVersion: v1
  data:
    init.sh: |
      #!/bin/bash
      set -e
      [ -d /var/lib/zookeeper/data ] || mkdir /var/lib/zookeeper/data
      [ -z "$ID_OFFSET" ] && ID_OFFSET=1
      export ZOOKEEPER_SERVER_ID=$((${HOSTNAME##*-} + $ID_OFFSET))
      echo "${ZOOKEEPER_SERVER_ID:-1}" | tee /var/lib/zookeeper/data/myid
      cp -Lur /etc/kafka-configmap/* /etc/kafka/
    log4j.properties: |
      log4j.rootLogger=INFO, stdout
      log4j.appender.stdout=org.apache.log4j.ConsoleAppender
      log4j.appender.stdout.layout=org.apache.log4j.PatternLayout
      log4j.appender.stdout.layout.ConversionPattern=[%d] %p %m (%c)%n
      # Suppress connection log messages, three lines per livenessProbe execution
      log4j.logger.org.apache.zookeeper.server.NIOServerCnxnFactory=WARN
      log4j.logger.org.apache.zookeeper.server.NIOServerCnxn=WARN
    zookeeper.properties: |
      4lw.commands.whitelist=ruok
      tickTime=2000
      dataDir=/var/lib/zookeeper/data
      dataLogDir=/var/lib/zookeeper/log
      clientPort=2181
      authProvider.sasl=org.apache.zookeeper.server.auth.SASLAuthenticationProvider
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: "zookeeper"
- name: NAMESPACE
  value: "openshift-logging"
`)

func loggingExternalLogStoresKafkaZookeeperConfigmapSslYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaZookeeperConfigmapSslYaml, nil
}

func loggingExternalLogStoresKafkaZookeeperConfigmapSslYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaZookeeperConfigmapSslYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/zookeeper/configmap-ssl.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaZookeeperConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: zookeeper-configmap-template
objects:
- apiVersion: v1
  data:
    init.sh: |
      #!/bin/bash
      set -e
      [ -d /var/lib/zookeeper/data ] || mkdir /var/lib/zookeeper/data
      [ -z "$ID_OFFSET" ] && ID_OFFSET=1
      export ZOOKEEPER_SERVER_ID=$((${HOSTNAME##*-} + $ID_OFFSET))
      echo "${ZOOKEEPER_SERVER_ID:-1}" | tee /var/lib/zookeeper/data/myid
      cp -Lur /etc/kafka-configmap/* /etc/kafka/
    log4j.properties: |
      log4j.rootLogger=INFO, stdout
      log4j.appender.stdout=org.apache.log4j.ConsoleAppender
      log4j.appender.stdout.layout=org.apache.log4j.PatternLayout
      log4j.appender.stdout.layout.ConversionPattern=[%d] %p %m (%c)%n
      # Suppress connection log messages, three lines per livenessProbe execution
      log4j.logger.org.apache.zookeeper.server.NIOServerCnxnFactory=WARN
      log4j.logger.org.apache.zookeeper.server.NIOServerCnxn=WARN
    zookeeper.properties: |
      4lw.commands.whitelist=ruok
      tickTime=2000
      dataDir=/var/lib/zookeeper/data
      dataLogDir=/var/lib/zookeeper/log
      clientPort=2181
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
parameters:
- name: NAME
  value: "zookeeper"
- name: NAMESPACE
  value: "openshift-logging"
`)

func loggingExternalLogStoresKafkaZookeeperConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaZookeeperConfigmapYaml, nil
}

func loggingExternalLogStoresKafkaZookeeperConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaZookeeperConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/zookeeper/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaZookeeperZookeeperStatefulsetYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: kafkaserver-template
objects:
- apiVersion: apps/v1
  kind: StatefulSet
  metadata:
    labels:
      app: zookeeper
      component: zookeeper
      provider: openshift
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    podManagementPolicy: Parallel
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        app: zookeeper
    serviceName: ${SERVICENAME}
    template:
      metadata:
        creationTimestamp: null
        labels:
          app: zookeeper
          component: zookeeper
          provider: openshift
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - command:
          - /opt/kafka/bin/zookeeper-server-start.sh
          - /etc/kafka/zookeeper.properties
          env:
          - name: KAFKA_LOG4J_OPTS
            value: -Dlog4j.configuration=file:/etc/kafka/log4j.properties
          image: quay.io/openshifttest/kafka@sha256:2411662d89dd5700e1fe49aa8be1219843948da90cfe51a1c7a49bcef9d22dab
          imagePullPolicy: IfNotPresent
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          lifecycle:
            preStop:
              exec:
                command:
                - sh
                - -ce
                - kill -s TERM 1; while $(kill -0 1 2>/dev/null); do sleep 1; done
          name: zookeeper
          ports:
          - containerPort: 2181
            name: client
            protocol: TCP
          - containerPort: 2888
            name: peer
            protocol: TCP
          - containerPort: 3888
            name: leader-election
            protocol: TCP
          resources:
            limits:
              memory: 120Mi
            requests:
              cpu: 10m
              memory: 100Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /etc/kafka
            name: config
          - mountPath: /opt/kafka/logs
            name: zookeeperlogs
          - mountPath: /var/lib/zookeeper
            name: data
          - mountPath: /etc/kafka-configmap
            name: configmap
        dnsPolicy: ClusterFirst
        initContainers:
        - command:
          - /bin/bash
          - /etc/kafka-configmap/init.sh
          image: quay.io/openshifttest/kafka-initutils@sha256:e73ff7a44b43b85b53849c0459ba32e704540852b885a5c78af9753f86a49d68
          imagePullPolicy: IfNotPresent
          name: init-config
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /etc/kafka-configmap
            name: configmap
          - mountPath: /etc/kafka
            name: config
          - mountPath: /var/lib/zookeeper
            name: data
        restartPolicy: Always
        nodeSelector:
          kubernetes.io/arch: amd64
          kubernetes.io/os: linux
        schedulerName: default-scheduler
        terminationGracePeriodSeconds: 10
        volumes:
        - configMap:
            defaultMode: 420
            name: ${CM_NAME}
          name: configmap
        - emptyDir: {}
          name: config
        - emptyDir: {}
          name: zookeeperlogs
        - emptyDir: {}
          name: data
    updateStrategy:
      type: RollingUpdate
parameters:
- name: NAME
  value: "zookeeper"
- name: NAMESPACE
  value: "openshift-logging"
- name: SERVICENAME
  value: "zookeeper"
- name: CM_NAME
  value: "zookeeper"
`)

func loggingExternalLogStoresKafkaZookeeperZookeeperStatefulsetYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaZookeeperZookeeperStatefulsetYaml, nil
}

func loggingExternalLogStoresKafkaZookeeperZookeeperStatefulsetYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaZookeeperZookeeperStatefulsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/zookeeper/zookeeper-statefulset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresKafkaZookeeperZookeeperSvcYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: kafkaserver-template
objects:
- apiVersion: v1
  kind: Service
  metadata:
    labels:
      logging-infra: support
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    ports:
    - name: client
      port: 2181
      protocol: TCP
      targetPort: 2181
    - name: peer
      port: 2888
      protocol: TCP
      targetPort: 2888
    - name: leader-election
      port: 3888
      protocol: TCP
      targetPort: 3888
    selector:
      component: zookeeper
      provider: openshift
    sessionAffinity: None
    type: ClusterIP
parameters:
- name: NAME
  value: "zookeeper"
- name: NAMESPACE
  value: "openshift-logging"
`)

func loggingExternalLogStoresKafkaZookeeperZookeeperSvcYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresKafkaZookeeperZookeeperSvcYaml, nil
}

func loggingExternalLogStoresKafkaZookeeperZookeeperSvcYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresKafkaZookeeperZookeeperSvcYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/kafka/zookeeper/zookeeper-svc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresLokiLokiConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: loki-config-template
objects:
- apiVersion: v1
  kind: ConfigMap
  metadata:
    name: ${LOKICMNAME}
    namespace: ${LOKINAMESPACE}
  data:
    local-config.yaml: |
      auth_enabled: false

      server:
        http_listen_port: 3100
        grpc_listen_port: 9096
        grpc_server_max_recv_msg_size: 8388608

      ingester:
        wal:
          enabled: true
          dir: /tmp/wal
        lifecycler:
          address: 127.0.0.1
          ring:
            kvstore:
              store: inmemory
            replication_factor: 1
          final_sleep: 0s
        chunk_idle_period: 1h       # Any chunk not receiving new logs in this time will be flushed
        chunk_target_size: 8388608
        max_chunk_age: 1h           # All chunks will be flushed when they hit this age, default is 1h
        chunk_retain_period: 30s    # Must be greater than index read cache TTL if using an index cache (Default index read cache TTL is 5m)
        max_transfer_retries: 0     # Chunk transfers disabled

      schema_config:
        configs:
          - from: 2020-10-24
            store: boltdb-shipper
            object_store: filesystem
            schema: v11
            index:
              prefix: index_
              period: 24h

      storage_config:
        boltdb_shipper:
          active_index_directory: /tmp/loki/boltdb-shipper-active
          cache_location: /tmp/loki/boltdb-shipper-cache
          cache_ttl: 24h         # Can be increased for faster performance over longer query periods, uses more disk space
          shared_store: filesystem
        filesystem:
          directory: /tmp/loki/chunks

      compactor:
        working_directory: /tmp/loki/boltdb-shipper-compactor
        shared_store: filesystem

      limits_config:
        reject_old_samples: true
        reject_old_samples_max_age: 12h
        ingestion_rate_mb: 8
        ingestion_burst_size_mb: 16

      chunk_store_config:
        max_look_back_period: 0s

      table_manager:
        retention_deletes_enabled: false
        retention_period: 0s

      ruler:
        storage:
          type: local
          local:
            directory: /tmp/loki/rules
        rule_path: /tmp/loki/rules-temp
        alertmanager_url: http://localhost:9093
        ring:
          kvstore:
            store: inmemory
        enable_api: true
parameters:
- name: LOKINAMESPACE
  value: "loki-aosqe"
- name: LOKICMNAME
  value: "loki-config"
`)

func loggingExternalLogStoresLokiLokiConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresLokiLokiConfigmapYaml, nil
}

func loggingExternalLogStoresLokiLokiConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresLokiLokiConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/loki/loki-configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresLokiLokiDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: loki-log-store-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    name:  ${LOKISERVERNAME}
    namespace: ${LOKINAMESPACE}
    labels:
      provider: aosqe
      component: "loki"
      appname: loki-server
  spec:
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        provider: aosqe
        component: "loki"
    strategy:
      type: Recreate
    template:
      metadata:
        labels:
          provider: aosqe
          component: "loki"
          appname: loki-server
      spec:
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - name: "loki"
          image: quay.io/openshifttest/grafana-loki@sha256:bbf6dbf3264af939a541b6f3c014cba21a2bdc8f22cb7962eee7e9048b41ea5e
          imagePullPolicy: IfNotPresent
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          ports:
          - containerPort: 3100
            name: tcp
            protocol: TCP
          volumeMounts:
          - mountPath: /etc/loki
            name: lokiconfig
            readOnly: true
        volumes:
        - configMap:
            defaultMode: 420
            name: ${LOKICMNAME}
          name: lokiconfig
parameters:
- name: LOKISERVERNAME
  value: "loki-server"
- name: LOKINAMESPACE
  value: "loki-aosqe"
- name: LOKICMNAME
  value: "loki-config"
`)

func loggingExternalLogStoresLokiLokiDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresLokiLokiDeploymentYaml, nil
}

func loggingExternalLogStoresLokiLokiDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresLokiLokiDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/loki/loki-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresOtelOtelCollectorYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: openTelemetryCollector-template
objects:
- apiVersion: opentelemetry.io/v1beta1
  kind: OpenTelemetryCollector
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    config:
      exporters:
        debug:
          verbosity: detailed
      receivers:
        otlp:
          protocols:
            http:
              endpoint: 0.0.0.0:4318
      service:
        pipelines:
          logs:
            exporters:
            - debug
            processors: []
            receivers:
            - otlp
    managementState: managed
    mode: deployment
    replicas: 1
    upgradeStrategy: automatic
parameters:
- name: NAME
  value: "otel"
- name: NAMESPACE
  value: "openshift-opentelemetry-operator"
`)

func loggingExternalLogStoresOtelOtelCollectorYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresOtelOtelCollectorYaml, nil
}

func loggingExternalLogStoresOtelOtelCollectorYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresOtelOtelCollectorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/otel/otel-collector.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresRsyslogInsecureConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: rsyslogserver-template
objects:
- apiVersion: v1
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
    labels:
      provider: aosqe
      component: ${NAME}
  data:
    rsyslog.conf: |+
      global(processInternalMessages="on")
      module(load="imptcp")
      module(load="imudp" TimeRequery="500")
      input(type="imptcp" port="10514")
      input(type="imudp" port="10514")
      :msg, contains, "\"log_type\":\"application\"" /var/log/clf/app-container.log
      :msg, contains, "\"log_type\":\"infrastructure\""{
        if $msg contains "\"log_source\":\"container\"" then /var/log/clf/infra-container.log
        if $msg contains "\"log_source\":\"node\"" then /var/log/clf/infra.log
      }
      :msg, contains, "\"log_type\":\"audit\"" /var/log/clf/audit.log
      :msg, contains, "\"log_source\":\"auditd\"" /var/log/clf/audit-linux.log
      :msg, contains, "\"log_source\":\"kubeAPI\"" /var/log/clf/audit-kubeAPI.log
      :msg, contains, "\"log_source\":\"openshiftAPI\"" /var/log/clf/audit-openshiftAPI.log
      :msg, contains, "\"log_source\":\"ovn\"" /var/log/clf/audit-ovn.log
      *.* /var/log/clf/other.log
parameters:
- name: NAME
  value: "rsyslogserver"
- name: NAMESPACE
  value: "openshift-logging"
`)

func loggingExternalLogStoresRsyslogInsecureConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresRsyslogInsecureConfigmapYaml, nil
}

func loggingExternalLogStoresRsyslogInsecureConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresRsyslogInsecureConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/rsyslog/insecure/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresRsyslogInsecureDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: rsyslogserver-template
objects:
- kind: Deployment
  apiVersion: apps/v1
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
    labels:
      provider: aosqe
      component: ${NAME}
  spec:
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        provider: aosqe
        component: ${NAME}
    strategy:
      type: Recreate
    template:
      metadata:
        labels:
          provider: aosqe
          component: ${NAME}
      spec:
        serviceAccount: ${NAME}
        serviceAccountName: ${NAME}
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - name: "rsyslog"
          command: ["/usr/sbin/rsyslogd", "-f", "/etc/rsyslog/conf/rsyslog.conf", "-n"]
          image: quay.io/openshifttest/rsyslogd-container@sha256:e806eb41f05d7cc6eec96bf09c7bcb692f97562d4a983cb019289bd048d9aee2
          imagePullPolicy: IfNotPresent
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          ports:
          - containerPort: 10514
            name: rsyslog-pod-tcp
            protocol: TCP
          - containerPort: 10514
            name: rsyslog-pod-udp
            protocol: UDP
          volumeMounts:
          - mountPath: /etc/rsyslog/conf
            name: main
            readOnly: true
        volumes:
        - configMap:
            defaultMode: 420
            name: ${NAME}
          name: main
parameters:
- name: NAME
  value: "rsyslogserver"
- name: NAMESPACE
  value: "openshift-logging"
`)

func loggingExternalLogStoresRsyslogInsecureDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresRsyslogInsecureDeploymentYaml, nil
}

func loggingExternalLogStoresRsyslogInsecureDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresRsyslogInsecureDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/rsyslog/insecure/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresRsyslogInsecureSvcYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: rsyslogserver-template
objects:
- apiVersion: v1
  kind: Service
  metadata:
    labels:
      provider: aosqe
      component: ${NAME}
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    ports:
    - name: rsyslogserver-tcp
      port: 514
      targetPort: 10514
      protocol: TCP
    - name: rsyslogserver-udp
      port: 514
      targetPort: 10514
      protocol: UDP
    selector:
      component: ${NAME}
      provider: aosqe
parameters:
- name: NAME
  value: "rsyslogserver"
- name: NAMESPACE
  value: "openshift-logging"
`)

func loggingExternalLogStoresRsyslogInsecureSvcYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresRsyslogInsecureSvcYaml, nil
}

func loggingExternalLogStoresRsyslogInsecureSvcYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresRsyslogInsecureSvcYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/rsyslog/insecure/svc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresRsyslogSecureConfigmapYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: rsyslogserver-template
objects:
- apiVersion: v1
  kind: ConfigMap
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
    labels:
      provider: aosqe
  data:
    rsyslog.conf: |+
      global(
        DefaultNetstreamDriverCAFile="/opt/app-root/tls/ca_bundle.crt"
        DefaultNetstreamDriverCertFile="/opt/app-root/tls/server.crt"
        DefaultNetstreamDriverKeyFile="/opt/app-root/tls/server.key"
      )
      module( load="imtcp"
        StreamDriver.Name = "gtls"
        StreamDriver.Mode = "1"
        #https://www.rsyslog.com/doc/master/concepts/ns_ossl.html
        StreamDriver.AuthMode = "anon"
      )
      module(load="imudp" TimeRequery="500")
      input(type="imtcp" port="6514")
      input(type="imudp" port="10514")
      :msg, contains, "\"log_type\":\"application\"" /var/log/clf/app-container.log
      :msg, contains, "\"log_type\":\"infrastructure\""{
        if $msg contains "\"log_source\":\"container\"" then /var/log/clf/infra-container.log
        if $msg contains "\"log_source\":\"node\"" then /var/log/clf/infra.log
      }
      :msg, contains, "\"log_type\":\"audit\"" /var/log/clf/audit.log
      :msg, contains, "\"log_source\":\"auditd\"" /var/log/clf/audit-linux.log
      :msg, contains, "\"log_source\":\"kubeAPI\"" /var/log/clf/audit-kubeAPI.log
      :msg, contains, "\"log_source\":\"openshiftAPI\"" /var/log/clf/audit-openshiftAPI.log
      :msg, contains, "\"log_source\":\"ovn\"" /var/log/clf/audit-ovn.log
      *.* /var/log/clf/other.log
parameters:
- name: NAME
  value: "rsyslogserver"
- name: NAMESPACE
  value: "openshift-logging"
`)

func loggingExternalLogStoresRsyslogSecureConfigmapYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresRsyslogSecureConfigmapYaml, nil
}

func loggingExternalLogStoresRsyslogSecureConfigmapYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresRsyslogSecureConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/rsyslog/secure/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresRsyslogSecureDeploymentYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: rsyslogserver-template
objects:
- kind: Deployment
  apiVersion: apps/v1
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
    labels:
      provider: aosqe
      component: ${NAME}
  spec:
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        provider: aosqe
        component: ${NAME}
    strategy:
      type: Recreate
    template:
      metadata:
        labels:
          provider: aosqe
          component: ${NAME}
      spec:
        serviceAccount: ${NAME}
        serviceAccountName: ${NAME}
        securityContext:
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        containers:
        - name: "rsyslog"
          command: ["/usr/sbin/rsyslogd", "-f", "/etc/rsyslog/conf/rsyslog.conf", "-n"]
          image: quay.io/openshifttest/rsyslogd-container@sha256:e806eb41f05d7cc6eec96bf09c7bcb692f97562d4a983cb019289bd048d9aee2
          imagePullPolicy: IfNotPresent
          securityContext:
            allowPrivilegeEscalation: false
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
            seccompProfile:
              type: RuntimeDefault
          ports:
          - containerPort: 10514
            name: rsyslog-pod-tcp
            protocol: TCP
          - containerPort: 10514
            name: rsyslog-pod-udp
            protocol: UDP
          - containerPort: 6514
            name: rsyslog-pod-tls
            protocol: TCP
          volumeMounts:
          - mountPath: /etc/rsyslog/conf
            name: main
            readOnly: true
          - mountPath: /opt/app-root/tls
            name: keys
            readOnly: true
        volumes:
        - configMap:
            defaultMode: 420
            name: ${NAME}
          name: main
        - secret:
            defaultMode: 420
            secretName: ${NAME}
          name: keys
parameters:
- name: NAME
  value: "rsyslogserver"
- name: NAMESPACE
  value: "openshift-logging"
`)

func loggingExternalLogStoresRsyslogSecureDeploymentYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresRsyslogSecureDeploymentYaml, nil
}

func loggingExternalLogStoresRsyslogSecureDeploymentYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresRsyslogSecureDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/rsyslog/secure/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresRsyslogSecureSvcYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: rsyslogserver-template
objects:
- apiVersion: v1
  kind: Service
  metadata:
    labels:
      provider: aosqe
      component: ${NAME}
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    ports:
    - name: rsyslogserver-tls
      port: 6514
      targetPort: 6514
      protocol: TCP
    - name: rsyslogserver-tcp
      port: 514
      targetPort: 10514
      protocol: TCP
    - name: rsyslogserver-udp
      port: 514
      targetPort: 10514
      protocol: UDP
    selector:
      component: ${NAME}
      provider: aosqe
parameters:
- name: NAME
  value: "rsyslogserver"
- name: NAMESPACE
  value: "openshift-logging"
`)

func loggingExternalLogStoresRsyslogSecureSvcYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresRsyslogSecureSvcYaml, nil
}

func loggingExternalLogStoresRsyslogSecureSvcYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresRsyslogSecureSvcYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/rsyslog/secure/svc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresSplunkRouteEdge_splunk_templateYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: route-edge-splunk-template
objects:
- apiVersion: route.openshift.io/v1
  kind: Route
  metadata:
    name: ${NAME}
  spec:
    host: ${ROUTE_HOST}
    port:
      targetPort: ${PORT_NAME} 
    tls:
      insecureEdgeTerminationPolicy: Allow
      termination: edge
    to:
      kind: Service
      name: ${SERVICE_NAME}
    wildcardPolicy: None
parameters:
- name: NAME
  value: "splunk-default-hec"
- name: PORT_NAME
  value: "http-hec"
- name: SERVICE_NAME
  value: "splunk-default-service"
- name: ROUTE_HOST
  value: ""
`)

func loggingExternalLogStoresSplunkRouteEdge_splunk_templateYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresSplunkRouteEdge_splunk_templateYaml, nil
}

func loggingExternalLogStoresSplunkRouteEdge_splunk_templateYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresSplunkRouteEdge_splunk_templateYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/splunk/route-edge_splunk_template.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresSplunkRoutePassthrough_splunk_templateYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: route-splunk-passthrough-template
objects:
- apiVersion: route.openshift.io/v1
  kind: Route
  metadata:
    name: ${NAME}
  spec:
    host: ${ROUTE_HOST}
    port:
      targetPort: ${PORT_NAME}
    tls:
      termination: passthrough
    to:
      kind: Service
      name: ${SERVICE_NAME}
parameters:
- name: NAME
  value: "splunk-default-hec"
- name: PORT_NAME
  value: "http-hec"
- name: SERVICE_NAME
  value: "splunk-default-service"
- name: ROUTE_HOST
  value: ""
`)

func loggingExternalLogStoresSplunkRoutePassthrough_splunk_templateYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresSplunkRoutePassthrough_splunk_templateYaml, nil
}

func loggingExternalLogStoresSplunkRoutePassthrough_splunk_templateYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresSplunkRoutePassthrough_splunk_templateYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/splunk/route-passthrough_splunk_template.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresSplunkSecret_splunk_templateYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: splunk-secret-template
objects:
- apiVersion: v1
  kind: Secret
  metadata:
    name: ${NAME}
  type: Opaque
  stringData:
    default.yml: |
      splunk:
        listenOnIPv6: "yes"
        hec_token: "${HEC_TOKEN}"
        password: "${PASSWORD}"
        pass4SymmKey: "${PASSWORD}"
        idxc:
            secret: "${PASSWORD}"
        shc:
            secret: "${PASSWORD}"
        hec:
            requireClientCert: False
            ssl: False
    hec_token: ${HEC_TOKEN}
    idxc_secret: ${PASSWORD}
    pass4SymmKey: ${PASSWORD}
    password: ${PASSWORD}
    shc_secret: ${PASSWORD}
parameters:
- name: NAME
  value: "splunk-default"
- name: HEC_TOKEN
  value: "555555555-BBBB-BBBB-BBBB-555555555555"
- name: PASSWORD
  value: ""
`)

func loggingExternalLogStoresSplunkSecret_splunk_templateYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresSplunkSecret_splunk_templateYaml, nil
}

func loggingExternalLogStoresSplunkSecret_splunk_templateYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresSplunkSecret_splunk_templateYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/splunk/secret_splunk_template.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresSplunkSecret_tls_passphrase_splunk_templateYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: splunk-secret-template
objects:
- apiVersion: v1
  kind: Secret
  metadata:
    name: ${NAME}
  type: Opaque
  stringData:
    default.yml: |
      splunk:
        listenOnIPv6: "yes"
        hec:
            enable: true
            token: "${HEC_TOKEN}"
            requireClientCert: ${HEC_CLIENTAUTH}
            cert: "/mnt/splunk-secrets/hec.pem"
            ssl: true
        http_enableSSL: 1
        http_enableSSL_cert: "/mnt/splunk-secrets/cert.pem"
        http_enableSSL_privKey: "/mnt/splunk-secrets/key.pem"
        http_enableSSL_privKey_password: ${PASSPHASE}
        password: "${PASSWORD}"
        pass4SymmKey: "${PASSWORD}"
        idxc:
            secret: "${PASSWORD}"
        shc:
            secret: "${PASSWORD}"
    hec_token: ${HEC_TOKEN}
    idxc_secret: ${PASSWORD}
    pass4SymmKey: ${PASSWORD}
    password: ${PASSWORD}
    shc_secret: ${PASSWORD}
parameters:
- name: NAME
  value: "splunk-default"
- name: PASSWORD
  value: "password"
- name: HEC_TOKEN
  value: "555555555-BBBB-BBBB-BBBB-555555555555"
- name: PASSPHASE
  value: "password"
- name: HEC_CLIENTAUTH
  value: "False"
`)

func loggingExternalLogStoresSplunkSecret_tls_passphrase_splunk_templateYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresSplunkSecret_tls_passphrase_splunk_templateYaml, nil
}

func loggingExternalLogStoresSplunkSecret_tls_passphrase_splunk_templateYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresSplunkSecret_tls_passphrase_splunk_templateYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/splunk/secret_tls_passphrase_splunk_template.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresSplunkSecret_tls_splunk_templateYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: splunk-secret-template
objects:
- apiVersion: v1
  kind: Secret
  metadata:
    name: ${NAME}
  type: Opaque
  stringData:
    default.yml: |
      splunk:
        listenOnIPv6: "yes"
        hec:
            enable: true
            ssl: true
            token: "${HEC_TOKEN}"
            requireClientCert: ${HEC_CLIENTAUTH}
            cert: "/mnt/splunk-secrets/hec.pem"
        http_enableSSL: 1
        http_enableSSL_cert: "/mnt/splunk-secrets/cert.pem"
        http_enableSSL_privKey: "/mnt/splunk-secrets/key.pem"
        password: "${PASSWORD}"
        pass4SymmKey: "${PASSWORD}"
        idxc:
            secret: "${PASSWORD}"
        shc:
            secret: "${PASSWORD}"
    hec_token: "${HEC_TOKEN}"
    idxc_secret: ${PASSWORD}
    pass4SymmKey: ${PASSWORD}
    password: ${PASSWORD}
    shc_secret: ${PASSWORD}
parameters:
- name: NAME
  value: "splunk-default"
- name: PASSWORD
  value: "password"
- name: HEC_TOKEN
  value: "555555555-BBBB-BBBB-BBBB-555555555555"
- name: HEC_CLIENTAUTH
  value: "False"
`)

func loggingExternalLogStoresSplunkSecret_tls_splunk_templateYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresSplunkSecret_tls_splunk_templateYaml, nil
}

func loggingExternalLogStoresSplunkSecret_tls_splunk_templateYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresSplunkSecret_tls_splunk_templateYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/splunk/secret_tls_splunk_template.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresSplunkStatefulset_splunk82_templateYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: splunk-s1-standalone-template
objects:
- apiVersion: apps/v1
  kind: StatefulSet
  metadata:
    name: ${NAME}
  spec:
    podManagementPolicy: Parallel
    replicas: 1
    selector:
      matchLabels:
        app.kubernetes.io/component: splunk
        app.kubernetes.io/instance: ${NAME}
        app.kubernetes.io/name: splunk
    serviceName: ${NAME}-headless
    template:
      metadata:
        annotations:
          traffic.sidecar.istio.io/excludeOutboundPorts: 8089,8191,9997
          traffic.sidecar.istio.io/includeInboundPorts: 8000,8088
        labels:
          app.kubernetes.io/component: splunk
          app.kubernetes.io/instance: ${NAME}
          app.kubernetes.io/name: splunk
      spec:
        containers:
        - env:
          - name: DEBUG
            value: "false"
          - name: ANSIBLE_EXTRA_FLAGS
            value: "-v"
          - name: SPLUNK_DECLARATIVE_ADMIN_PASSWORD
            value: "true"
          - name: SPLUNK_DEFAULTS_URL
            value: /mnt/splunk-secrets/default.yml
          - name: SPLUNK_HOME
            value: /opt/splunk
          - name: SPLUNK_HOME_OWNERSHIP_ENFORCEMENT
            value: "false"
          - name: SPLUNK_ROLE
            value: splunk_standalone
          - name: SPLUNK_START_ARGS
            value: --accept-license
          image: quay.io/openshifttest/splunk@sha256:fbfae0b70a4884a3d23a05d3f45fa35646ea56ccd98ab73fb147b31715a41c42
          imagePullPolicy: IfNotPresent
          livenessProbe:
            exec:
              command:
              - /sbin/checkstate.sh
            failureThreshold: 3
            initialDelaySeconds: 300
            periodSeconds: 30
            successThreshold: 1
            timeoutSeconds: 30
          name: splunk
          ports:
          - containerPort: 8000
            name: http-splunkweb
            protocol: TCP
          - containerPort: 8088
            name: http-hec
            protocol: TCP
          - containerPort: 8089
            name: https-splunkd
            protocol: TCP
          - containerPort: 9997
            name: tcp-s2s
            protocol: TCP
          readinessProbe:
            exec:
              command:
              - /bin/grep
              - started
              - /opt/container_artifact/splunk-container.state
            failureThreshold: 3
            initialDelaySeconds: 10
            periodSeconds: 5
            successThreshold: 1
            timeoutSeconds: 5
          resources:
            limits:
              cpu: "4"
              memory: 8Gi
            requests:
              cpu: 100m
              memory: 512Mi
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              drop:
              - ALL
            runAsNonRoot: true
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /mnt/splunk-secrets
            name: mnt-splunk-secrets
          - mountPath: /opt/splunk/etc
            name: pvc-etc
          - mountPath: /opt/splunk/var
            name: pvc-var
        dnsPolicy: ClusterFirst
        nodeSelector:
          kubernetes.io/arch: amd64
          kubernetes.io/os: linux
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          fsGroup: 41812
          runAsNonRoot: true
          runAsUser: 41812
        terminationGracePeriodSeconds: 30
        volumes:
        - name: mnt-splunk-secrets
          secret:
            defaultMode: 420
            secretName: ${NAME}
    updateStrategy:
      type: OnDelete
    volumeClaimTemplates:
    - apiVersion: v1
      kind: PersistentVolumeClaim
      metadata:
        labels:
          app.kubernetes.io/component: splunk
          app.kubernetes.io/instance: ${NAME}
          app.kubernetes.io/name: splunk
        name: pvc-etc
      spec:
        accessModes:
        - ReadWriteOnce
        resources:
          requests:
            storage: 1Gi
        volumeMode: Filesystem
    - apiVersion: v1
      kind: PersistentVolumeClaim
      metadata:
        labels:
          app.kubernetes.io/component: splunk
          app.kubernetes.io/instance: ${NAME}
          app.kubernetes.io/name: splunk
        name: pvc-var
      spec:
        accessModes:
        - ReadWriteOnce
        resources:
          requests:
            storage: 10Gi
        volumeMode: Filesystem
- apiVersion: v1
  kind: Service
  metadata:
    labels:
      app.kubernetes.io/component: splunk
      app.kubernetes.io/instance: ${NAME}
      app.kubernetes.io/name: splunk
    name: ${NAME}-headless
  spec:
    type: ClusterIP
    clusterIP: None
    selector:
      app.kubernetes.io/component: splunk
      app.kubernetes.io/instance: ${NAME}
      app.kubernetes.io/name: splunk
    ports:
    - name: http-splunkweb
      port: 8000
      protocol: TCP
      targetPort: 8000
    - name: http-hec
      port: 8088
      protocol: TCP
      targetPort: 8088
    - name: https-splunkd
      port: 8089
      protocol: TCP
      targetPort: 8089
    - name: tcp-s2s
      port: 9997
      protocol: TCP
      targetPort: 9997
- apiVersion: v1
  kind: Service
  metadata:
    labels:
      app.kubernetes.io/component: splunk
      app.kubernetes.io/instance: ${NAME}
      app.kubernetes.io/name: splunk
    name: ${NAME}-0
  spec:
    type: ClusterIP
    selector:
      app.kubernetes.io/component: splunk
      app.kubernetes.io/instance: ${NAME}
      app.kubernetes.io/name: splunk
    internalTrafficPolicy: Cluster
    ports:
    - name: http-splunkweb
      port: 8000
      protocol: TCP
      targetPort: 8000
    - name: http-hec
      port: 8088
      protocol: TCP
      targetPort: 8088
    - name: https-splunkd
      port: 8089
      protocol: TCP
      targetPort: 8089
    - name: tcp-s2s
      port: 9997
      protocol: TCP
      targetPort: 9997
parameters:
- name: NAME
  value: "splunk-s1-standalone"
`)

func loggingExternalLogStoresSplunkStatefulset_splunk82_templateYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresSplunkStatefulset_splunk82_templateYaml, nil
}

func loggingExternalLogStoresSplunkStatefulset_splunk82_templateYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresSplunkStatefulset_splunk82_templateYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/splunk/statefulset_splunk-8.2_template.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingExternalLogStoresSplunkStatefulset_splunk90_templateYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: splunk-s1-standalone-template
objects:
- apiVersion: apps/v1
  kind: StatefulSet
  metadata:
    name: ${NAME}
  spec:
    podManagementPolicy: Parallel
    replicas: 1
    selector:
      matchLabels:
        app.kubernetes.io/component: splunk
        app.kubernetes.io/instance: ${NAME}
        app.kubernetes.io/name: splunk
    serviceName: ${NAME}-headless
    template:
      metadata:
        annotations:
          traffic.sidecar.istio.io/excludeOutboundPorts: 8089,8191,9997
          traffic.sidecar.istio.io/includeInboundPorts: 8000,8088
        labels:
          app.kubernetes.io/component: splunk
          app.kubernetes.io/instance: ${NAME}
          app.kubernetes.io/name: splunk
      spec:
        containers:
        - env:
          - name: DEBUG
            value: "false"
          - name: ANSIBLE_EXTRA_FLAGS
            value: "-v"
          - name: SPLUNK_DECLARATIVE_ADMIN_PASSWORD
            value: "true"
          - name: SPLUNK_DEFAULTS_URL
            value: /mnt/splunk-secrets/default.yml
          - name: SPLUNK_HOME
            value: /opt/splunk
          - name: SPLUNK_HOME_OWNERSHIP_ENFORCEMENT
            value: "false"
          - name: SPLUNK_ROLE
            value: splunk_standalone
          - name: SPLUNK_START_ARGS
            value: --accept-license
          image: quay.io/openshifttest/splunk@sha256:5762a3b61ad5090f24ad33360fc03f3ced469e16c3c75f6d8590b5ef39d95751
          imagePullPolicy: IfNotPresent
          livenessProbe:
            exec:
              command:
              - /sbin/checkstate.sh
            failureThreshold: 3
            initialDelaySeconds: 300
            periodSeconds: 30
            successThreshold: 1
            timeoutSeconds: 30
          name: splunk
          ports:
          - containerPort: 8000
            name: http-splunkweb
            protocol: TCP
          - containerPort: 8088
            name: http-hec
            protocol: TCP
          - containerPort: 8089
            name: https-splunkd
            protocol: TCP
          - containerPort: 9997
            name: tcp-s2s
            protocol: TCP
          readinessProbe:
            exec:
              command:
              - /bin/grep
              - started
              - /opt/container_artifact/splunk-container.state
            failureThreshold: 3
            initialDelaySeconds: 10
            periodSeconds: 5
            successThreshold: 1
            timeoutSeconds: 5
          resources:
            limits:
              cpu: "4"
              memory: 8Gi
            requests:
              cpu: 100m
              memory: 512Mi
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              drop:
              - ALL
            runAsNonRoot: true
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /mnt/splunk-secrets
            name: mnt-splunk-secrets
          - mountPath: /opt/splunk/etc
            name: pvc-etc
          - mountPath: /opt/splunk/var
            name: pvc-var
        dnsPolicy: ClusterFirst
        nodeSelector:
          kubernetes.io/arch: amd64
          kubernetes.io/os: linux
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          fsGroup: 41812
          runAsNonRoot: true
          runAsUser: 41812
        terminationGracePeriodSeconds: 30
        volumes:
        - name: mnt-splunk-secrets
          secret:
            defaultMode: 420
            secretName: ${NAME}
    updateStrategy:
      type: OnDelete
    volumeClaimTemplates:
    - apiVersion: v1
      kind: PersistentVolumeClaim
      metadata:
        labels:
          app.kubernetes.io/component: splunk
          app.kubernetes.io/instance: ${NAME}
          app.kubernetes.io/name: splunk
        name: pvc-etc
      spec:
        accessModes:
        - ReadWriteOnce
        resources:
          requests:
            storage: 1Gi
        volumeMode: Filesystem
    - apiVersion: v1
      kind: PersistentVolumeClaim
      metadata:
        labels:
          app.kubernetes.io/component: splunk
          app.kubernetes.io/instance: ${NAME}
          app.kubernetes.io/name: splunk
        name: pvc-var
      spec:
        accessModes:
        - ReadWriteOnce
        resources:
          requests:
            storage: 10Gi
        volumeMode: Filesystem
- apiVersion: v1
  kind: Service
  metadata:
    labels:
      app.kubernetes.io/component: splunk
      app.kubernetes.io/instance: ${NAME}
      app.kubernetes.io/name: splunk
    name: ${NAME}-headless
  spec:
    type: ClusterIP
    clusterIP: None
    selector:
      app.kubernetes.io/component: splunk
      app.kubernetes.io/instance: ${NAME}
      app.kubernetes.io/name: splunk
    ports:
    - name: http-splunkweb
      port: 8000
      protocol: TCP
      targetPort: 8000
    - name: http-hec
      port: 8088
      protocol: TCP
      targetPort: 8088
    - name: https-splunkd
      port: 8089
      protocol: TCP
      targetPort: 8089
    - name: tcp-s2s
      port: 9997
      protocol: TCP
      targetPort: 9997
- apiVersion: v1
  kind: Service
  metadata:
    labels:
      app.kubernetes.io/component: splunk
      app.kubernetes.io/instance: ${NAME}
      app.kubernetes.io/name: splunk
    name: ${NAME}-0
  spec:
    type: ClusterIP
    selector:
      app.kubernetes.io/component: splunk
      app.kubernetes.io/instance: ${NAME}
      app.kubernetes.io/name: splunk
    internalTrafficPolicy: Cluster
    ports:
    - name: http-splunkweb
      port: 8000
      protocol: TCP
      targetPort: 8000
    - name: http-hec
      port: 8088
      protocol: TCP
      targetPort: 8088
    - name: https-splunkd
      port: 8089
      protocol: TCP
      targetPort: 8089
    - name: tcp-s2s
      port: 9997
      protocol: TCP
      targetPort: 9997
parameters:
- name: NAME
  value: "splunk-s1-standalone"
`)

func loggingExternalLogStoresSplunkStatefulset_splunk90_templateYamlBytes() ([]byte, error) {
	return _loggingExternalLogStoresSplunkStatefulset_splunk90_templateYaml, nil
}

func loggingExternalLogStoresSplunkStatefulset_splunk90_templateYaml() (*asset, error) {
	bytes, err := loggingExternalLogStoresSplunkStatefulset_splunk90_templateYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/external-log-stores/splunk/statefulset_splunk-9.0_template.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingGeneratelog42981Yaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: ovn-audit-log-gen-template
objects:
  - kind: Namespace
    apiVersion: v1
    metadata:
      annotations:
        k8s.ovn.org/acl-logging: '{ "deny": "alert", "allow": "alert" }'
      name: ${NAMESPACE}
    spec:
      finalizers:
      - kubernetes

  - kind: Deployment
    apiVersion: apps/v1
    metadata:
      labels:
        app: ovn-app
      name: ovn-app
      namespace: ${NAMESPACE}
    spec:
      replicas: 2
      selector:
        matchLabels:
          app: ovn-app
      strategy: {}
      template:
        metadata:
          labels:
            app: ovn-app
        spec:
          containers:
          - image: quay.io/openshifttest/hello-sdn@sha256:c89445416459e7adea9a5a416b3365ed3d74f2491beb904d61dc8d1eb89a72a4
            name: hello-sdn
            resources:
              limits:
                memory: 340Mi

  - kind: Service
    apiVersion: v1
    metadata:
      labels:
        name: test-service
      name: test-service
      namespace: ${NAMESPACE}
    spec:
      ports:
      - name: http
        port: 27017
        protocol: TCP
        targetPort: 8080
      selector:
        app: ovn-app

  - kind: NetworkPolicy
    apiVersion: networking.k8s.io/v1
    metadata:
      name: default-deny
      namespace: ${NAMESPACE}
    spec:
      podSelector:

  - kind: NetworkPolicy
    apiVersion: networking.k8s.io/v1
    metadata:
      name: allow-same-namespace
      namespace: ${NAMESPACE}
    spec:
      podSelector:
      ingress:
      - from:
        - podSelector: {}

  - apiVersion: networking.k8s.io/v1
    kind: NetworkPolicy
    metadata:
      name: bad-np
      namespace: ${NAMESPACE}
    spec:
      egress:
      - {}
      podSelector:
        matchLabels:
          never-gonna: match
      policyTypes:
      - Egress

parameters:
  - name: NAMESPACE
    value: "openshift-logging"
`)

func loggingGeneratelog42981YamlBytes() ([]byte, error) {
	return _loggingGeneratelog42981Yaml, nil
}

func loggingGeneratelog42981Yaml() (*asset, error) {
	bytes, err := loggingGeneratelog42981YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/generatelog/42981.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingGeneratelogContainer_json_log_templateJson = []byte(`{
  "apiVersion": "template.openshift.io/v1",
  "kind": "Template",
  "metadata": {
    "name": "centos-logtest-template"
  },
  "objects": [
    {
      "apiVersion": "v1",
      "data": {
        "ocp_logtest.cfg": "--raw --file /var/lib/svt/json.example  --text-type input --rate ${RATE}",
        "json.example": "{\"message\": \"MERGE_JSON_LOG=true\", \"level\": \"debug\",\"Layer1\": \"layer1 0\", \"layer2\": {\"name\":\"Layer2 1\", \"tips\":\"Decide by PRESERVE_JSON_LOG\"}, \"StringNumber\":\"10\", \"Number\": 10,\"foo.bar\":\"Dot Item\",\"{foobar}\":\"Brace Item\",\"[foobar]\":\"Bracket Item\", \"foo:bar\":\"Colon Item\",\"foo bar\":\"Space Item\" }"
      },
      "kind": "ConfigMap",
      "metadata": {
        "name": "${CONFIGMAP}"
      }
    },
    {
      "apiVersion": "v1",
      "kind": "ReplicationController",
      "metadata": {
        "name": "${REPLICATIONCONTROLLER}",
        "labels": "${{LABELS}}"
      },
      "spec": {
        "replicas": "${{REPLICAS}}",
        "template": {
          "metadata": {
            "generateName": "centos-logtest-",
            "annotations": {
              "containerType.logging.openshift.io/${CONTAINER}": "${CONTAINER}"
            },
            "labels": "${{LABELS}}"
          },
          "spec": {
            "containers": [
              {
                "env": [],
                "image": "quay.io/openshifttest/ocp-logtest@sha256:6e2973d7d454ce412ad90e99ce584bf221866953da42858c4629873e53778606",
                "imagePullPolicy": "IfNotPresent",
                "name": "${CONTAINER}",
                "resources": {},
                "volumeMounts": [
                  {
                    "name": "config",
                    "mountPath": "/var/lib/svt"
                  }
                ],
                "securityContext": {
                  "allowPrivilegeEscalation": false,
                  "capabilities": {
                    "drop": [
                      "ALL"
                    ]
                  }
                },
                "terminationMessagePath": "/dev/termination-log"
              }
            ],
            "securityContext": {
              "runAsNonRoot": true,
              "seccompProfile": {
                "type": "RuntimeDefault"
              }
            },
            "volumes": [
              {
                "name": "config",
                "configMap": {
                  "name": "${CONFIGMAP}"
                }
              }
            ]
          }
        }
      }
    }
  ],
  "parameters": [
    {
      "name": "REPLICAS",
      "displayName": "Replicas",
      "value": "1"
    },
    {
      "name": "LABELS",
      "displayName": "labels",
      "value": "{\"run\": \"centos-logtest\", \"test\": \"centos-logtest\"}"
    },
    {
      "name": "REPLICATIONCONTROLLER",
      "displayName": "ReplicationController",
      "value": "logging-centos-logtest"
    },
    {
      "name": "CONFIGMAP",
      "displayName": "ConfigMap",
      "value": "logtest-config"
    },
    {
      "name": "CONTAINER",
      "value": "logging-centos-logtest"
    },
    {
      "name": "RATE",
      "value": "60.0"
    }
  ]
}
`)

func loggingGeneratelogContainer_json_log_templateJsonBytes() ([]byte, error) {
	return _loggingGeneratelogContainer_json_log_templateJson, nil
}

func loggingGeneratelogContainer_json_log_templateJson() (*asset, error) {
	bytes, err := loggingGeneratelogContainer_json_log_templateJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/generatelog/container_json_log_template.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingGeneratelogContainer_json_log_template_unannotedJson = []byte(`{
  "apiVersion": "template.openshift.io/v1",
  "kind": "Template",
  "metadata": {
    "name": "centos-logtest-template"
  },
  "objects": [
    {
      "apiVersion": "v1",
      "data": {
        "ocp_logtest.cfg": "--raw --file /var/lib/svt/json.example  --text-type input --rate 60.0",
        "json.example": "{\"message\": \"MERGE_JSON_LOG=true\", \"level\": \"debug\",\"Layer1\": \"layer1 0\", \"layer2\": {\"name\":\"Layer2 1\", \"tips\":\"Decide by PRESERVE_JSON_LOG\"}, \"StringNumber\":\"10\", \"Number\": 10,\"foo.bar\":\"Dot Item\",\"{foobar}\":\"Brace Item\",\"[foobar]\":\"Bracket Item\", \"foo:bar\":\"Colon Item\",\"foo bar\":\"Space Item\" }"
      },
      "kind": "ConfigMap",
      "metadata": {
        "name": "${CONFIGMAP}"
      }
    },
    {
      "apiVersion": "v1",
      "kind": "ReplicationController",
      "metadata": {
        "name": "${{REPLICATIONCONTROLLER}}",
        "labels": "${{LABELS}}"
      },
      "spec": {
        "replicas": "${{REPLICAS}}",
        "template": {
          "metadata": {
            "generateName": "centos-logtest-",
            "labels": "${{LABELS}}"
          },
          "spec": {
            "containers": [
              {
                "env": [],
                "image": "quay.io/openshifttest/ocp-logtest@sha256:6e2973d7d454ce412ad90e99ce584bf221866953da42858c4629873e53778606",
                "imagePullPolicy": "IfNotPresent",
                "name": "${CONTAINER}",
                "resources": {},
                "volumeMounts": [
                  {
                    "name": "config",
                    "mountPath": "/var/lib/svt"
                  }
                ],
                "securityContext": {
                  "allowPrivilegeEscalation": false,
                  "capabilities": {
                    "drop": [
                      "ALL"
                    ]
                  }
                },
                "terminationMessagePath": "/dev/termination-log"
              }
            ],
            "securityContext": {
              "runAsNonRoot": true,
              "seccompProfile": {
                "type": "RuntimeDefault"
              }
            },
            "volumes": [
              {
                "name": "config",
                "configMap": {
                  "name": "${{CONFIGMAP}}"
                }
              }
            ]
          }
        }
      }
    }
  ],
  "parameters": [
    {
      "name": "REPLICAS",
      "displayName": "Replicas",
      "value": "1"
    },
    {
      "name": "LABELS",
      "displayName": "labels",
      "value": "{\"run\": \"centos-logtest\", \"test\": \"centos-logtest\"}"
    },
    {
      "name": "REPLICATIONCONTROLLER",
      "displayName": "ReplicationController",
      "value": "logging-centos-logtest"
    },
    {
      "name": "CONFIGMAP",
      "displayName": "ConfigMap",
      "value": "logtest-config"
    },
    {
      "name": "CONTAINER",
      "value": "logging-centos-logtest"
    }
  ]
}
`)

func loggingGeneratelogContainer_json_log_template_unannotedJsonBytes() ([]byte, error) {
	return _loggingGeneratelogContainer_json_log_template_unannotedJson, nil
}

func loggingGeneratelogContainer_json_log_template_unannotedJson() (*asset, error) {
	bytes, err := loggingGeneratelogContainer_json_log_template_unannotedJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/generatelog/container_json_log_template_unannoted.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingGeneratelogContainer_non_json_log_templateJson = []byte(`{
  "apiVersion": "template.openshift.io/v1",
  "kind": "Template",
  "metadata": {
    "name": "centos-logtest-template"
  },
  "objects": [
    {
      "apiVersion": "v1",
      "data": {
        "ocp_logtest.cfg": "--raw --file /var/lib/svt/json.example  --text-type input --rate 60.0",
        "json.example": "ㄅㄉˇˋㄓˊ˙ㄚㄞㄢㄦㄆ 中国 883.317µs ā á ǎ à ō ó ▅ ▆ ▇ █ 々"
      },
      "kind": "ConfigMap",
      "metadata": {
        "name": "${{CONFIGMAP}}"
      }
    },
    {
      "apiVersion": "v1",
      "kind": "ReplicationController",
      "metadata": {
        "name": "${{REPLICATIONCONTROLLER}}",
        "labels": {
          "run": "${{LABELS}}",
          "test": "${{LABELS}}"
        }
      },
      "spec": {
        "replicas": "${{REPLICAS}}",
        "template": {
          "metadata": {
            "generateName": "centos-logtest-",
            "labels": {
              "run": "${{LABELS}}",
              "test": "${{LABELS}}"
            }
          },
          "spec": {
            "containers": [
              {
                "env": [],
                "image": "quay.io/openshifttest/ocp-logtest@sha256:6e2973d7d454ce412ad90e99ce584bf221866953da42858c4629873e53778606",
                "imagePullPolicy": "IfNotPresent",
                "name": "logging-centos-logtest",
                "resources": {},
                "volumeMounts": [
                  {
                    "name": "config",
                    "mountPath": "/var/lib/svt"
                  }
                ],
                "securityContext": {
                  "allowPrivilegeEscalation": false,
                  "capabilities": {
                    "drop": [
                      "ALL"
                    ]
                  }
                },
                "terminationMessagePath": "/dev/termination-log"
              }
            ],
            "securityContext": {
              "runAsNonRoot": true,
              "seccompProfile": {
                "type": "RuntimeDefault"
              }
            },
            "volumes": [
              {
                "name": "config",
                "configMap": {
                  "name": "${{CONFIGMAP}}"
                }
              }
            ]
          }
        }
      }
    }
  ],
  "parameters": [
    {
      "name": "REPLICAS",
      "displayName": "Replicas",
      "value": "1"
    },
    {
      "name": "LABELS",
      "displayName": "labels",
      "value": "centos-logtest"
    },
    {
      "name": "REPLICATIONCONTROLLER",
      "displayName": "ReplicationController",
      "value": "logging-centos-logtest"
    },
    {
      "name": "CONFIGMAP",
      "displayName": "ConfigMap",
      "value": "logtest-config"
    }
  ]
}
`)

func loggingGeneratelogContainer_non_json_log_templateJsonBytes() ([]byte, error) {
	return _loggingGeneratelogContainer_non_json_log_templateJson, nil
}

func loggingGeneratelogContainer_non_json_log_templateJson() (*asset, error) {
	bytes, err := loggingGeneratelogContainer_non_json_log_templateJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/generatelog/container_non_json_log_template.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingGeneratelogLoggingPerformanceAppGeneratorJson = []byte(`{
    "apiVersion": "template.openshift.io/v1",
    "kind": "Template",
    "metadata": {
      "name": "centos-logtest-template"
    },
    "objects": [
      {
        "apiVersion": "v1",
        "data": {
          "ocp_logtest.cfg": "--raw --file /var/lib/svt/json.example  --text-type input --rate ${RATE} --num-lines ${NUM_LINES}",
          "json.example": "{\"message\": \"MERGE_JSON_LOG=true\", \"level\": \"debug\",\"Layer1\": \"layer1 0\", \"layer2\": {\"name\":\"Layer2 1\", \"tips\":\"Decide by PRESERVE_JSON_LOG\"}, \"StringNumber\":\"10\", \"Number\": 10,\"foo.bar\":\"Dot Item\",\"{foobar}\":\"Brace Item\",\"[foobar]\":\"Bracket Item\", \"foo:bar\":\"Colon Item\",\"foo bar\":\"Space Item\" }"
        },
        "kind": "ConfigMap",
        "metadata": {
          "name": "${CONFIGMAP}"
        }
      },
      {
        "apiVersion": "v1",
        "kind": "ReplicationController",
        "metadata": {
          "name": "${REPLICATIONCONTROLLER}",
          "labels": "${{LABELS}}"
        },
        "spec": {
          "replicas": "${{REPLICAS}}",
          "template": {
            "metadata": {
              "generateName": "centos-logtest-",
              "annotations": {
                "containerType.logging.openshift.io/${CONTAINER}": "${CONTAINER}"
              },
              "labels": "${{LABELS}}"
            },
            "spec": {
              "nodeSelector": "${{NODE_SELECTOR}}",
              "containers": [
                {
                  "env": [],
                  "image": "quay.io/openshifttest/ocp-logtest@sha256:6e2973d7d454ce412ad90e99ce584bf221866953da42858c4629873e53778606",
                  "imagePullPolicy": "IfNotPresent",
                  "name": "${CONTAINER}",
                  "resources": {},
                  "volumeMounts": [
                    {
                      "name": "config",
                      "mountPath": "/var/lib/svt"
                    }
                  ],
                  "securityContext": {
                    "allowPrivilegeEscalation": false,
                    "capabilities": {
                      "drop": [
                        "ALL"
                      ]
                    }
                  },
                  "terminationMessagePath": "/dev/termination-log"
                }
              ],
              "securityContext": {
                "runAsNonRoot": true,
                "seccompProfile": {
                  "type": "RuntimeDefault"
                }
              },
              "volumes": [
                {
                  "name": "config",
                  "configMap": {
                    "name": "${CONFIGMAP}"
                  }
                }
              ]
            }
          }
        }
      }
    ],
    "parameters": [
      {
        "name": "REPLICAS",
        "displayName": "Replicas",
        "value": "1"
      },
      {
        "name": "LABELS",
        "displayName": "labels",
        "value": "{\"run\": \"centos-logtest\", \"test\": \"centos-logtest\"}"
      },
      {
        "name": "REPLICATIONCONTROLLER",
        "displayName": "ReplicationController",
        "value": "logging-centos-logtest"
      },
      {
        "name": "CONFIGMAP",
        "displayName": "ConfigMap",
        "value": "logtest-config"
      },
      {
        "name": "CONTAINER",
        "value": "logging-centos-logtest"
      },
      {
        "name": "RATE",
        "value": "60.0"
      },
      {
        "name": "NUM_LINES",
        "value": "1000"
      },
      {
        "name": "NODE_SELECTOR",
        "displayName": "Node Selector",
        "description": "Node selector to schedule pods on specific nodes",
        "value": "{\"node-role.kubernetes.io/worker\": \"\"}"
      }
    ]
}
`)

func loggingGeneratelogLoggingPerformanceAppGeneratorJsonBytes() ([]byte, error) {
	return _loggingGeneratelogLoggingPerformanceAppGeneratorJson, nil
}

func loggingGeneratelogLoggingPerformanceAppGeneratorJson() (*asset, error) {
	bytes, err := loggingGeneratelogLoggingPerformanceAppGeneratorJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/generatelog/logging-performance-app-generator.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingGeneratelogMulti_container_json_log_templateYaml = []byte(`apiVersion: template.openshift.io/v1
kind: Template
metadata:
  name: multi-container-json-log-template
objects:
- apiVersion: v1
  data:
    ocp_logtest.cfg: "--raw --file /var/lib/svt/json.example  --text-type input --rate ${RATE}"
    json.example: "{\"message\": \"MERGE_JSON_LOG=true\", \"level\": \"debug\",\"Layer1\": \"layer1 0\", \"layer2\": {\"name\":\"Layer2 1\", \"tips\":\"Decide by PRESERVE_JSON_LOG\"}, \"StringNumber\":\"10\", \"Number\": 10,\"foo.bar\":\"Dot Item\",\"{foobar}\":\"Brace Item\",\"[foobar]\":\"Bracket Item\", \"foo:bar\":\"Colon Item\",\"foo bar\":\"Space Item\"}"
  kind: ConfigMap
  metadata:
    name: ${CMNAME}
- apiVersion: v1
  kind: ReplicationController
  metadata:
    name: ${NAME}
    labels: ${{LABELS}}
  spec:
    replicas: ${{REPLICAS}}
    template:
      metadata:
        generateName: logging-logtest-
        annotations:
          containerType.logging.openshift.io/${CONTAINER}-0: ${CONTAINER}-0
          containerType.logging.openshift.io/${CONTAINER}-1: ${CONTAINER}-1
        labels: ${{LABELS}}
      spec:
        containers:
        - env: []
          image: quay.io/openshifttest/ocp-logtest@sha256:6e2973d7d454ce412ad90e99ce584bf221866953da42858c4629873e53778606
          imagePullPolicy: IfNotPresent
          name: ${CONTAINER}-0
          resources: {}
          volumeMounts:
          - name: config
            mountPath: /var/lib/svt
          terminationMessagePath: /dev/termination-log
        - env: []
          image: quay.io/openshifttest/ocp-logtest@sha256:6e2973d7d454ce412ad90e99ce584bf221866953da42858c4629873e53778606
          imagePullPolicy: IfNotPresent
          name: ${CONTAINER}-1
          resources: {}
          volumeMounts:
          - name: config
            mountPath: /var/lib/svt
          terminationMessagePath: /dev/termination-log
        - env: []
          image: quay.io/openshifttest/ocp-logtest@sha256:6e2973d7d454ce412ad90e99ce584bf221866953da42858c4629873e53778606
          imagePullPolicy: IfNotPresent
          name: ${CONTAINER}-2
          resources: {}
          volumeMounts:
          - name: config
            mountPath: /var/lib/svt
          terminationMessagePath: /dev/termination-log
        volumes:
        - name: config
          configMap:
            name: ${CMNAME}
parameters:
- name: REPLICAS
  value: "1"
- name: LABELS
  displayName: labels
  value: "{\"run\": \"logging-logtest\", \"test\": \"logging-logtest\"}"
- name: NAME
  value: logging-logtest
- name: CMNAME
  value: multi-containers-logtest-config
- name: CONTAINER
  value: centos-logtest-container
- name: RATE
  value: "30.0"
`)

func loggingGeneratelogMulti_container_json_log_templateYamlBytes() ([]byte, error) {
	return _loggingGeneratelogMulti_container_json_log_templateYaml, nil
}

func loggingGeneratelogMulti_container_json_log_templateYaml() (*asset, error) {
	bytes, err := loggingGeneratelogMulti_container_json_log_templateYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/generatelog/multi_container_json_log_template.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingGeneratelogMultilineErrorLogYaml = []byte(`apiVersion: template.openshift.io/v1
kind: Template
metadata:
  name: multiline-log-template
objects:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: ${NAME}
    labels:
      name: multiline-log
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        name: multiline-log
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        annotations:
          capabilities: Seamless Upgrades
          containerImage: quay.io/openshifttest/multiline-log@sha256:31cabe5ffb849e79e12d7105e0d8cba68b9218d302521c8b656bad78987b0502
          support: OpenShift Logging QE
        creationTimestamp: null
        labels:
          name: multiline-log
      spec:
        containers:
        - args:
          - /run-go.sh
          command:
          - /bin/sh
          image: quay.io/openshifttest/multiline-log@sha256:31cabe5ffb849e79e12d7105e0d8cba68b9218d302521c8b656bad78987b0502
          imagePullPolicy: IfNotPresent
          name: multiline-log
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /var/lib/logging/multiline-log.cfg
            subPath: multiline-log.cfg
            name: config
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          seccompProfile:
            type: RuntimeDefault
        terminationGracePeriodSeconds: 30
        volumes:
        - configMap:
            defaultMode: 420
            name: ${NAME}
          name: config
- apiVersion: v1
  data:
    multiline-log.cfg: |
      --stream ${OUT_STREAM} --rate ${RATE} --log-type ${LOG_TYPE}
  kind: ConfigMap
  metadata:
    name: ${NAME}
parameters:
- name: NAME
  value: "multiline-log"
- name: LOG_TYPE
  value: "all"
- name: RATE
  value: "30.00"
- name: OUT_STREAM
  value: "stdout"
`)

func loggingGeneratelogMultilineErrorLogYamlBytes() ([]byte, error) {
	return _loggingGeneratelogMultilineErrorLogYaml, nil
}

func loggingGeneratelogMultilineErrorLogYaml() (*asset, error) {
	bytes, err := loggingGeneratelogMultilineErrorLogYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/generatelog/multiline-error-log.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingLogfilemetricexporterLfmeYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: logfilesmetricexporter-template
objects:
- apiVersion: "logging.openshift.io/v1alpha1"
  kind: LogFileMetricExporter
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    resources:
      limits:
        cpu: ${LIMIT_CPU}
        memory: ${LIMIT_MEMORY}
      requests:
        cpu: ${REQUEST_CPU}
        memory: ${REQUEST_MEMORY}
    tolerations: ${{TOLERATIONS}}
    nodeSelector: ${{NODE_SELECTOR}}
parameters:
- name: NAME
  value: "instance"
- name: NAMESPACE
  value: "openshift-logging"
- name: TOLERATIONS
  value: "[]"
- name: NODE_SELECTOR
  value: "{}"
- name: LIMIT_CPU
  value: "500m"
- name: LIMIT_MEMORY
  value: "256Mi"
- name: REQUEST_CPU
  value: "200m"
- name: REQUEST_MEMORY
  value: "128Mi"
`)

func loggingLogfilemetricexporterLfmeYamlBytes() ([]byte, error) {
	return _loggingLogfilemetricexporterLfmeYaml, nil
}

func loggingLogfilemetricexporterLfmeYaml() (*asset, error) {
	bytes, err := loggingLogfilemetricexporterLfmeYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/logfilemetricexporter/lfme.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingLokiLogAlertsClusterMonitoringConfigYaml = []byte(`apiVersion: template.openshift.io/v1
kind: Template
metadata:
  name: cluster-monitoring-config-temp
objects:
- apiVersion: v1
  kind: ConfigMap
  metadata:
    name: cluster-monitoring-config
    namespace: openshift-monitoring
  data:
    config.yaml: |
      enableUserWorkload: true
`)

func loggingLokiLogAlertsClusterMonitoringConfigYamlBytes() ([]byte, error) {
	return _loggingLokiLogAlertsClusterMonitoringConfigYaml, nil
}

func loggingLokiLogAlertsClusterMonitoringConfigYaml() (*asset, error) {
	bytes, err := loggingLokiLogAlertsClusterMonitoringConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/loki-log-alerts/cluster-monitoring-config.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingLokiLogAlertsLokiAppAlertingRuleTemplateYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: loki-app-alerting-rule-template
objects:
- apiVersion: loki.grafana.com/v1
  kind: AlertingRule
  metadata:
    labels:
      openshift.io/cluster-monitoring: 'true'
    name: ${ALERTING_RULE_NAME}
    namespace: ${NAMESPACE}
  spec:
    groups:
      - interval: 1m
        name: MyAppLogVolumeAlert
        rules:
          - alert: MyAppLogVolumeIsHigh
            annotations:
              description: My application has high amount of logs.
              summary: Your application project has high amount of logs.
            expr: ${ALERT_LOGQL_EXPR}
            for: 1m
            labels:
              severity: info
              project: ${NAMESPACE}
    tenantID: application
parameters:
- name: NAMESPACE
  value: "my-app-1"
- name: ALERTING_RULE_NAME
  value: "my-app-workload-alert"
- name: ALERT_LOGQL_EXPR
  value: ""
`)

func loggingLokiLogAlertsLokiAppAlertingRuleTemplateYamlBytes() ([]byte, error) {
	return _loggingLokiLogAlertsLokiAppAlertingRuleTemplateYaml, nil
}

func loggingLokiLogAlertsLokiAppAlertingRuleTemplateYaml() (*asset, error) {
	bytes, err := loggingLokiLogAlertsLokiAppAlertingRuleTemplateYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/loki-log-alerts/loki-app-alerting-rule-template.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingLokiLogAlertsLokiAppRecordingRuleTemplateYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: loki-app-recording-rule-template
objects:
- apiVersion: loki.grafana.com/v1
  kind: RecordingRule
  metadata:
    labels:
      openshift.io/cluster-monitoring: 'true'
    name: ${RECORDING_RULE_NAME}
    namespace: ${NAMESPACE}
  spec:
    groups:
      - interval: 1m
        name: HighAppLogsToLoki1m
        rules:
          - expr: >
              count_over_time({kubernetes_namespace_name="${NAMESPACE}"}[1m]) > 10
            record: 'loki:operator:applogs:rate1m'
    tenantID: application
parameters:
- name: NAMESPACE
  value: "my-app-1"
- name: RECORDING_RULE_NAME
  value: "my-app-workload-record"
`)

func loggingLokiLogAlertsLokiAppRecordingRuleTemplateYamlBytes() ([]byte, error) {
	return _loggingLokiLogAlertsLokiAppRecordingRuleTemplateYaml, nil
}

func loggingLokiLogAlertsLokiAppRecordingRuleTemplateYaml() (*asset, error) {
	bytes, err := loggingLokiLogAlertsLokiAppRecordingRuleTemplateYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/loki-log-alerts/loki-app-recording-rule-template.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingLokiLogAlertsLokiInfraAlertingRuleTemplateYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: loki-infra-alerting-rule-template
objects:
- apiVersion: loki.grafana.com/v1
  kind: AlertingRule
  metadata:
    labels:
      openshift.io/cluster-monitoring: 'true'
    name: ${ALERTING_RULE_NAME}
    namespace: ${NAMESPACE}
  spec:
    groups:
      - interval: 1m
        name: LokiOperatorLogsHigh
        rules:
          - alert: LokiOperatorLogsAreHigh
            annotations:
              description: Your Loki Operator has High amount of logs
              summary: Loki Operator Log volume is High
            expr: >
              count_over_time({kubernetes_namespace_name="${NAMESPACE}"}[1m]) > 10
            for: 1m
            labels:
              severity: info
    tenantID: infrastructure
parameters:
- name: NAMESPACE
  value: "openshift-operators-redhat"
- name: ALERTING_RULE_NAME
  value: "my-infra-workload-alert"
`)

func loggingLokiLogAlertsLokiInfraAlertingRuleTemplateYamlBytes() ([]byte, error) {
	return _loggingLokiLogAlertsLokiInfraAlertingRuleTemplateYaml, nil
}

func loggingLokiLogAlertsLokiInfraAlertingRuleTemplateYaml() (*asset, error) {
	bytes, err := loggingLokiLogAlertsLokiInfraAlertingRuleTemplateYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/loki-log-alerts/loki-infra-alerting-rule-template.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingLokiLogAlertsLokiInfraRecordingRuleTemplateYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: loki-infra-recording-rule-template
objects:
- apiVersion: loki.grafana.com/v1
  kind: RecordingRule
  metadata:
    labels:
      openshift.io/cluster-monitoring: 'true'
    name: ${RECORDING_RULE_NAME}
    namespace: ${NAMESPACE}
  spec:
    groups:
      - interval: 1m
        name: LokiOperatorLogsAreHigh1m
        rules:
          - expr: >
              count_over_time({kubernetes_namespace_name="${NAMESPACE}"}[1m]) > 10
            record: 'loki:operator:infralogs:rate1m'
    tenantID: infrastructure
parameters:
- name: NAMESPACE
  value: "openshift-operators-redhat"
- name: RECORDING_RULE_NAME
  value: "my-infra-workload-record"
`)

func loggingLokiLogAlertsLokiInfraRecordingRuleTemplateYamlBytes() ([]byte, error) {
	return _loggingLokiLogAlertsLokiInfraRecordingRuleTemplateYaml, nil
}

func loggingLokiLogAlertsLokiInfraRecordingRuleTemplateYaml() (*asset, error) {
	bytes, err := loggingLokiLogAlertsLokiInfraRecordingRuleTemplateYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/loki-log-alerts/loki-infra-recording-rule-template.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingLokiLogAlertsUserWorkloadMonitoringConfigYaml = []byte(`apiVersion: template.openshift.io/v1
kind: Template
metadata:
  name: user-workload-monitoring-config-temp
objects:
- apiVersion: v1
  kind: ConfigMap
  metadata:
    name: user-workload-monitoring-config
    namespace: openshift-user-workload-monitoring
  data:
    config.yaml: |
      alertmanager:
        enabled: true
        enableAlertmanagerConfig: true
`)

func loggingLokiLogAlertsUserWorkloadMonitoringConfigYamlBytes() ([]byte, error) {
	return _loggingLokiLogAlertsUserWorkloadMonitoringConfigYaml, nil
}

func loggingLokiLogAlertsUserWorkloadMonitoringConfigYaml() (*asset, error) {
	bytes, err := loggingLokiLogAlertsUserWorkloadMonitoringConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/loki-log-alerts/user-workload-monitoring-config.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingLokistackLokistackSimpleIpv6TlsYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: lokiStack-template
objects:
- kind: "LokiStack"
  apiVersion: "loki.grafana.com/v1"
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    networkPolicies:
      ruleSet: ${LOKISTACK_NETWORK_POLICIES_RULESET}
    limits:
      global:
        retention:
          days: 20
          streams:
          - days: 4
            priority: 1
            selector: '{kubernetes_namespace_name=~"e2e.+"}'
          - days: 1
            priority: 1
            selector: '{kubernetes_namespace_name="kube.+"}'
          - days: 15
            priority: 1
            selector: '{log_type="audit"}'
      tenants:
        application:
          retention:
            days: 1
            streams:
            - days: 4
              selector: '{kubernetes_namespace_name=~"test.+"}'
        audit:
          retention:
            days: 15
        infrastructure:
          retention:
            days: 5
            streams:
            - days: 1
              selector: '{kubernetes_namespace_name=~"openshift-cluster.+"}'
    hashRing:
      memberlist:
        enableIPv6: true
      type: memberlist
    managementState: "Managed"
    size: ${SIZE}
    storage:
      secret:
        name: ${SECRET_NAME}
        type: ${STORAGE_TYPE}
      schemas:
      - version: ${STORAGE_SCHEMA_VERSION}
        effectiveDate: ${SCHEMA_EFFECTIVE_DATE}
      tls:
        caName: ${CA_NAME}
        caKey: ${CA_KEY_NAME}
    storageClassName: ${STORAGE_CLASS}
    tenants:
      mode: "openshift-logging"
      openshift:
        adminGroups: ${{ADMIN_GROUPS}}
    rules:
      enabled: true
      selector:
        matchLabels:
          openshift.io/cluster-monitoring: "true"
      namespaceSelector:
        matchLabels:
          openshift.io/cluster-monitoring: "true"
parameters:
- name: NAME
  value: "my-loki"
- name: NAMESPACE
  value: "openshift-logging"
- name: SIZE
  value: "1x.demo"
- name: SECRET_NAME
  value: "s3-secret"
- name: STORAGE_TYPE
  value: "s3"
- name: STORAGE_CLASS
  value: "gp2"
- name: "ADMIN_GROUPS"
  value: "[]"
- name: STORAGE_SCHEMA_VERSION
  value: "v13"
- name: SCHEMA_EFFECTIVE_DATE
  value: "2023-10-15"
- name: CA_NAME
  value: ""
- name: CA_KEY_NAME
  value: "service-ca.crt"
- name: LOKISTACK_NETWORK_POLICIES_RULESET
  value: "None"
`)

func loggingLokistackLokistackSimpleIpv6TlsYamlBytes() ([]byte, error) {
	return _loggingLokistackLokistackSimpleIpv6TlsYaml, nil
}

func loggingLokistackLokistackSimpleIpv6TlsYaml() (*asset, error) {
	bytes, err := loggingLokistackLokistackSimpleIpv6TlsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/lokistack/lokistack-simple-ipv6-tls.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingLokistackLokistackSimpleIpv6Yaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: lokiStack-template
objects:
- kind: "LokiStack"
  apiVersion: "loki.grafana.com/v1"
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    networkPolicies:
      ruleSet: ${LOKISTACK_NETWORK_POLICIES_RULESET}
    limits:
      global:
        retention:
          days: 20
          streams:
          - days: 4
            priority: 1
            selector: '{kubernetes_namespace_name=~"e2e.+"}'
          - days: 1
            priority: 1
            selector: '{kubernetes_namespace_name="kube.+"}'
          - days: 15
            priority: 1
            selector: '{log_type="audit"}'
      tenants:
        application:
          retention:
            days: 1
            streams:
            - days: 4
              selector: '{kubernetes_namespace_name=~"test.+"}'
        audit:
          retention:
            days: 15
        infrastructure:
          retention:
            days: 5
            streams:
            - days: 1
              selector: '{kubernetes_namespace_name=~"openshift-cluster.+"}'
    hashRing:
      memberlist:
        enableIPv6: true
      type: memberlist
    managementState: "Managed"
    size: ${SIZE}
    storage:
      secret:
        name: ${SECRET_NAME}
        type: ${STORAGE_TYPE}
      schemas:
      - version: ${STORAGE_SCHEMA_VERSION}
        effectiveDate: ${SCHEMA_EFFECTIVE_DATE}
    storageClassName: ${STORAGE_CLASS}
    tenants:
      mode: "openshift-logging"
      openshift:
        adminGroups: ${{ADMIN_GROUPS}}
    rules:
      enabled: true
      selector:
        matchLabels:
          openshift.io/cluster-monitoring: "true"
      namespaceSelector:
        matchLabels:
          openshift.io/cluster-monitoring: "true"
parameters:
- name: NAME
  value: "my-loki"
- name: NAMESPACE
  value: "openshift-logging"
- name: SIZE
  value: "1x.demo"
- name: SECRET_NAME
  value: "s3-secret"
- name: STORAGE_TYPE
  value: "s3"
- name: STORAGE_CLASS
  value: "gp2"
- name: "ADMIN_GROUPS"
  value: "[]"
- name: STORAGE_SCHEMA_VERSION
  value: "v13"
- name: SCHEMA_EFFECTIVE_DATE
  value: "2023-10-15"
- name: LOKISTACK_NETWORK_POLICIES_RULESET
  value: "None"
`)

func loggingLokistackLokistackSimpleIpv6YamlBytes() ([]byte, error) {
	return _loggingLokistackLokistackSimpleIpv6Yaml, nil
}

func loggingLokistackLokistackSimpleIpv6Yaml() (*asset, error) {
	bytes, err := loggingLokistackLokistackSimpleIpv6YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/lokistack/lokistack-simple-ipv6.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loggingLokistackLokistackSimpleTlsYaml = []byte(`kind: Template
apiVersion: template.openshift.io/v1
metadata:
  name: lokiStack-template
objects:
- kind: "LokiStack"
  apiVersion: "loki.grafana.com/v1"
  metadata:
    name: ${NAME}
    namespace: ${NAMESPACE}
  spec:
    networkPolicies:
      ruleSet: ${LOKISTACK_NETWORK_POLICIES_RULESET}
    limits:
      global:
        retention:
          days: 20
          streams:
          - days: 4
            priority: 1
            selector: '{kubernetes_namespace_name=~"e2e.+"}'
          - days: 1
            priority: 1
            selector: '{kubernetes_namespace_name="kube.+"}'
          - days: 15
            priority: 1
            selector: '{log_type="audit"}'
      tenants:
        application:
          retention:
            days: 1
            streams:
            - days: 4
              selector: '{kubernetes_namespace_name=~"test.+"}'
        audit:
          retention:
            days: 15
        infrastructure:
          retention:
            days: 5
            streams:
            - days: 1
              selector: '{kubernetes_namespace_name=~"openshift-cluster.+"}'
    managementState: "Managed"
    size: ${SIZE}
    storage:
      secret:
        name: ${SECRET_NAME}
        type: ${STORAGE_TYPE}
      schemas:
      - version: ${STORAGE_SCHEMA_VERSION}
        effectiveDate: ${SCHEMA_EFFECTIVE_DATE}
      tls:
        caName: ${CA_NAME}
        caKey: ${CA_KEY_NAME}
    storageClassName: ${STORAGE_CLASS}
    tenants:
      mode: "openshift-logging"
      openshift:
        adminGroups: ${{ADMIN_GROUPS}}
    rules:
      enabled: true
      selector:
        matchLabels:
          openshift.io/cluster-monitoring: "true"
      namespaceSelector:
        matchLabels:
          openshift.io/cluster-monitoring: "true"
parameters:
- name: NAME
  value: "my-loki"
- name: NAMESPACE
  value: "openshift-logging"
- name: SIZE
  value: "1x.demo"
- name: SECRET_NAME
  value: "s3-secret"
- name: STORAGE_TYPE
  value: "s3"
- name: STORAGE_CLASS
  value: "gp2"
- name: "ADMIN_GROUPS"
  value: "[]"
- name: STORAGE_SCHEMA_VERSION
  value: "v13"
- name: SCHEMA_EFFECTIVE_DATE
  value: "2023-10-15"
- name: CA_NAME
  value: ""
- name: CA_KEY_NAME
  value: "service-ca.crt"
- name: LOKISTACK_NETWORK_POLICIES_RULESET
  value: "None"
`)

func loggingLokistackLokistackSimpleTlsYamlBytes() ([]byte, error) {
	return _loggingLokistackLokistackSimpleTlsYaml, nil
}

func loggingLokistackLokistackSimpleTlsYaml() (*asset, error) {
	bytes, err := loggingLokistackLokistackSimpleTlsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logging/lokistack/lokistack-simple-tls.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"fixtures.go":                                    fixturesGo,
	"logging/OWNERS":                                 loggingOwners,
	"logging/UIPlugin/UIPlugin.yaml":                 loggingUipluginUipluginYaml,
	"logging/eventrouter/eventrouter.yaml":           loggingEventrouterEventrouterYaml,
	"logging/external-log-stores/cert_generation.sh": loggingExternalLogStoresCert_generationSh,
	"logging/external-log-stores/elasticsearch/6/http/no_user/configmap.yaml":         loggingExternalLogStoresElasticsearch6HttpNo_userConfigmapYaml,
	"logging/external-log-stores/elasticsearch/6/http/no_user/deployment.yaml":        loggingExternalLogStoresElasticsearch6HttpNo_userDeploymentYaml,
	"logging/external-log-stores/elasticsearch/6/http/user_auth/configmap.yaml":       loggingExternalLogStoresElasticsearch6HttpUser_authConfigmapYaml,
	"logging/external-log-stores/elasticsearch/6/http/user_auth/deployment.yaml":      loggingExternalLogStoresElasticsearch6HttpUser_authDeploymentYaml,
	"logging/external-log-stores/elasticsearch/6/https/no_user/configmap.yaml":        loggingExternalLogStoresElasticsearch6HttpsNo_userConfigmapYaml,
	"logging/external-log-stores/elasticsearch/6/https/no_user/deployment.yaml":       loggingExternalLogStoresElasticsearch6HttpsNo_userDeploymentYaml,
	"logging/external-log-stores/elasticsearch/6/https/user_auth/configmap.yaml":      loggingExternalLogStoresElasticsearch6HttpsUser_authConfigmapYaml,
	"logging/external-log-stores/elasticsearch/6/https/user_auth/deployment.yaml":     loggingExternalLogStoresElasticsearch6HttpsUser_authDeploymentYaml,
	"logging/external-log-stores/elasticsearch/7/http/no_user/configmap.yaml":         loggingExternalLogStoresElasticsearch7HttpNo_userConfigmapYaml,
	"logging/external-log-stores/elasticsearch/7/http/no_user/deployment.yaml":        loggingExternalLogStoresElasticsearch7HttpNo_userDeploymentYaml,
	"logging/external-log-stores/elasticsearch/7/http/user_auth/configmap.yaml":       loggingExternalLogStoresElasticsearch7HttpUser_authConfigmapYaml,
	"logging/external-log-stores/elasticsearch/7/http/user_auth/deployment.yaml":      loggingExternalLogStoresElasticsearch7HttpUser_authDeploymentYaml,
	"logging/external-log-stores/elasticsearch/7/https/no_user/configmap.yaml":        loggingExternalLogStoresElasticsearch7HttpsNo_userConfigmapYaml,
	"logging/external-log-stores/elasticsearch/7/https/no_user/deployment.yaml":       loggingExternalLogStoresElasticsearch7HttpsNo_userDeploymentYaml,
	"logging/external-log-stores/elasticsearch/7/https/user_auth/configmap.yaml":      loggingExternalLogStoresElasticsearch7HttpsUser_authConfigmapYaml,
	"logging/external-log-stores/elasticsearch/7/https/user_auth/deployment.yaml":     loggingExternalLogStoresElasticsearch7HttpsUser_authDeploymentYaml,
	"logging/external-log-stores/elasticsearch/8/http/no_user/configmap.yaml":         loggingExternalLogStoresElasticsearch8HttpNo_userConfigmapYaml,
	"logging/external-log-stores/elasticsearch/8/http/no_user/deployment.yaml":        loggingExternalLogStoresElasticsearch8HttpNo_userDeploymentYaml,
	"logging/external-log-stores/elasticsearch/8/http/user_auth/configmap.yaml":       loggingExternalLogStoresElasticsearch8HttpUser_authConfigmapYaml,
	"logging/external-log-stores/elasticsearch/8/http/user_auth/deployment.yaml":      loggingExternalLogStoresElasticsearch8HttpUser_authDeploymentYaml,
	"logging/external-log-stores/elasticsearch/8/https/no_user/configmap.yaml":        loggingExternalLogStoresElasticsearch8HttpsNo_userConfigmapYaml,
	"logging/external-log-stores/elasticsearch/8/https/no_user/deployment.yaml":       loggingExternalLogStoresElasticsearch8HttpsNo_userDeploymentYaml,
	"logging/external-log-stores/elasticsearch/8/https/user_auth/configmap.yaml":      loggingExternalLogStoresElasticsearch8HttpsUser_authConfigmapYaml,
	"logging/external-log-stores/elasticsearch/8/https/user_auth/deployment.yaml":     loggingExternalLogStoresElasticsearch8HttpsUser_authDeploymentYaml,
	"logging/external-log-stores/fluentd/insecure/configmap.yaml":                     loggingExternalLogStoresFluentdInsecureConfigmapYaml,
	"logging/external-log-stores/fluentd/insecure/deployment.yaml":                    loggingExternalLogStoresFluentdInsecureDeploymentYaml,
	"logging/external-log-stores/fluentd/insecure/http-configmap.yaml":                loggingExternalLogStoresFluentdInsecureHttpConfigmapYaml,
	"logging/external-log-stores/fluentd/secure/cm-mtls-share.yaml":                   loggingExternalLogStoresFluentdSecureCmMtlsShareYaml,
	"logging/external-log-stores/fluentd/secure/cm-mtls.yaml":                         loggingExternalLogStoresFluentdSecureCmMtlsYaml,
	"logging/external-log-stores/fluentd/secure/cm-serverauth-share.yaml":             loggingExternalLogStoresFluentdSecureCmServerauthShareYaml,
	"logging/external-log-stores/fluentd/secure/cm-serverauth.yaml":                   loggingExternalLogStoresFluentdSecureCmServerauthYaml,
	"logging/external-log-stores/fluentd/secure/deployment.yaml":                      loggingExternalLogStoresFluentdSecureDeploymentYaml,
	"logging/external-log-stores/fluentd/secure/http-cm-mtls.yaml":                    loggingExternalLogStoresFluentdSecureHttpCmMtlsYaml,
	"logging/external-log-stores/fluentd/secure/http-cm-serverauth.yaml":              loggingExternalLogStoresFluentdSecureHttpCmServerauthYaml,
	"logging/external-log-stores/kafka/amqstreams/kafka-no-auth-cluster.yaml":         loggingExternalLogStoresKafkaAmqstreamsKafkaNoAuthClusterYaml,
	"logging/external-log-stores/kafka/amqstreams/kafka-no-auth-consumer-job.yaml":    loggingExternalLogStoresKafkaAmqstreamsKafkaNoAuthConsumerJobYaml,
	"logging/external-log-stores/kafka/amqstreams/kafka-sasl-cluster.yaml":            loggingExternalLogStoresKafkaAmqstreamsKafkaSaslClusterYaml,
	"logging/external-log-stores/kafka/amqstreams/kafka-sasl-consumer-job.yaml":       loggingExternalLogStoresKafkaAmqstreamsKafkaSaslConsumerJobYaml,
	"logging/external-log-stores/kafka/amqstreams/kafka-sasl-consumers-config.yaml":   loggingExternalLogStoresKafkaAmqstreamsKafkaSaslConsumersConfigYaml,
	"logging/external-log-stores/kafka/amqstreams/kafka-sasl-user.yaml":               loggingExternalLogStoresKafkaAmqstreamsKafkaSaslUserYaml,
	"logging/external-log-stores/kafka/amqstreams/kafka-topic.yaml":                   loggingExternalLogStoresKafkaAmqstreamsKafkaTopicYaml,
	"logging/external-log-stores/kafka/cert_generation.sh":                            loggingExternalLogStoresKafkaCert_generationSh,
	"logging/external-log-stores/kafka/kafka-rbac.yaml":                               loggingExternalLogStoresKafkaKafkaRbacYaml,
	"logging/external-log-stores/kafka/kafka-svc.yaml":                                loggingExternalLogStoresKafkaKafkaSvcYaml,
	"logging/external-log-stores/kafka/plaintext-ssl/consumer-configmap.yaml":         loggingExternalLogStoresKafkaPlaintextSslConsumerConfigmapYaml,
	"logging/external-log-stores/kafka/plaintext-ssl/kafka-configmap.yaml":            loggingExternalLogStoresKafkaPlaintextSslKafkaConfigmapYaml,
	"logging/external-log-stores/kafka/plaintext-ssl/kafka-consumer-deployment.yaml":  loggingExternalLogStoresKafkaPlaintextSslKafkaConsumerDeploymentYaml,
	"logging/external-log-stores/kafka/plaintext-ssl/kafka-statefulset.yaml":          loggingExternalLogStoresKafkaPlaintextSslKafkaStatefulsetYaml,
	"logging/external-log-stores/kafka/sasl-plaintext/consumer-configmap.yaml":        loggingExternalLogStoresKafkaSaslPlaintextConsumerConfigmapYaml,
	"logging/external-log-stores/kafka/sasl-plaintext/kafka-configmap.yaml":           loggingExternalLogStoresKafkaSaslPlaintextKafkaConfigmapYaml,
	"logging/external-log-stores/kafka/sasl-plaintext/kafka-consumer-deployment.yaml": loggingExternalLogStoresKafkaSaslPlaintextKafkaConsumerDeploymentYaml,
	"logging/external-log-stores/kafka/sasl-plaintext/kafka-statefulset.yaml":         loggingExternalLogStoresKafkaSaslPlaintextKafkaStatefulsetYaml,
	"logging/external-log-stores/kafka/sasl-ssl/consumer-configmap.yaml":              loggingExternalLogStoresKafkaSaslSslConsumerConfigmapYaml,
	"logging/external-log-stores/kafka/sasl-ssl/kafka-configmap.yaml":                 loggingExternalLogStoresKafkaSaslSslKafkaConfigmapYaml,
	"logging/external-log-stores/kafka/sasl-ssl/kafka-consumer-deployment.yaml":       loggingExternalLogStoresKafkaSaslSslKafkaConsumerDeploymentYaml,
	"logging/external-log-stores/kafka/sasl-ssl/kafka-statefulset.yaml":               loggingExternalLogStoresKafkaSaslSslKafkaStatefulsetYaml,
	"logging/external-log-stores/kafka/zookeeper/configmap-ssl.yaml":                  loggingExternalLogStoresKafkaZookeeperConfigmapSslYaml,
	"logging/external-log-stores/kafka/zookeeper/configmap.yaml":                      loggingExternalLogStoresKafkaZookeeperConfigmapYaml,
	"logging/external-log-stores/kafka/zookeeper/zookeeper-statefulset.yaml":          loggingExternalLogStoresKafkaZookeeperZookeeperStatefulsetYaml,
	"logging/external-log-stores/kafka/zookeeper/zookeeper-svc.yaml":                  loggingExternalLogStoresKafkaZookeeperZookeeperSvcYaml,
	"logging/external-log-stores/loki/loki-configmap.yaml":                            loggingExternalLogStoresLokiLokiConfigmapYaml,
	"logging/external-log-stores/loki/loki-deployment.yaml":                           loggingExternalLogStoresLokiLokiDeploymentYaml,
	"logging/external-log-stores/otel/otel-collector.yaml":                            loggingExternalLogStoresOtelOtelCollectorYaml,
	"logging/external-log-stores/rsyslog/insecure/configmap.yaml":                     loggingExternalLogStoresRsyslogInsecureConfigmapYaml,
	"logging/external-log-stores/rsyslog/insecure/deployment.yaml":                    loggingExternalLogStoresRsyslogInsecureDeploymentYaml,
	"logging/external-log-stores/rsyslog/insecure/svc.yaml":                           loggingExternalLogStoresRsyslogInsecureSvcYaml,
	"logging/external-log-stores/rsyslog/secure/configmap.yaml":                       loggingExternalLogStoresRsyslogSecureConfigmapYaml,
	"logging/external-log-stores/rsyslog/secure/deployment.yaml":                      loggingExternalLogStoresRsyslogSecureDeploymentYaml,
	"logging/external-log-stores/rsyslog/secure/svc.yaml":                             loggingExternalLogStoresRsyslogSecureSvcYaml,
	"logging/external-log-stores/splunk/route-edge_splunk_template.yaml":              loggingExternalLogStoresSplunkRouteEdge_splunk_templateYaml,
	"logging/external-log-stores/splunk/route-passthrough_splunk_template.yaml":       loggingExternalLogStoresSplunkRoutePassthrough_splunk_templateYaml,
	"logging/external-log-stores/splunk/secret_splunk_template.yaml":                  loggingExternalLogStoresSplunkSecret_splunk_templateYaml,
	"logging/external-log-stores/splunk/secret_tls_passphrase_splunk_template.yaml":   loggingExternalLogStoresSplunkSecret_tls_passphrase_splunk_templateYaml,
	"logging/external-log-stores/splunk/secret_tls_splunk_template.yaml":              loggingExternalLogStoresSplunkSecret_tls_splunk_templateYaml,
	"logging/external-log-stores/splunk/statefulset_splunk-8.2_template.yaml":         loggingExternalLogStoresSplunkStatefulset_splunk82_templateYaml,
	"logging/external-log-stores/splunk/statefulset_splunk-9.0_template.yaml":         loggingExternalLogStoresSplunkStatefulset_splunk90_templateYaml,
	"logging/generatelog/42981.yaml":                                                  loggingGeneratelog42981Yaml,
	"logging/generatelog/container_json_log_template.json":                            loggingGeneratelogContainer_json_log_templateJson,
	"logging/generatelog/container_json_log_template_unannoted.json":                  loggingGeneratelogContainer_json_log_template_unannotedJson,
	"logging/generatelog/container_non_json_log_template.json":                        loggingGeneratelogContainer_non_json_log_templateJson,
	"logging/generatelog/logging-performance-app-generator.json":                      loggingGeneratelogLoggingPerformanceAppGeneratorJson,
	"logging/generatelog/multi_container_json_log_template.yaml":                      loggingGeneratelogMulti_container_json_log_templateYaml,
	"logging/generatelog/multiline-error-log.yaml":                                    loggingGeneratelogMultilineErrorLogYaml,
	"logging/logfilemetricexporter/lfme.yaml":                                         loggingLogfilemetricexporterLfmeYaml,
	"logging/loki-log-alerts/cluster-monitoring-config.yaml":                          loggingLokiLogAlertsClusterMonitoringConfigYaml,
	"logging/loki-log-alerts/loki-app-alerting-rule-template.yaml":                    loggingLokiLogAlertsLokiAppAlertingRuleTemplateYaml,
	"logging/loki-log-alerts/loki-app-recording-rule-template.yaml":                   loggingLokiLogAlertsLokiAppRecordingRuleTemplateYaml,
	"logging/loki-log-alerts/loki-infra-alerting-rule-template.yaml":                  loggingLokiLogAlertsLokiInfraAlertingRuleTemplateYaml,
	"logging/loki-log-alerts/loki-infra-recording-rule-template.yaml":                 loggingLokiLogAlertsLokiInfraRecordingRuleTemplateYaml,
	"logging/loki-log-alerts/user-workload-monitoring-config.yaml":                    loggingLokiLogAlertsUserWorkloadMonitoringConfigYaml,
	"logging/lokistack/lokistack-simple-ipv6-tls.yaml":                                loggingLokistackLokistackSimpleIpv6TlsYaml,
	"logging/lokistack/lokistack-simple-ipv6.yaml":                                    loggingLokistackLokistackSimpleIpv6Yaml,
	"logging/lokistack/lokistack-simple-tls.yaml":                                     loggingLokistackLokistackSimpleTlsYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"fixtures.go": {fixturesGo, map[string]*bintree{}},
	"logging": {nil, map[string]*bintree{
		"OWNERS": {loggingOwners, map[string]*bintree{}},
		"UIPlugin": {nil, map[string]*bintree{
			"UIPlugin.yaml": {loggingUipluginUipluginYaml, map[string]*bintree{}},
		}},
		"eventrouter": {nil, map[string]*bintree{
			"eventrouter.yaml": {loggingEventrouterEventrouterYaml, map[string]*bintree{}},
		}},
		"external-log-stores": {nil, map[string]*bintree{
			"cert_generation.sh": {loggingExternalLogStoresCert_generationSh, map[string]*bintree{}},
			"elasticsearch": {nil, map[string]*bintree{
				"6": {nil, map[string]*bintree{
					"http": {nil, map[string]*bintree{
						"no_user": {nil, map[string]*bintree{
							"configmap.yaml":  {loggingExternalLogStoresElasticsearch6HttpNo_userConfigmapYaml, map[string]*bintree{}},
							"deployment.yaml": {loggingExternalLogStoresElasticsearch6HttpNo_userDeploymentYaml, map[string]*bintree{}},
						}},
						"user_auth": {nil, map[string]*bintree{
							"configmap.yaml":  {loggingExternalLogStoresElasticsearch6HttpUser_authConfigmapYaml, map[string]*bintree{}},
							"deployment.yaml": {loggingExternalLogStoresElasticsearch6HttpUser_authDeploymentYaml, map[string]*bintree{}},
						}},
					}},
					"https": {nil, map[string]*bintree{
						"no_user": {nil, map[string]*bintree{
							"configmap.yaml":  {loggingExternalLogStoresElasticsearch6HttpsNo_userConfigmapYaml, map[string]*bintree{}},
							"deployment.yaml": {loggingExternalLogStoresElasticsearch6HttpsNo_userDeploymentYaml, map[string]*bintree{}},
						}},
						"user_auth": {nil, map[string]*bintree{
							"configmap.yaml":  {loggingExternalLogStoresElasticsearch6HttpsUser_authConfigmapYaml, map[string]*bintree{}},
							"deployment.yaml": {loggingExternalLogStoresElasticsearch6HttpsUser_authDeploymentYaml, map[string]*bintree{}},
						}},
					}},
				}},
				"7": {nil, map[string]*bintree{
					"http": {nil, map[string]*bintree{
						"no_user": {nil, map[string]*bintree{
							"configmap.yaml":  {loggingExternalLogStoresElasticsearch7HttpNo_userConfigmapYaml, map[string]*bintree{}},
							"deployment.yaml": {loggingExternalLogStoresElasticsearch7HttpNo_userDeploymentYaml, map[string]*bintree{}},
						}},
						"user_auth": {nil, map[string]*bintree{
							"configmap.yaml":  {loggingExternalLogStoresElasticsearch7HttpUser_authConfigmapYaml, map[string]*bintree{}},
							"deployment.yaml": {loggingExternalLogStoresElasticsearch7HttpUser_authDeploymentYaml, map[string]*bintree{}},
						}},
					}},
					"https": {nil, map[string]*bintree{
						"no_user": {nil, map[string]*bintree{
							"configmap.yaml":  {loggingExternalLogStoresElasticsearch7HttpsNo_userConfigmapYaml, map[string]*bintree{}},
							"deployment.yaml": {loggingExternalLogStoresElasticsearch7HttpsNo_userDeploymentYaml, map[string]*bintree{}},
						}},
						"user_auth": {nil, map[string]*bintree{
							"configmap.yaml":  {loggingExternalLogStoresElasticsearch7HttpsUser_authConfigmapYaml, map[string]*bintree{}},
							"deployment.yaml": {loggingExternalLogStoresElasticsearch7HttpsUser_authDeploymentYaml, map[string]*bintree{}},
						}},
					}},
				}},
				"8": {nil, map[string]*bintree{
					"http": {nil, map[string]*bintree{
						"no_user": {nil, map[string]*bintree{
							"configmap.yaml":  {loggingExternalLogStoresElasticsearch8HttpNo_userConfigmapYaml, map[string]*bintree{}},
							"deployment.yaml": {loggingExternalLogStoresElasticsearch8HttpNo_userDeploymentYaml, map[string]*bintree{}},
						}},
						"user_auth": {nil, map[string]*bintree{
							"configmap.yaml":  {loggingExternalLogStoresElasticsearch8HttpUser_authConfigmapYaml, map[string]*bintree{}},
							"deployment.yaml": {loggingExternalLogStoresElasticsearch8HttpUser_authDeploymentYaml, map[string]*bintree{}},
						}},
					}},
					"https": {nil, map[string]*bintree{
						"no_user": {nil, map[string]*bintree{
							"configmap.yaml":  {loggingExternalLogStoresElasticsearch8HttpsNo_userConfigmapYaml, map[string]*bintree{}},
							"deployment.yaml": {loggingExternalLogStoresElasticsearch8HttpsNo_userDeploymentYaml, map[string]*bintree{}},
						}},
						"user_auth": {nil, map[string]*bintree{
							"configmap.yaml":  {loggingExternalLogStoresElasticsearch8HttpsUser_authConfigmapYaml, map[string]*bintree{}},
							"deployment.yaml": {loggingExternalLogStoresElasticsearch8HttpsUser_authDeploymentYaml, map[string]*bintree{}},
						}},
					}},
				}},
			}},
			"fluentd": {nil, map[string]*bintree{
				"insecure": {nil, map[string]*bintree{
					"configmap.yaml":      {loggingExternalLogStoresFluentdInsecureConfigmapYaml, map[string]*bintree{}},
					"deployment.yaml":     {loggingExternalLogStoresFluentdInsecureDeploymentYaml, map[string]*bintree{}},
					"http-configmap.yaml": {loggingExternalLogStoresFluentdInsecureHttpConfigmapYaml, map[string]*bintree{}},
				}},
				"secure": {nil, map[string]*bintree{
					"cm-mtls-share.yaml":       {loggingExternalLogStoresFluentdSecureCmMtlsShareYaml, map[string]*bintree{}},
					"cm-mtls.yaml":             {loggingExternalLogStoresFluentdSecureCmMtlsYaml, map[string]*bintree{}},
					"cm-serverauth-share.yaml": {loggingExternalLogStoresFluentdSecureCmServerauthShareYaml, map[string]*bintree{}},
					"cm-serverauth.yaml":       {loggingExternalLogStoresFluentdSecureCmServerauthYaml, map[string]*bintree{}},
					"deployment.yaml":          {loggingExternalLogStoresFluentdSecureDeploymentYaml, map[string]*bintree{}},
					"http-cm-mtls.yaml":        {loggingExternalLogStoresFluentdSecureHttpCmMtlsYaml, map[string]*bintree{}},
					"http-cm-serverauth.yaml":  {loggingExternalLogStoresFluentdSecureHttpCmServerauthYaml, map[string]*bintree{}},
				}},
			}},
			"kafka": {nil, map[string]*bintree{
				"amqstreams": {nil, map[string]*bintree{
					"kafka-no-auth-cluster.yaml":       {loggingExternalLogStoresKafkaAmqstreamsKafkaNoAuthClusterYaml, map[string]*bintree{}},
					"kafka-no-auth-consumer-job.yaml":  {loggingExternalLogStoresKafkaAmqstreamsKafkaNoAuthConsumerJobYaml, map[string]*bintree{}},
					"kafka-sasl-cluster.yaml":          {loggingExternalLogStoresKafkaAmqstreamsKafkaSaslClusterYaml, map[string]*bintree{}},
					"kafka-sasl-consumer-job.yaml":     {loggingExternalLogStoresKafkaAmqstreamsKafkaSaslConsumerJobYaml, map[string]*bintree{}},
					"kafka-sasl-consumers-config.yaml": {loggingExternalLogStoresKafkaAmqstreamsKafkaSaslConsumersConfigYaml, map[string]*bintree{}},
					"kafka-sasl-user.yaml":             {loggingExternalLogStoresKafkaAmqstreamsKafkaSaslUserYaml, map[string]*bintree{}},
					"kafka-topic.yaml":                 {loggingExternalLogStoresKafkaAmqstreamsKafkaTopicYaml, map[string]*bintree{}},
				}},
				"cert_generation.sh": {loggingExternalLogStoresKafkaCert_generationSh, map[string]*bintree{}},
				"kafka-rbac.yaml":    {loggingExternalLogStoresKafkaKafkaRbacYaml, map[string]*bintree{}},
				"kafka-svc.yaml":     {loggingExternalLogStoresKafkaKafkaSvcYaml, map[string]*bintree{}},
				"plaintext-ssl": {nil, map[string]*bintree{
					"consumer-configmap.yaml":        {loggingExternalLogStoresKafkaPlaintextSslConsumerConfigmapYaml, map[string]*bintree{}},
					"kafka-configmap.yaml":           {loggingExternalLogStoresKafkaPlaintextSslKafkaConfigmapYaml, map[string]*bintree{}},
					"kafka-consumer-deployment.yaml": {loggingExternalLogStoresKafkaPlaintextSslKafkaConsumerDeploymentYaml, map[string]*bintree{}},
					"kafka-statefulset.yaml":         {loggingExternalLogStoresKafkaPlaintextSslKafkaStatefulsetYaml, map[string]*bintree{}},
				}},
				"sasl-plaintext": {nil, map[string]*bintree{
					"consumer-configmap.yaml":        {loggingExternalLogStoresKafkaSaslPlaintextConsumerConfigmapYaml, map[string]*bintree{}},
					"kafka-configmap.yaml":           {loggingExternalLogStoresKafkaSaslPlaintextKafkaConfigmapYaml, map[string]*bintree{}},
					"kafka-consumer-deployment.yaml": {loggingExternalLogStoresKafkaSaslPlaintextKafkaConsumerDeploymentYaml, map[string]*bintree{}},
					"kafka-statefulset.yaml":         {loggingExternalLogStoresKafkaSaslPlaintextKafkaStatefulsetYaml, map[string]*bintree{}},
				}},
				"sasl-ssl": {nil, map[string]*bintree{
					"consumer-configmap.yaml":        {loggingExternalLogStoresKafkaSaslSslConsumerConfigmapYaml, map[string]*bintree{}},
					"kafka-configmap.yaml":           {loggingExternalLogStoresKafkaSaslSslKafkaConfigmapYaml, map[string]*bintree{}},
					"kafka-consumer-deployment.yaml": {loggingExternalLogStoresKafkaSaslSslKafkaConsumerDeploymentYaml, map[string]*bintree{}},
					"kafka-statefulset.yaml":         {loggingExternalLogStoresKafkaSaslSslKafkaStatefulsetYaml, map[string]*bintree{}},
				}},
				"zookeeper": {nil, map[string]*bintree{
					"configmap-ssl.yaml":         {loggingExternalLogStoresKafkaZookeeperConfigmapSslYaml, map[string]*bintree{}},
					"configmap.yaml":             {loggingExternalLogStoresKafkaZookeeperConfigmapYaml, map[string]*bintree{}},
					"zookeeper-statefulset.yaml": {loggingExternalLogStoresKafkaZookeeperZookeeperStatefulsetYaml, map[string]*bintree{}},
					"zookeeper-svc.yaml":         {loggingExternalLogStoresKafkaZookeeperZookeeperSvcYaml, map[string]*bintree{}},
				}},
			}},
			"loki": {nil, map[string]*bintree{
				"loki-configmap.yaml":  {loggingExternalLogStoresLokiLokiConfigmapYaml, map[string]*bintree{}},
				"loki-deployment.yaml": {loggingExternalLogStoresLokiLokiDeploymentYaml, map[string]*bintree{}},
			}},
			"otel": {nil, map[string]*bintree{
				"otel-collector.yaml": {loggingExternalLogStoresOtelOtelCollectorYaml, map[string]*bintree{}},
			}},
			"rsyslog": {nil, map[string]*bintree{
				"insecure": {nil, map[string]*bintree{
					"configmap.yaml":  {loggingExternalLogStoresRsyslogInsecureConfigmapYaml, map[string]*bintree{}},
					"deployment.yaml": {loggingExternalLogStoresRsyslogInsecureDeploymentYaml, map[string]*bintree{}},
					"svc.yaml":        {loggingExternalLogStoresRsyslogInsecureSvcYaml, map[string]*bintree{}},
				}},
				"secure": {nil, map[string]*bintree{
					"configmap.yaml":  {loggingExternalLogStoresRsyslogSecureConfigmapYaml, map[string]*bintree{}},
					"deployment.yaml": {loggingExternalLogStoresRsyslogSecureDeploymentYaml, map[string]*bintree{}},
					"svc.yaml":        {loggingExternalLogStoresRsyslogSecureSvcYaml, map[string]*bintree{}},
				}},
			}},
			"splunk": {nil, map[string]*bintree{
				"route-edge_splunk_template.yaml":            {loggingExternalLogStoresSplunkRouteEdge_splunk_templateYaml, map[string]*bintree{}},
				"route-passthrough_splunk_template.yaml":     {loggingExternalLogStoresSplunkRoutePassthrough_splunk_templateYaml, map[string]*bintree{}},
				"secret_splunk_template.yaml":                {loggingExternalLogStoresSplunkSecret_splunk_templateYaml, map[string]*bintree{}},
				"secret_tls_passphrase_splunk_template.yaml": {loggingExternalLogStoresSplunkSecret_tls_passphrase_splunk_templateYaml, map[string]*bintree{}},
				"secret_tls_splunk_template.yaml":            {loggingExternalLogStoresSplunkSecret_tls_splunk_templateYaml, map[string]*bintree{}},
				"statefulset_splunk-8.2_template.yaml":       {loggingExternalLogStoresSplunkStatefulset_splunk82_templateYaml, map[string]*bintree{}},
				"statefulset_splunk-9.0_template.yaml":       {loggingExternalLogStoresSplunkStatefulset_splunk90_templateYaml, map[string]*bintree{}},
			}},
		}},
		"generatelog": {nil, map[string]*bintree{
			"42981.yaml":                                 {loggingGeneratelog42981Yaml, map[string]*bintree{}},
			"container_json_log_template.json":           {loggingGeneratelogContainer_json_log_templateJson, map[string]*bintree{}},
			"container_json_log_template_unannoted.json": {loggingGeneratelogContainer_json_log_template_unannotedJson, map[string]*bintree{}},
			"container_non_json_log_template.json":       {loggingGeneratelogContainer_non_json_log_templateJson, map[string]*bintree{}},
			"logging-performance-app-generator.json":     {loggingGeneratelogLoggingPerformanceAppGeneratorJson, map[string]*bintree{}},
			"multi_container_json_log_template.yaml":     {loggingGeneratelogMulti_container_json_log_templateYaml, map[string]*bintree{}},
			"multiline-error-log.yaml":                   {loggingGeneratelogMultilineErrorLogYaml, map[string]*bintree{}},
		}},
		"logfilemetricexporter": {nil, map[string]*bintree{
			"lfme.yaml": {loggingLogfilemetricexporterLfmeYaml, map[string]*bintree{}},
		}},
		"loki-log-alerts": {nil, map[string]*bintree{
			"cluster-monitoring-config.yaml":          {loggingLokiLogAlertsClusterMonitoringConfigYaml, map[string]*bintree{}},
			"loki-app-alerting-rule-template.yaml":    {loggingLokiLogAlertsLokiAppAlertingRuleTemplateYaml, map[string]*bintree{}},
			"loki-app-recording-rule-template.yaml":   {loggingLokiLogAlertsLokiAppRecordingRuleTemplateYaml, map[string]*bintree{}},
			"loki-infra-alerting-rule-template.yaml":  {loggingLokiLogAlertsLokiInfraAlertingRuleTemplateYaml, map[string]*bintree{}},
			"loki-infra-recording-rule-template.yaml": {loggingLokiLogAlertsLokiInfraRecordingRuleTemplateYaml, map[string]*bintree{}},
			"user-workload-monitoring-config.yaml":    {loggingLokiLogAlertsUserWorkloadMonitoringConfigYaml, map[string]*bintree{}},
		}},
		"lokistack": {nil, map[string]*bintree{
			"lokistack-simple-ipv6-tls.yaml": {loggingLokistackLokistackSimpleIpv6TlsYaml, map[string]*bintree{}},
			"lokistack-simple-ipv6.yaml":     {loggingLokistackLokistackSimpleIpv6Yaml, map[string]*bintree{}},
			"lokistack-simple-tls.yaml":      {loggingLokistackLokistackSimpleTlsYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
