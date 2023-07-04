package config

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
)

const (
	delimiter = "."

	tagName = "koanf"

	upTemplate     = "================ Loaded Configuration ================"
	bottomTemplate = "======================================================"
)

func Load(print bool) *Config {
	k := koanf.New(delimiter)

	// load default configuration from file
	if err := LoadValues(k); err != nil {
		log.Fatalf("error loading default values: \n%v", err)
	}

	config := Config{}
	var tag = koanf.UnmarshalConf{Tag: tagName}
	if err := k.UnmarshalWithConf("", &config, tag); err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}

	if print {
		// pretty print loaded configuration using provided template
		log.Printf("%s\n%v\n%s\n", upTemplate, spew.Sdump(config), bottomTemplate)
	}

	return &config
}

//go:embed default.yml
var values []byte

func LoadValues(k *koanf.Koanf) error {
	if err := k.Load(rawbytes.Provider(values), yaml.Parser()); err != nil {
		return fmt.Errorf("error loading values: %s", err)
	}

	return nil
}
