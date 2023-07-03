package config

import (
	"github.com/mohammadne/dns-resolver/pkg/logger"
	"github.com/mohammadne/dns-resolver/pkg/tracing"
)

type Config struct {
	Logger  *logger.Config `koanf:"logger"`
}
