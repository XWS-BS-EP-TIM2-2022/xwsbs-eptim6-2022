package api

import (
	"fmt"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/store"
	pb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
)

func mapUser(user store.User) *pb.User {
	userPb := &pb.User{
		Username: user.Username,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
	}

	return userPb
}

func mapUser2(user *pb.User) store.User {
	fmt.Println("MAP USER" + user.Username)
	userPb := store.User{
		Username: user.Username,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
	}
	return userPb
}

func mapUser3(user *store.User) *pb.User {
	userPb := &pb.User{
		Username: user.Username,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
	}

	return userPb
}
