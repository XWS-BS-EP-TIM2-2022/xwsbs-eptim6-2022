package api

import (
	pb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/profile_service"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/profile-service/store"
)

func mapUser(user *pb.User) store.User {
	userPb := store.User{
		Username: user.Username,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
		Email: user.Email,
		Telephone: user.Telephone,
		Gender: user.Gender,
		BirthDate: user.BirthDate,
		Biography: user.Biography,
		Experiences: user.Expiriences,
		Educations: user.Educations,
		Skills: user.Skills,
		Interests: user.Inte
	}
	return userPb
}
