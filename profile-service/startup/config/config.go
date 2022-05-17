package config

type Config struct {
	Port          string
	ProfileDBHost string
	ProfileDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:          "8081",      //os.Getenv("AUTH_SERVICE_PORT"),
		ProfileDBHost: "localhost", //os.Getenv("AUTH_DB_HOST"),
		ProfileDBPort: "27017",     //os.Getenv("AUTH_DB_PORT"),
	}
}
