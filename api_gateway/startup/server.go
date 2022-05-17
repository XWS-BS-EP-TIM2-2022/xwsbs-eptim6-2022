package startup

import (
	"context"
	"fmt"
	cfg "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/api_gateway/startup/config"
	authGw "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	profGw "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/profile_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	authEndPoint := fmt.Sprintf("%s:%s", server.config.AuthHost, server.config.AuthPort)
	fmt.Println("Api gateway")
	err := authGw.RegisterAuthServiceHandlerFromEndpoint(context.TODO(), server.mux, authEndPoint, opts)
	if err != nil {
		panic(err)
	}
	profileEndPoint := fmt.Sprintf("%s:%s", server.config.ProfileHost, server.config.ProfilePort)
	err = profGw.RegisterProfileServiceHandlerFromEndpoint(context.TODO(), server.mux, profileEndPoint, opts)
	if err != nil {
		panic(err)
	}
	/*Ovo se radi za sve ostale servise*/
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}
