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
		GatewayPort:         os.Getenv("GATEWAY_POSTS_PORT"),
		GrpcPort:            os.Getenv("GRPC_POSTS_PORT"),
		MongoDbUri:          os.Getenv("MONGO_DB_URI_POSTS"),
		Host:                os.Getenv("PROFILE_SERVICE_HOST"),
		AuthServiceGrpcHost: os.Getenv("AUTH_SERVICE_GRPC_HOST"),
		AuthServiceGrpcPort: os.Getenv("AUTH_SERVICE_GRPC_PORT"),
	}
}
