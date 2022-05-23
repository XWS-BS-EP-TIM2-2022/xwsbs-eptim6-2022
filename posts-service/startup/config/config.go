package config

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
		GatewayPort:         "8012",            //os.Getenv("GATEWAY_PORT"),
		GrpcPort:            "8011",            //os.Getenv("GRPC_PORT"),
		MongoDbUri:          "localhost:27017", //os.Getenv("MONGODB_URI"),
		Host:                "http://localhost",
		AuthServiceGrpcHost: "localhost",
		AuthServiceGrpcPort: "8051",
	}
}
