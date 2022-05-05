package api

import (
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/store"
	pb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	service *store.UsersStore
}

func NewAuthHandler(service *store.UsersStore) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}
