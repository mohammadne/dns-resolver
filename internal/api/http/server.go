package http

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	logger *zap.Logger
	app    *fiber.App
}

func New(log *zap.Logger) *Server {
	server := &Server{logger: log}

	server.app = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	server.app.Get("/healthz/liveness", server.liveness)
	server.app.Get("/healthz/readiness", server.readiness)

	return server
}

func (server *Server) Serve(port int) error {
	address := fmt.Sprintf(":%d", port)
	if err := server.app.Listen(address); err != nil {
		server.logger.Error("error resolving server", zap.Error(err))
		return err
	}
	return nil
}
