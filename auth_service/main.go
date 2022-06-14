package main

import (
	"context"
	"fmt"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/startup"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/startup/config"
	logger "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/logger"
	authServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jasonlvhit/gocron"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
)

var log *logger.LoggerWrapper

func init() {
	s := gocron.NewScheduler()
	log = logger.InitLogger("auth_service", s)
}
func main() {
	serverConfig := config.NewConfig()
	err := createGrpcServer(serverConfig)
	if err != nil {
		log.Writeln(logger.LogMessage{Message: err.Error(), Level: logrus.FatalLevel, Component: "auth_service.main.createGrpcServer"})
		return
	}
	err = creategRPCGateway(serverConfig)
	if err != nil {
		log.Writeln(logger.LogMessage{Message: err.Error(), Level: logrus.FatalLevel, Component: "auth_service.main.creategRPCGateway"})
		return
	}
	<-log.Scheduler.Start()
}

func creategRPCGateway(serverConfig *config.Config) error {
	conn, err := grpc.DialContext(
		context.Background(),
		serverConfig.Host+":"+serverConfig.GrpcPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Writeln(logger.LogMessage{Message: fmt.Sprintf("Failed to dial server: %s", err), Level: logrus.FatalLevel, Component: "auth_service.main.creategRPCGateway"})
		return err
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = authServicePb.RegisterAuthServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Writeln(logger.LogMessage{Message: fmt.Sprintf("Failed to register gateway: %s", err), Level: logrus.FatalLevel, Component: "auth_service.main.creategRPCGateway"})
		return err
	}

	gwServer := &http.Server{
		Addr:    ":" + serverConfig.GatewayPort,
		Handler: tracingWrapper(gwmux),
	}
	log.Writeln(logger.LogMessage{Message: fmt.Sprintf("Serving gRPC-Gateway on http://0.0.0.0:%s", serverConfig.GatewayPort), Level: logrus.InfoLevel, Component: startup.GetComponentName(creategRPCGateway)})
	err = gwServer.ListenAndServe()
	if err != nil {
		log.Writeln(logger.LogMessage{Message: fmt.Sprintf("Serving gRPC-Gateway on http://0.0.0.0:%s", serverConfig.GatewayPort), Level: logrus.FatalLevel, Component: startup.GetComponentName(creategRPCGateway)})
		return err
	}
	return nil
}

func createGrpcServer(serverConfig *config.Config) error {
	lis, err := net.Listen("tcp", ":"+serverConfig.GrpcPort)
	if err != nil {
		return err
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	service, err := startup.NewServer(serverConfig, log)
	if err != nil {
		return err
	}

	authServicePb.RegisterAuthServiceServer(s, service)
	log.Writeln(logger.LogMessage{Message: fmt.Sprintf("Serving gRPC on: %s", serverConfig.GrpcPort), Level: logrus.InfoLevel, Component: startup.GetComponentName(createGrpcServer)})
	go func() {
		err := s.Serve(lis)
		if err != nil {
			log.Writeln(logger.LogMessage{Message: err.Error(), Level: logrus.FatalLevel, Component: startup.GetComponentName(createGrpcServer)})
			return
		}
	}()
	return err
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
