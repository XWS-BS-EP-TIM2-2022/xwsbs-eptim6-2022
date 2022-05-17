package api

import (
	"context"
	pb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/profile_service"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/profile-service/application"
)

type AuthHandler struct {
	pb.UnimplementedProfileServiceServer
	service *application.ProfileService
}

func NewAuthHandler(service *application.ProfileService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (handler *AuthHandler) AddNew(ctx context.Context, request *pb.CreateNewUser) (*pb.CreateNewUser, error) {

}
