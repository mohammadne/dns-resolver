package dns

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DNS Server", func() {
	It("records", func() {
		Expect(len(testServer.records)).ToNot(BeZero())
		Expect(testServer.records).To(HaveKeyWithValue("example10.com.", "172.1.0.10"))
		Expect(testServer.records).NotTo(HaveKey("example100.com."))
	})
})
