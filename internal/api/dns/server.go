package dns

import (
	"fmt"
	"log"
	"net"

	"github.com/mohammadne/dns-resolver/internal/models"
	"go.uber.org/zap"
)

type Server struct {
	logger     *zap.Logger
	datasource models.Records
}

func New(log *zap.Logger) *Server {
	server := &Server{logger: log}

	return server
}

func (server *Server) Serve(port int) error {
	address := fmt.Sprintf(":%d", port)

	// Create a UDP address from the string address
	udpAddress, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal(err)
	}

	// Create a UDP connection
	conn, err := net.ListenUDP("udp", udpAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Listening on UDP port", udpAddress.Port)

	for {
		server.handleDNS(conn)
	}
}
