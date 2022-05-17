package api

import "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/profile-service/application"

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	service *application.AuthService
}
