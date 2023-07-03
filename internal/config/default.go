package config

import (
	"github.com/mohammadne/dns-resolver/pkg/logger"
	"github.com/mohammadne/dns-resolver/pkg/tracing"
)

func Default() *Config {
	return &Config{
		Logger: &logger.Config{
			Development: true,
			Level:       "debug",
			Encoding:    "console",
		},
	}
}
