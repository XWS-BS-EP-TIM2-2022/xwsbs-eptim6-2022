package startup

import (
	"context"
	"fmt"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/api_gateway/infrastructure/service"
	cfg "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/api_gateway/startup/config"
	authGw "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	profileGw "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/profile_service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type Server struct {
	config      *cfg.Config
	mux         *runtime.ServeMux
	authService *service.AuthService
}

type SecurityServer struct {
	mainServer *Server
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config:      config,
		mux:         runtime.NewServeMux(),
		authService: service.InitAuthService(config),
	}
	server.initHandlers()
	server.initCustomHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	authEndpoint := fmt.Sprintf("%s:%s", server.config.AuthHost, server.config.AuthPort)
	err := authGw.RegisterAuthServiceHandlerFromEndpoint(context.TODO(), server.mux, authEndpoint, opts)
	if err != nil {
		panic(err)
	}
	profileEndpoint := fmt.Sprintf("%s:%s", server.config.ProfileHost, server.config.ProfilePort)
	err = profileGw.RegisterProfileServiceHandlerFromEndpoint(context.TODO(), server.mux, profileEndpoint, opts)
	if err != nil {
		panic(err)
	}
	/*shippingEmdpoint := fmt.Sprintf("%s:%s", server.config.ShippingHost, server.config.ShippingPort)
	err = shippingGw.RegisterShippingServiceHandlerFromEndpoint(context.TODO(), server.mux, shippingEmdpoint, opts)
	if err != nil {
		panic(err)
	}
	inventoryEmdpoint := fmt.Sprintf("%s:%s", server.config.InventoryHost, server.config.InventoryPort)
	err = inventoryGw.RegisterInventoryServiceHandlerFromEndpoint(context.TODO(), server.mux, inventoryEmdpoint, opts)
	if err != nil {
		panic(err)
	}*/
}

func (server *Server) initCustomHandlers() {
	/*authEndpoint := fmt.Sprintf("%s:%s", server.config.AuthHost, server.config.AuthPort)
	orderingEmdpoint := fmt.Sprintf("%s:%s", server.config.OrderingHost, server.config.OrderingPort)
	shippingEmdpoint := fmt.Sprintf("%s:%s", server.config.ShippingHost, server.config.ShippingPort)
	orderingHandler := api.NewOrderingHandler(orderingEmdpoint, catalogueEmdpoint, shippingEmdpoint)
	orderingHandler.Init(server.mux)*/
}

func authWrapper(h http.Handler, s *Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request := fmt.Sprintf("%s%s", r.Method, r.RequestURI)
		token := s.authService.GetAuthToken(r)
		if token == "" {
			if s.config.SecurityPermissions.ValidateUnauthorizedRequest(request) {
				h.ServeHTTP(w, r)
				return
			} else {
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, "Not Authorized")
			}
		}
		userPermissions, _ := s.authService.GetUserPermissions(token)
		hasPermissions := s.config.SecurityPermissions.ValidatePermission(userPermissions, request)
		if hasPermissions {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func (server *Server) Start() {
	gwServer := &http.Server{
		Addr:    ":" + server.config.Port,
		Handler: authWrapper(server.mux, server),
	}
	s := fmt.Sprintf("Serving gRPC-Gateway on http://0.0.0.0:%s", server.config.Port)
	fmt.Println(s)
	log.Fatalln(gwServer.ListenAndServe())

	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), )
}
