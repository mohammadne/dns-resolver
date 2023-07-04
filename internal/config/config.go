package config

import (
	"github.com/mohammadne/dns-resolver/pkg/logger"
)

type Config struct {
	Logger *logger.Config `koanf:"logger"`
}
