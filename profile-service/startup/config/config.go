package config

type Config struct {
	Host                string
	GatewayPort         string
	GrpcPort            string
	MongoDbUri          string
	MongoDbName         string
	MongoDbCollection   string
	AuthServiceGrpcHost string
	AuthServiceGrpcPort string
}

func NewConfig() *Config {
	return &Config{
		GatewayPort:         "8004",
		GrpcPort:            "5003",
		MongoDbUri:          "localhost:27017",
		Host:                "localhost",
		AuthServiceGrpcHost: "localhost",
		AuthServiceGrpcPort: "5001",
		MongoDbName:         "users_profiles", //os.Getenv("MONGODB_URI"),
		MongoDbCollection:   "profiles",       //os.Getenv("MONGODB_URI"),
	}
	//return &Config{
	//	GatewayPort:         os.Getenv("GATEWAY_PROFILE_PORT"),
	//	GrpcPort:            os.Getenv("GRPC_PROFILE_PORT"),
	//	MongoDbUri:          os.Getenv("MONGO_DB_URI_PROFILE"),
	//	Host:                os.Getenv("PROFILE_SERVICE_HOST"),
	//	AuthServiceGrpcHost: os.Getenv("AUTH_SERVICE_GRPC_HOST"),
	//	AuthServiceGrpcPort: os.Getenv("AUTH_SERVICE_GRPC_PORT"),
	//	MongoDbName:         "users_profiles",  //os.Getenv("MONGODB_URI"),
	//	MongoDbCollection:   "profiles",        //os.Getenv("MONGODB_URI"),
	//}
}
