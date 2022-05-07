package config

type Config struct {
	Port     string
	AuthHost string
	AuthPort string
}

func NewConfig() *Config {
	return &Config{
		Port:     "8082",      //os.Getenv("GATEWAY_PORT"),
		AuthHost: "localhost", //os.Getenv("AUTH_SERVICE_HOST"),
		AuthPort: "8081",      //os.Getenv("AUTH_SERVICE_PORT"),
	}
}
