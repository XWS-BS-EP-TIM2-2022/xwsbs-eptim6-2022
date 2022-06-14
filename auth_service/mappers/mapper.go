package mappers

import (
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/store"
	authServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
)

func MapPermissions(permissions *store.UserPermission) *authServicePb.UserPermissions {
	var permissionsResponse authServicePb.UserPermissions
	for i := 0; i < len(permissions.Permissions); i++ {
		permissionsResponse.Permissions = append(permissionsResponse.Permissions, &authServicePb.Permission{Value: permissions.Permissions[i]})
	}
	return &permissionsResponse
}
func MapUserToPb(user *store.User) *authServicePb.User {
	userPb := &authServicePb.User{
		Username: user.Username,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
		ApiKey:   user.ApiToken,
	}
	return userPb
}
func MapUsersToPb(users []store.User) (*authServicePb.GetAllResponse, error) {
	var response authServicePb.GetAllResponse
	for i := 0; i < len(users); i++ {
		response.Users = append(response.Users, MapUserToPb(&users[i]))
	}
	return &response, nil
}

func MapPbToUser(user *authServicePb.User) store.User {
	userPb := store.User{
		Username: user.Username,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
		Email:    user.Email,
	}
	return userPb
}

func MapMessToRequest(mess *authServicePb.ChangePasswordRequest) store.ChangePasswordRequest {
	req := store.ChangePasswordRequest{
		Username:    mess.ChangePasswordMessage.Username,
		OldPassword: mess.ChangePasswordMessage.OldPassword,
		NewPassword: mess.ChangePasswordMessage.NewPassword,
	}
	return req
}
