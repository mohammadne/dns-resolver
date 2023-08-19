package config

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("Load Config", func() {
	It("load config without print", func() {
		cfg := Load(false)
		Expect(cfg).NotTo(BeNil())
	})

	It("load config with print", func() {
		cfg := Load(true)
		Expect(cfg).NotTo(BeNil())
	})
})
