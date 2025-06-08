package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Service struct {
	Name string `yaml:"name"`
	Cmd  string `yaml:"cmd"`
}

type Config struct {
	Services []Service `yaml:"services"`
}

func Load() *Config {
	data, err := os.ReadFile("internal/config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	// (3) Unmarshal into your struct.
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("error parsing YAML: %v", err)
	}

	return &cfg
}