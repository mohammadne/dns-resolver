package dns

import (
	"net"

	"github.com/miekg/dns"
	"go.uber.org/zap"
)

func (handler *Server) handleDNS(request []byte) ([]byte, error) {
	msg := new(dns.Msg)
	if err := msg.Unpack(request); err != nil {
		handler.logger.Error("Error parsing DNS query", zap.Error(err))
		return nil, err
	}

	// Process the DNS query and prepare the response
	response := new(dns.Msg)
	response.SetReply(msg)

	if len(msg.Question) > 0 {
		question := msg.Question[0]

		ip, ok := handler.records[question.Name]
		if !ok {
			handler.setNXDomainResponse(response)
			handler.logger.Error("Domain mapping doesn't exists", zap.String("domain", question.Name))
			return handler.packResponse(response)
		}

		handler.setAResponse(response, &question, ip)
	}

	// Set the DNS response message ID to match the query ID
	response.Id = msg.Id

	return handler.packResponse(response)
}

func (handler *Server) setNXDomainResponse(response *dns.Msg) {
	response.Authoritative = true
	response.RecursionAvailable = true
	response.Rcode = dns.RcodeNameError
}

func (handler *Server) setAResponse(response *dns.Msg, question *dns.Question, ip string) {
	// Assuming the query is for an A record
	if question.Qtype == dns.TypeA {
		// Create an A record for the domain name
		aRecord := &dns.A{
			Hdr: dns.RR_Header{
				Name:   question.Name,
				Rrtype: dns.TypeA,
				Class:  dns.ClassINET,
				Ttl:    60, // Time to live in seconds
			},
			A: net.ParseIP(ip), // Set the IP address
		}

		// Add the A record to the DNS response
		response.Answer = []dns.RR{aRecord}
	}
}

func (handler *Server) packResponse(response *dns.Msg) ([]byte, error) {
	// Convert the DNS response to a byte slice
	responseBytes, err := response.Pack()
	if err != nil {
		errMessage := "Error packing DNS response"
		handler.logger.Error(errMessage, zap.Error(err))
		return nil, err
	}

	return responseBytes, nil
}
