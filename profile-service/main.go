package main

import (
	"context"
	usersServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/profile_service"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"profile-service/handlers"
)

type Server struct {
	usersServicePb.UnsafeProfileServiceServer
	userHandler *handlers.UserHandler
}

func (s *Server) GetAllUsers(ctx context.Context, in *usersServicePb.EmptyRequest) (*usersServicePb.UsersResponse, error) {
	return nil, nil
}
func (s *Server) AddNewUser(ctx context.Context, in *usersServicePb.UserRequest) (*usersServicePb.UserResponse, error) {
	return nil, nil
}
func (s *Server) AddSkill(ctx context.Context, in *usersServicePb.SkillRequest) (*usersServicePb.UserResponse, error) {
	return nil, nil
}
func (s *Server) AddInterest(ctx context.Context, in *usersServicePb.InterestRequest) (*usersServicePb.UserResponse, error) {
	return nil, nil
}
func (s *Server) AddExperience(ctx context.Context, in *usersServicePb.ExperienceRequest) (*usersServicePb.UserResponse, error) {
	return nil, nil
}
func (s *Server) AddEducation(ctx context.Context, in *usersServicePb.EducationRequest) (*usersServicePb.UserResponse, error) {
	return nil, nil
}
func (s *Server) FollowUser(ctx context.Context, in *usersServicePb.FollowUserRequest) (*usersServicePb.UserResponse, error) {
	return nil, nil
}
func (s *Server) UnFollowUser(ctx context.Context, in *usersServicePb.FollowUserRequest) (*usersServicePb.UserResponse, error) {
	return nil, nil
}
func (s *Server) AcceptFollow(ctx context.Context, in *usersServicePb.FollowUserRequest) (*usersServicePb.UserResponse, error) {
	return nil, nil
}
func (s *Server) RejectFollow(ctx context.Context, in *usersServicePb.FollowUserRequest) (*usersServicePb.UserResponse, error) {
	return nil, nil
}

func RegisterRouts() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)

	uh := handlers.InitUserHandler()

	myRouter.HandleFunc("/user", uh.GetUser).Methods("GET")
	myRouter.HandleFunc("/users", uh.AddNewUser).Methods("POST")
	myRouter.HandleFunc("/users", uh.GetAll).Methods("GET")
	myRouter.HandleFunc("/users", uh.UpdateUser).Methods("PUT")

	myRouter.HandleFunc("/users/experience", uh.AddExperience).Methods("POST")
	myRouter.HandleFunc("/users/education", uh.AddEducation).Methods("POST")
	myRouter.HandleFunc("/users/skill", uh.AddSkill).Methods("POST")
	myRouter.HandleFunc("/users/interest", uh.AddInterest).Methods("POST")

	myRouter.HandleFunc("/users/follow/{id}", uh.FollowUser).Methods("PUT")
	myRouter.HandleFunc("/users/unfollow/{id}", uh.UnfollowUser).Methods("PUT")
	myRouter.HandleFunc("/users/accept-follow-request/{id}", uh.AcceptFollow).Methods("PUT")
	myRouter.HandleFunc("/users/reject-follow-request/{id}", uh.RejectFollow).Methods("PUT")

	return myRouter
}

func main() {
	//router := RegisterRouts()
	//fmt.Println("START Listening")
	//log.Fatal(http.ListenAndServe(":8080", router))
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8003")
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
	usersServicePb.RegisterProfileServiceServer(s, service)
	log.Println("Serving gRPC on 0.0.0.0:8003")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8003",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = usersServicePb.RegisterProfileServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8093",
		Handler: tracingWrapper(gwmux),
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8093")
	log.Fatalln(gwServer.ListenAndServe())
}
func NewServer() (*Server, error) {
	return &Server{userHandler: handlers.InitUserHandler()}, nil
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
