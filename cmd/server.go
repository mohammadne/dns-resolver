package cmd

import (
	"os"

	"github.com/mohammadne/dns-resolver/internal/api/dns"
	"github.com/mohammadne/dns-resolver/internal/api/http"
	"github.com/mohammadne/dns-resolver/internal/config"
	"github.com/mohammadne/dns-resolver/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type Server struct{}

func (cmd Server) Command(trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		cmd.main(config.Load(true), trap)
	}

	return &cobra.Command{
		Use:   "server",
		Short: "run dns-resolver server",
		Run:   run,
	}
}

func (cmd *Server) main(cfg *config.Config, trap chan os.Signal) {
	logger := logger.NewZap(cfg.Logger)

	go dns.New(logger).Serve(5354)
	go http.New(logger).Serve(8080)

	// Keep this at the bottom of the main function
	field := zap.String("signal trap", (<-trap).String())
	logger.Info("exiting by receiving a unix signal", field)
}
