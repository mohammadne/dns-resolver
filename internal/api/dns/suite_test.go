package dns

import (
	"testing"

	"github.com/mohammadne/dns-resolver/pkg/logger"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var testServer *Server

// Perform setup steps before the entire test suite
var _ = BeforeSuite(func() {
	testServer = New(logger.NewNoop())
})

func TestHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DNS Suite")
}
