package dns

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func (handler *Server) handleDNS(conn *net.UDPConn) {
	message := make([]byte, 1024)
	rlen, remote, err := conn.ReadFromUDP(message[:])
	if err != nil {
		log.Println("Error reading from UDP:", err)
		panic(err)
	}

	data := strings.TrimSpace(string(message[:rlen]))
	fmt.Printf("received: %s from %s\n", data, remote)

	// Process the received data here
	// ...

	// Send a response back to the client
	_, err = conn.WriteToUDP([]byte("Message received"), remote)
	if err != nil {
		log.Println("Error sending response:", err)
	}
}
