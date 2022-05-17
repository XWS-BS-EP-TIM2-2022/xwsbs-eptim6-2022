package startup

import (
	"fmt"
	prof "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/profile_service"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/profile-service/application"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/profile-service/infrastructure/api"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/profile-service/startup/config"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/profile-service/store"
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
	profileStore := server.initProfileStore(mongoClient)

	profileService := server.initProfileService(profileStore)

	profileHandler := server.initProfileHandler(profileService)

	server.startGrpcServer(profileHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := store.GetClient(server.config.ProfileDBHost, server.config.ProfileDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initProfileStore(client *mongo.Client) *store.UsersStore {
	store := store.InitUsersStore(client)
	return store
}

func (server *Server) initProfileService(store store.ProfileStore) *application.ProfileService {
	return application.NewProfileService(store)
}

func (server *Server) initProfileHandler(service *application.ProfileService) *api.AuthHandler {
	return api.NewAuthHandler(service)
}

func (server *Server) startGrpcServer(authHandler *api.AuthHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	prof.RegisterProfileServiceServer(grpcServer, authHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
