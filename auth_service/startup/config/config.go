package config

import "os"

type Config struct {
	Host                   string
	GatewayPort            string
	GrpcPort               string
	MongoDbUri             string
	MongoDbName            string
	MongoDbCollection      string
	ProfileServiceGrpcHost string
	ProfileServiceGrpcPort string
	SecretKey              string
	EmailPort              string
	EmailHost              string
	EmailFrom              string
	EmailPassword          string
	FrontendUri            string
}

func NewConfig() *Config {
	return &Config{
		GatewayPort:       "8001",               //os.Getenv("GATEWAY_PORT"),
		GrpcPort:          "8051",               //os.Getenv("GRPC_PORT"),
		MongoDbUri:        "localhost:27017",    //os.Getenv("MONGODB_URI"),
		MongoDbName:       "auth_service_users", //os.Getenv("MONGODB_URI"),
		MongoDbCollection: "users",              //os.Getenv("MONGODB_URI"),

		Host:                   "localhost",
		ProfileServiceGrpcHost: "localhost",
		ProfileServiceGrpcPort: "8007",
		SecretKey:              "secret-key",
		EmailPort:              "587",
		EmailHost:              "smtp.gmail.com",
		EmailFrom:              os.Getenv("DISLINKT_MAIL"),          //bsepdislinkt@gmail.com
		EmailPassword:          os.Getenv("DISLINKT_MAIL_PASSWORD"), //Dislinkt123
		FrontendUri:            "localhost:4200",
	}
}
