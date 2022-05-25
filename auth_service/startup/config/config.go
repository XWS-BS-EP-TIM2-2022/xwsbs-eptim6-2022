package config

import "os"

type Config struct {
	Host                   string
	GatewayPort            string
	GrpcPort               string
	MongoDbUri             string
	ProfileServiceGrpcHost string
	ProfileServiceGrpcPort string
	SecretKey              string
}

func NewConfig() *Config {
	return &Config{
		GatewayPort:            os.Getenv("GATEWAY_AUTH_PORT"),
		GrpcPort:               os.Getenv("GRPC_AUTH_PORT"),
		MongoDbUri:             os.Getenv("MONGO_DB_URI_AUTH"),
		Host:                   os.Getenv("AUTH_HOST"),
		ProfileServiceGrpcHost: os.Getenv("PROFILE_SERVICE_GRPC_HOST"),
		ProfileServiceGrpcPort: os.Getenv("PROFILE_SERVICE_GRPC_PORT"),
		SecretKey:              os.Getenv("SECRET_KEY_AUTH"),
	}
}
