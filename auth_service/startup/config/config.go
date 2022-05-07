package config

type Config struct {
	Port       string
	AuthDBHost string
	AuthDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:       "8081",      //os.Getenv("AUTH_SERVICE_PORT"),
		AuthDBHost: "localhost", //os.Getenv("AUTH_DB_HOST"),
		AuthDBPort: "27017",     //os.Getenv("AUTH_DB_PORT"),
	}
}
