package config

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var Cfg *Config

type Service struct {
	Name string `yaml:"name"`
	Cmd  string `yaml:"cmd"`
}

type Auth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	Services []Service `yaml:"services"`
	Auth     Auth      `yaml:"auth"`
	Port     int       `yaml:"port"`
	Slack    struct {
		WebhookURL string `yaml:"webhook-url"`
	}
}

func Load() *Config {

	cfgPath := flag.String("config", "config.yaml", "Path to config file")

	flag.Parse()

	path := *cfgPath

	data, err := os.ReadFile(path)
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

func init() {
	Cfg = Load()
}
