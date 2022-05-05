package config

import "os"

type Config struct {
	Port       string
	AuthDBHost string
	AuthDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:       os.Getenv("CATALOGUE_SERVICE_PORT"),
		AuthDBHost: os.Getenv("CATALOGUE_DB_HOST"),
		AuthDBPort: os.Getenv("CATALOGUE_DB_PORT"),
	}
}
