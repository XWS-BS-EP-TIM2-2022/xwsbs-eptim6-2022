package config

import "os"

type Config struct {
	Port     string
	AuthHost string
	AuthPort string
}

func NewConfig() *Config {
	return &Config{
		Port:     os.Getenv("GATEWAY_PORT"),
		AuthHost: os.Getenv("AUTH_SERVICE_HOST"),
		AuthPort: os.Getenv("AUTH_SERVICE_PORT"),
	}
}
