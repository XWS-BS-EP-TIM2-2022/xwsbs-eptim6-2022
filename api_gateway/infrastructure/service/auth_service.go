package service

import (
	"context"
	"fmt"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/api_gateway/startup/config"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/logger"
	authGw "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strings"
)

type AuthService struct {
	client authGw.AuthServiceClient
	log    *logger.LoggerWrapper
}

func InitAuthService(config *config.Config, wrapper *logger.LoggerWrapper) *AuthService {
	authEndpoint := fmt.Sprintf("%s:%s", config.AuthHost, config.AuthPort)
	conn, err := getConnection(authEndpoint)
	if err != nil {
		fmt.Println("Fatal error init auth service connection!")
		return nil
	}
	authConn := authGw.NewAuthServiceClient(conn)
	return &AuthService{client: authConn, log: wrapper}
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
func (as *AuthService) GetAuthToken(r *http.Request) string {
	if r.Header["Authorization"] != nil {
		token := strings.Split(r.Header["Authorization"][0], " ")[1]
		fmt.Println(token)
		return token
	}
	return ""
}

func (as *AuthService) GetUserPermissions(token string) ([]string, error) {
	permissions, err := as.client.GetUserPermissions(context.TODO(), &authGw.ValidateToken{Token: &authGw.Token{Token: token}})
	if err != nil {
		return nil, err
	}
	return mapPermisionToStrings(permissions.Permissions), nil
}

func mapPermisionToStrings(permissions []*authGw.Permission) []string {
	var returnValues []string
	for i := 0; i < len(permissions); i++ {
		returnValues = append(returnValues, permissions[i].Value)
	}
	return returnValues
}
