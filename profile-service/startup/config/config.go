package config

import "os"

type Config struct {
	Host                string
	GatewayPort         string
	GrpcPort            string
	MongoDbUri          string
	AuthServiceGrpcHost string
	AuthServiceGrpcPort string
}

func NewConfig() *Config {
	return &Config{
		GatewayPort:         os.Getenv("GATEWAY_PROFILE_PORT"),
		GrpcPort:            os.Getenv("GRPC_PROFILE_PORT"),
		MongoDbUri:          os.Getenv("MONGO_DB_URI_PROFILE"),
		Host:                os.Getenv("PROFILE_SERVICE_HOST"),
		AuthServiceGrpcHost: os.Getenv("AUTH_SERVICE_GRPC_HOST"),
		AuthServiceGrpcPort: os.Getenv("AUTH_SERVICE_GRPC_PORT"),
	}
}
