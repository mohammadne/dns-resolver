package dns

import (
	"fmt"

	"github.com/miekg/dns"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DNS Handler", func() {
	request := func(domain string) (*dns.Msg, error) {
		msg := dns.Msg{} // Create a new DNS message
		msg.SetQuestion(dns.Fqdn(domain), dns.TypeA)

		// Convert DNS message to byte slice
		msgBytes, err := msg.Pack()
		if err != nil {
			fmt.Printf("Failed to pack DNS message: \n%v", err)
			return nil, err
		}

		responseBytes, err := testServer.handleDNS(msgBytes)
		if err != nil {
			fmt.Printf("Error processing DNS request: \n%v", err)
			return nil, err
		}

		// Parse the DNS response
		response := dns.Msg{}
		if err = response.Unpack(responseBytes); err != nil {
			fmt.Printf("Failed to unpack DNS response: \n%v", err)
			return nil, err
		}

		return &response, nil
	}

	It("should handle successful but non-existing dns requests", func() {
		response, err := request("example.com")
		Expect(err).ToNot(HaveOccurred())

		// Check if the response is NXDOMAIN
		Expect(response.Rcode).To(Equal(dns.RcodeNameError), "Response is not NXDOMAIN")
	})

	testCases := []struct {
		description string
		domain      string
	}{
		{
			description: "should handle successful and existing dns requests",
			domain:      "example1.com",
		},
		{
			description: "should handle empty dns request",
			domain:      "",
		},
		{
			description: "should handle invalid dns request",
			domain:      "I love F1",
		},
	}

	for _, testCase := range testCases {
		It(testCase.description, func() {
			response, err := request(testCase.domain)
			Expect(err).ToNot(HaveOccurred())

			// Print the IP addresses from the DNS response
			for _, answer := range response.Answer {
				if a, ok := answer.(*dns.A); ok {
					fmt.Printf("IP address: %s\n", a.A.String())
				} else {
					Fail("Mismatch in expected IPs")
				}
			}
		})
	}
})
