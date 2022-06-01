package config

type Config struct {
	Host                  string
	GatewayPort           string
	GrpcPort              string
	MongoDbUri            string
	MongoDbCollectionName string
	MongoDbName           string
	AuthServiceGrpcHost   string
	AuthServiceGrpcPort   string
}

func NewConfig() *Config {
	return &Config{
		GatewayPort:           "8020",
		GrpcPort:              "5020",
		MongoDbUri:            "localhost:27017",
		MongoDbName:           "job-offers",
		MongoDbCollectionName: "job-offers",
		Host:                  "localhost",
		AuthServiceGrpcHost:   "localhost",
		AuthServiceGrpcPort:   "5001",
	}
	//return &Config{
	//	GatewayPort:         os.Getenv("GATEWAY_POSTS_PORT"),
	//	GrpcPort:            os.Getenv("GRPC_POSTS_PORT"),
	//	MongoDbUri:          os.Getenv("MONGO_DB_URI_POSTS"),
	//	Host:                os.Getenv("PROFILE_SERVICE_HOST"),
	//	AuthServiceGrpcHost: os.Getenv("AUTH_SERVICE_GRPC_HOST"),
	//	AuthServiceGrpcPort: os.Getenv("AUTH_SERVICE_GRPC_PORT"),
	//}
}
