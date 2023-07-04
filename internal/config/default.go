package config

import (
	"github.com/mohammadne/dns-resolver/pkg/logger"
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
