package main

import (
	"context"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/startup"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/startup/config"
	authServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {

	serverConfig := config.NewConfig()
	_, done := createGrpcServer(serverConfig)
	if done {
		return
	}
	creategRPCGateway(serverConfig)
}

func creategRPCGateway(serverConfig *config.Config) {
	conn, err := grpc.DialContext(
		context.Background(),
		serverConfig.Host+":"+serverConfig.GrpcPort,
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
		Addr:    serverConfig.GatewayPort,
		Handler: tracingWrapper(gwmux),
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:" + serverConfig.GatewayPort)
	log.Fatalln(gwServer.ListenAndServe())
}

func createGrpcServer(serverConfig *config.Config) (error, bool) {
	lis, err := net.Listen("tcp", ":"+serverConfig.GrpcPort)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	service, err := startup.NewServer(serverConfig)
	if err != nil {
		log.Fatal(err.Error())
		return nil, true
	}

	authServicePb.RegisterAuthServiceServer(s, service)
	log.Println("Serving gRPC on:" + serverConfig.GrpcPort)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()
	return err, false
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

/*func RegisterRouts() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	rg := handlers.InitAuthHandler()
	router.HandleFunc("/api/auth/users", rg.GetAll).Methods("GET")
	router.HandleFunc("/api/auth/users", rg.AddNewUser).Methods("POST")
	router.HandleFunc("/api/auth/session", rg.LoginUserRequest).Methods("PUT")
	router.HandleFunc("/api/auth/session/validations", rg.AuthorizeJWT).Methods("PUT")
	return router
}
*/
