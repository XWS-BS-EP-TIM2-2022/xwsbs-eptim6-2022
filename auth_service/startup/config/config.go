package config

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
		GatewayPort:            "8001",            //os.Getenv("GATEWAY_PORT"),
		GrpcPort:               "8051",            //os.Getenv("GRPC_PORT"),
		MongoDbUri:             "localhost:27017", //os.Getenv("MONGODB_URI"),
		Host:                   "localhost",
		ProfileServiceGrpcHost: "localhost",
		ProfileServiceGrpcPort: "8003",
		SecretKey:              "secret-key",
	}
}
