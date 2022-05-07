package api

import (
	"context"
	"fmt"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/application"
	pb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	service *application.AuthService
}

func NewAuthHandler(service *application.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (handler *AuthHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	fmt.Println("GetAll")
	users := handler.service.FindAll()

	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}

	for _, user := range users {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler *AuthHandler) AddNewUser(ctx context.Context, request *pb.CreateNewUser) (*pb.User, error) {
	user := mapUser2(request.User)
	handler.service.AddNew(user)
	return mapUser(user), nil
}

func (handler *AuthHandler) LoginUser(ctx context.Context, request *pb.CreateNewUser) (*pb.Token, error) {
	fmt.Println("GRPC LOGIN")
	if request.User == nil {
		fmt.Println("request je nil")
	}
	user := mapUser2(request.User)
	fmt.Println("USER")
	token := handler.service.LoginUser(user)
	fmt.Println("TOKEN")
	response := &pb.Token{
		Token: token,
	}
	fmt.Println("RETURN")
	return response, nil
}
