package dns

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"go.uber.org/zap"
)

type Server struct {
	logger  *zap.Logger
	records map[string]string
	metrics *metrics
}

//go:embed mappings.csv
var mappings []byte

func New(log *zap.Logger) *Server {
	server := &Server{logger: log, metrics: newMetrics()}

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
		buffer := make([]byte, 1500)
		rlen, remote, err := conn.ReadFromUDP(buffer[:])
		if err != nil {
			server.logger.Error("Error reading from UDP", zap.Error(err))
			continue
		}

		go func() {
			start := time.Now()

			server.metrics.requestsInProgress.WithLabelValues().Inc()
			server.metrics.requestsTotal.WithLabelValues().Inc()

			response, err := server.handleDNS(buffer[:rlen])
			if err != nil {
				server.logger.Error("Error processing the request", zap.Error(err))
				return
			}

			// Send a response back to the client
			if _, err := conn.WriteToUDP(response, remote); err != nil {
				server.logger.Error("Error sending response", zap.Error(err))
			}

			elapsed := float64(time.Since(start).Nanoseconds()) / 1e9
			server.metrics.requestsDuration.WithLabelValues().Observe(elapsed)
			server.metrics.requestsInProgress.WithLabelValues().Dec()
		}()
	}
}
