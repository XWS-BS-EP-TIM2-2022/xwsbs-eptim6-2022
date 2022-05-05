package startup

import (
	"auth_service/application"
	"auth_service/infrastructure/api"
	"auth_service/store"
	"fmt"
	//"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/api_gateway/startup/config"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/startup/config"
	auths "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	productStore := server.initAuthStore(mongoClient)

	productService := server.initAuthService(productStore)

	productHandler := server.initAuthHandler(productService)

	server.startGrpcServer(productHandler)
}

//Da li je ovo initUserStore iz users_store.go ???????
func (server *Server) initMongoClient() *mongo.Client {
	client, err := store.GetClient()
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUsersStore(client *mongo.Client) store.AuthStore {
	store := store.InitUsersStore()
	return store
}

func (server *Server) initAuthService(store store.AuthStore) *application.AuthService {
	return application.NewAuthService(store)
}

func (server *Server) initAuthHandler(service *application.AuthService) *api.AuthHandlerHandler {
	return api.NewAuthHandler(service)
}

func (server *Server) startGrpcServer(authHandler *api.AuthHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	auths.RegisterAuthServiceServer(grpcServer, authHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
