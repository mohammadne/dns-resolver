package dns

import (
	"net"

	"github.com/miekg/dns"
	"go.uber.org/zap"
)

func (handler *Server) handleDNS(conn *net.UDPConn) {
	buffer := make([]byte, 1024)
	rlen, remote, err := conn.ReadFromUDP(buffer[:])
	if err != nil {
		handler.logger.Error("Error reading from UDP", zap.Error(err))
		return
	}
	buffer = buffer[:rlen]

	msg := new(dns.Msg)
	if err := msg.Unpack(buffer); err != nil {
		handler.logger.Error("Error parsing DNS query", zap.Error(err))
		return
	}

	// Process the DNS query and prepare the response
	response := new(dns.Msg)
	response.SetReply(msg)

	sendResponse := func() {
		// Convert the DNS response to a byte slice
		responseBytes, err := response.Pack()
		if err != nil {
			errMessage := "Error packing DNS response"
			handler.logger.Error(errMessage, zap.Error(err))
			return
		}

		// Send a response back to the client
		if _, err := conn.WriteToUDP(responseBytes, remote); err != nil {
			handler.logger.Error("Error sending response", zap.Error(err))
		}
	}

	if len(msg.Question) > 0 {
		question := msg.Question[0]

		ip, ok := handler.records[question.Name]
		if !ok {
			errMessage := "Domain mapping doesn't exists"
			handler.logger.Error(errMessage, zap.String("domain", question.Name))
			sendResponse()
			return
		}

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

	// Set the DNS response message ID to match the query ID
	response.Id = msg.Id

	sendResponse()
}
