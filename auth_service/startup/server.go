package startup

import (
	"context"
	"fmt"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/handlers"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/mappers"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/startup/config"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/store"
	authServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	"google.golang.org/grpc/metadata"
	"strings"
)

type Server struct {
	authServicePb.UnimplementedAuthServiceServer
	AuthHandler     *handlers.AuthHandler
	PermissionStore *store.PermissionsStore
}

func NewServer(serverConfig *config.Config) (*Server, error) {
	return &Server{AuthHandler: handlers.InitAuthHandler(serverConfig),
		PermissionStore: store.InitPermissionsStore(serverConfig.MongoDbUri)}, nil
}
func (s *Server) AddNewUser(ctx context.Context, in *authServicePb.CreateNewUser) (*authServicePb.CreateNewUser, error) {
	user := mappers.MapPbToUser(in.User)
	user.Role = "USER"
	err := s.AuthHandler.AddNewUser(user)
	if err != nil {
		return &authServicePb.CreateNewUser{}, err
	}
	err = s.AuthHandler.NotifyProfileServiceAboutRegistration(in.User)
	if err != nil {
		return nil, err
	}
	return in, nil
}
func (s *Server) GetAll(ctx context.Context, in *authServicePb.GetAllRequest) (*authServicePb.GetAllResponse, error) {
	return mappers.MapUsersToPb(s.AuthHandler.GetAllUsers())
}
func (s *Server) LoginUser(ctx context.Context, user *authServicePb.CreateNewUser) (*authServicePb.Token, error) {
	jwt, err := s.AuthHandler.LoginUser(mappers.MapPbToUser(user.User))
	return &authServicePb.Token{Token: jwt.Token}, err
}
func (s *Server) AuthorizeJWT(ctx context.Context, token *authServicePb.ValidateToken) (*authServicePb.CreateNewUser, error) {
	user, err := s.AuthHandler.ValidateToken(token.Token.Token)
	if err != nil {
		return nil, err
	}
	return &authServicePb.CreateNewUser{User: mappers.MapUserToPb(user)}, err
}

func (s *Server) GetUserPermissions(ctx context.Context, in *authServicePb.ValidateToken) (*authServicePb.UserPermissions, error) {
	user, err := s.AuthHandler.ValidateToken(in.Token.Token)
	if err != nil {
		return nil, err
	}
	userDb, err := s.AuthHandler.UserStore.FindByUsername(user.Username)
	if err != nil {
		return nil, err
	}
	permission, err := s.PermissionStore.FindByUserRole(userDb.Role)
	if err != nil {
		return nil, err
	}
	return mappers.MapPermissions(permission), nil
}

func (s *Server) UpdateUserPassword(ctx context.Context, in *authServicePb.ChangePasswordRequest) (*authServicePb.CreateNewUser, error) {
	user, err := s.AuthHandler.ValidateToken(getTokenFromContext(ctx))
	changePasswordData := mappers.MapMessToRequest(in)
	changePasswordData.Username = user.Username
	user, err = s.AuthHandler.ChangePassword(changePasswordData)
	return &authServicePb.CreateNewUser{User: mappers.MapUserToPb(user)}, err
}
func (s *Server) ActivateUserAccount(ctx context.Context, in *authServicePb.ActivationToken) (*authServicePb.ActivationResponse, error) {
	err := s.AuthHandler.ActivateAccount(in.Token)
	if err != nil {
		return &authServicePb.ActivationResponse{ResponseStatus: err.Error()}, err
	}

	return &authServicePb.ActivationResponse{ResponseStatus: "Account successfully activated!"}, err
}
func (s *Server) ForgottenPassword(ctx context.Context, in *authServicePb.UserEmailMessage) (*authServicePb.ActivationResponse, error) {
	err := s.AuthHandler.AccountRecoveryEmail(in.Email.Email)
	if err != nil {
		return &authServicePb.ActivationResponse{ResponseStatus: err.Error()}, err
	}

	return &authServicePb.ActivationResponse{ResponseStatus: "Recovery email successfully sent!"}, err
}
func (s *Server) ResetPassword(ctx context.Context, in *authServicePb.ResetPasswordWithTokenMessage) (*authServicePb.ActivationResponse, error) {
	err := s.AuthHandler.ResetPassword(in.Details.Token, in.Details.NewPassword)
	if err != nil {
		return &authServicePb.ActivationResponse{ResponseStatus: err.Error()}, err
	}

	return &authServicePb.ActivationResponse{ResponseStatus: "Password changed successfully!"}, err
}
func (s *Server) GeneratePasswordlessLoginToken(ctx context.Context, in *authServicePb.UserEmailMessage) (*authServicePb.ActivationResponse, error) {
	err := s.AuthHandler.SendPasswordlessLoginEmail(in.Email.Email)
	if err != nil {
		return &authServicePb.ActivationResponse{ResponseStatus: err.Error()}, err
	}

	return &authServicePb.ActivationResponse{ResponseStatus: "Email successfully sent!"}, err
}
func (s *Server) PasswordlessLogin(ctx context.Context, in *authServicePb.ActivationTokenMessage) (*authServicePb.ActivationResponse, error) {
	jwt, err := s.AuthHandler.PasswordlessLogin(in.Token.Token)
	if err != nil {
		return &authServicePb.ActivationResponse{ResponseStatus: err.Error()}, err
	}

	return &authServicePb.ActivationResponse{ResponseStatus: jwt.Token}, err
}

func getTokenFromContext(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Println(md.Get("authorization")[0])
	return strings.Split(md.Get("authorization")[0], " ")[1]
}
