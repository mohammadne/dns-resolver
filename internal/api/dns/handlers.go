package dns

import (
	"net"
	"net/url"
	"strings"

	"go.uber.org/zap"
)

func (handler *Server) handleDNS(conn *net.UDPConn) {
	message := make([]byte, 1024)
	rlen, remote, err := conn.ReadFromUDP(message[:])
	if err != nil {
		handler.logger.Error("Error reading from UDP", zap.Error(err))
		return
	}

	// sendResponse sends a response back to the client
	sendResponse := func(body string) {
		if _, err := conn.WriteToUDP([]byte(body), remote); err != nil {
			handler.logger.Error("Error sending response", zap.Error(err))
		}
	}

	data := strings.TrimSpace(string(message[:rlen]))
	handler.logger.Info("received message", zap.String("remote", remote.String()), zap.String("data", data))

	// Process the received data here
	if _, err := url.ParseRequestURI(data); err != nil {
		handler.logger.Error("Invalid URL", zap.String("URL", data))
		sendResponse("Invalid URL has been provided")
		return
	}

	// Send a response back to the client
	sendResponse("Message received")
}
