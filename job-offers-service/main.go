package main

import (
	"context"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/logger"
	jobOffersPb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/job_offers_service"
	"github.com/jasonlvhit/gocron"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"job-offers-service/startup"
	"job-offers-service/startup/config"
	"log"
	"net"
	"net/http"
)

var logs *logger.LoggerWrapper

func init() {
	s := gocron.NewScheduler()
	logs = logger.InitLogger("job-offers-service", s)
}

func main() {
	serverConfig := config.NewConfig()
	lis, err := net.Listen("tcp", ":"+serverConfig.GrpcPort)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	service, err := startup.NewServer(*serverConfig, logs)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	jobOffersPb.RegisterJobOffersServiceServer(s, service)
	log.Println("Serving gRPC on 0.0.0.0:" + serverConfig.GrpcPort)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0"+":"+serverConfig.GrpcPort,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = jobOffersPb.RegisterJobOffersServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":" + serverConfig.GatewayPort,
		Handler: tracingWrapper(gwmux),
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0" + ":" + serverConfig.GatewayPort)
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
