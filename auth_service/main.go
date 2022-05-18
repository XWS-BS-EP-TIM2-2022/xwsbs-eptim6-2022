package main

import (
	"auth_service/handlers"
	"auth_service/store"
	"context"
	"fmt"
	authServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type Server struct {
	authServicePb.UnimplementedAuthServiceServer
	authHandler *handlers.AuthHandler
}

func mapUser(user *store.User) *authServicePb.User {
	userPb := &authServicePb.User{
		Username: user.Username,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
	}

	return userPb
}
func mapPbToUser(user *authServicePb.User) store.User {
	fmt.Println("MAP USER" + user.Username)
	userPb := store.User{
		Username: user.Username,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
	}
	return userPb
}
func (s *Server) AddNewUser(ctx context.Context, in *authServicePb.CreateNewUser) (*authServicePb.CreateNewUser, error) {
	s.authHandler.UserStore.AddNew(mapPbToUser(in.User))
	return in, nil
}
func (s *Server) GetAll(ctx context.Context, in *authServicePb.GetAllRequest) (*authServicePb.GetAllResponse, error) {
	return nil, nil
}
func (s *Server) LoginUser(ctx context.Context, user *authServicePb.CreateNewUser) (*authServicePb.Token, error) {
	jwt, err := s.authHandler.LoginUser(mapPbToUser(user.User))
	return &authServicePb.Token{Token: jwt.Token}, err
}
func (s *Server) AuthorizeJWT(ctx context.Context, token *authServicePb.ValidateToken) (*authServicePb.CreateNewUser, error) {
	user, err := s.authHandler.ValidateToken(token.Token.Token)
	if err != nil {
		return nil, err
	}
	return &authServicePb.CreateNewUser{User: mapUser(user)}, err
}
func main() {
	//router := RegisterRouts()
	//fmt.Println("START Listening")
	//log.Fatal(http.ListenAndServe(":8080", router))
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	service, err := NewServer()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	authServicePb.RegisterAuthServiceServer(s, service)
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = authServicePb.RegisterAuthServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: tracingWrapper(gwmux),
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
func NewServer() (*Server, error) {
	return &Server{authHandler: handlers.InitAuthHandler()}, nil
}
func RegisterRouts() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	rg := handlers.InitAuthHandler()
	router.HandleFunc("/api/auth/users", rg.GetAll).Methods("GET")
	router.HandleFunc("/api/auth/users", rg.AddNewUser).Methods("POST")
	router.HandleFunc("/api/auth/session", rg.LoginUserRequest).Methods("PUT")
	router.HandleFunc("/api/auth/session/validations", rg.AuthorizeJWT).Methods("PUT")
	return router
}

var grpcGatewayTag = otgo.Tag{Key: string(ext.Component), Value: "grpc-gateway"}

func tracingWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parentSpanContext, err := otgo.GlobalTracer().Extract(
			otgo.HTTPHeaders,
			otgo.HTTPHeadersCarrier(r.Header))
		if err == nil || err == otgo.ErrSpanContextNotFound {
			serverSpan := otgo.GlobalTracer().StartSpan(
				"ServeHTTP",
				// this is magical, it attaches the new span to the parent parentSpanContext, and creates an unparented one if empty.
				ext.RPCServerOption(parentSpanContext),
				grpcGatewayTag,
			)
			r = r.WithContext(otgo.ContextWithSpan(r.Context(), serverSpan))
			defer serverSpan.Finish()
		}
		h.ServeHTTP(w, r)
	})
}
