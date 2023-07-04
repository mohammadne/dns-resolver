package dns

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"log"
	"net"
	"strings"

	"go.uber.org/zap"
)

type Server struct {
	logger  *zap.Logger
	records map[string]string
}

//go:embed mappings.csv
var mappings []byte

func New(log *zap.Logger) *Server {
	server := &Server{logger: log}

	{
		reader := csv.NewReader(strings.NewReader(string(mappings)))
		records, err := reader.ReadAll()
		if err != nil {
			log.Fatal("Error reading csv mapping file", zap.Error(err))
		}

		server.records = make(map[string]string, len(records))
		for _, record := range records {
			server.records[record[0]] = record[1]
		}
	}

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
