package infrastructure

import (
	"fmt"
	authGw "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"profile-service/startup/config"
)

type AuthService struct {
	client authGw.AuthServiceClient
}

func InitAuthServiceClient(config *config.Config) authGw.AuthServiceClient {
	authEndpoint := fmt.Sprintf("%s:%s", config.AuthServiceGrpcHost, config.AuthServiceGrpcPort)
	conn, err := getConnection(authEndpoint)
	if err != nil {
		fmt.Println("Fatal error init auth service connection!")
		return nil
	}
	authConn := authGw.NewAuthServiceClient(conn)
	return authConn
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
