package main

import (
	"context"
	"fmt"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/logger"
	postsServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/posts_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jasonlvhit/gocron"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"xwsbs-eptim6-2022/posts-service/startup"
	"xwsbs-eptim6-2022/posts-service/startup/config"
)

var logging *logger.LoggerWrapper

func init() {
	s := gocron.NewScheduler()
	logging = logger.InitLogger("posts-service", s)
}

func main() {
	serverConfig := config.NewConfig()
	lis, err := net.Listen("tcp", ":"+serverConfig.GrpcPort)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	service, err := startup.NewServer(serverConfig, logging)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	postsServicePb.RegisterPostsServiceServer(s, service)
	log.Println("Serving gRPC on 0.0.0.0:" + serverConfig.GrpcPort)
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
		logging.Writeln(logger.LogMessage{Message: fmt.Sprintf("Failed to dial server: %s", err), Level: logrus.FatalLevel, Component: "posts-service.main"})
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = postsServicePb.RegisterPostsServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		logging.Writeln(logger.LogMessage{Message: fmt.Sprintf("Failed to register gateway: %s", err), Level: logrus.FatalLevel, Component: "posts-service.main"})
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: tracingWrapper(gwmux),
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
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
