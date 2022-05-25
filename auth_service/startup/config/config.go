package config

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
		GatewayPort:            "8003",
		GrpcPort:               "5001",
		MongoDbUri:             "localhost:27017",
		Host:                   "localhost",
		ProfileServiceGrpcHost: "localhost",
		ProfileServiceGrpcPort: "5003",
		SecretKey:              "security-key",
		MongoDbName:            "auth_service_users", //os.Getenv("MONGODB_URI"),
		MongoDbCollection:      "users",              //os.Getenv("MONGODB_URI"),
		EmailPort:              "587",
		EmailHost:              "smtp.gmail.com",
		EmailFrom:              "bsepdislinkt@gmail.com", //bsepdislinkt@gmail.com
		EmailPassword:          "Dislinkt123",            //Dislinkt123
		FrontendUri:            "localhost:4200",
	}
	//return &Config{
	//	GatewayPort:            os.Getenv("GATEWAY_AUTH_PORT"),
	//	GrpcPort:               os.Getenv("GRPC_AUTH_PORT"),
	//	MongoDbUri:             os.Getenv("MONGO_DB_URI_AUTH"),
	//	Host:                   os.Getenv("AUTH_HOST"),
	//	ProfileServiceGrpcHost: os.Getenv("PROFILE_SERVICE_GRPC_HOST"),
	//	ProfileServiceGrpcPort: os.Getenv("PROFILE_SERVICE_GRPC_PORT"),
	//	SecretKey:              os.Getenv("SECRET_KEY_AUTH"),
	//	MongoDbName:       "auth_service_users", //os.Getenv("MONGODB_URI"),
	//	MongoDbCollection: "users",              //os.Getenv("MONGODB_URI"),
	//	EmailPort:              "587",
	//	EmailHost:              "smtp.gmail.com",
	//	EmailFrom:              os.Getenv("DISLINKT_MAIL"),          //bsepdislinkt@gmail.com
	//	EmailPassword:          os.Getenv("DISLINKT_MAIL_PASSWORD"), //Dislinkt123
	//	FrontendUri:            "localhost:4200",
	//}
}
