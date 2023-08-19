package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mohammadne/dns-resolver/pkg/logger"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HTTP Handlers Suite")
}

var testServer *Server

// Perform setup steps before the entire test suite
var _ = BeforeSuite(func() {
	testServer = New(logger.NewNoop())
})

var _ = Describe("Handlers", func() {
	It("random endpoint", func() {
		request := httptest.NewRequest(http.MethodGet, "/random-endpoint", nil)
		response, _ := testServer.app.Test(request)

		// Check the expected response code
		Expect(response.StatusCode).Should(Equal(http.StatusNotFound))
	})

	It("liveness", func() {
		request := httptest.NewRequest(http.MethodGet, "/healthz/liveness", nil)
		response, _ := testServer.app.Test(request)

		// Check the expected response code
		Expect(response.StatusCode).Should(Equal(http.StatusOK))
	})

	It("readiness", func() {
		request := httptest.NewRequest(http.MethodGet, "/healthz/readiness", nil)
		response, _ := testServer.app.Test(request)

		// Check the expected response code
		Expect(response.StatusCode).Should(Equal(http.StatusOK))
	})
})
