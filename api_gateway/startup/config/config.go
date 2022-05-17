package config

type Config struct {
	Port        string
	AuthHost    string
	AuthPort    string
	ProfileHost string
	ProfilePort string
}

func NewConfig() *Config {
	return &Config{
		Port:        "8082",      //os.Getenv("GATEWAY_PORT"),
		AuthHost:    "localhost", //os.Getenv("AUTH_SERVICE_HOST"),
		AuthPort:    "8081",      //os.Getenv("AUTH_SERVICE_PORT"),
		ProfileHost: "localhost", //os.Getenv("AUTH_SERVICE_HOST"),
		ProfilePort: "8081",
	}
}
