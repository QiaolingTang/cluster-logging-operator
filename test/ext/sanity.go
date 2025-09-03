package ext

import (
	g "github.com/onsi/ginkgo/v2"
	o "github.com/onsi/gomega"
)

var _ = g.Describe("[Jira:Logging][sig-openshift-logging] sanity test", func() {
	g.It("should always pass [Suite:logging/parallel]", func() {
		o.Expect(true).To(o.BeTrue())
	})
})
