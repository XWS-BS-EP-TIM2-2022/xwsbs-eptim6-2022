package startup

import (
	"context"
	"fmt"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/handlers"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/mappers"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/startup/config"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/store"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/consts"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/logger"
	authServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"reflect"
	"runtime"
	"strings"
)

type Server struct {
	authServicePb.UnimplementedAuthServiceServer
	AuthHandler     *handlers.AuthHandler
	PermissionStore *store.PermissionsStore
	log             *logger.LoggerWrapper
}

func NewServer(serverConfig *config.Config, log *logger.LoggerWrapper) (*Server, error) {
	authHandler, err := handlers.InitAuthHandler(serverConfig, log)
	if err != nil {
		return nil, err
	}
	return &Server{AuthHandler: authHandler,
		PermissionStore: store.InitPermissionsStore(serverConfig.MongoDbUri), log: log}, nil
}
func (s *Server) AddNewUser(ctx context.Context, in *authServicePb.CreateNewUser) (*authServicePb.CreateNewUser, error) {
	user := mappers.MapPbToUser(in.User)
	user.Role = "USER"
	err := s.AuthHandler.AddNewUser(&user)
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: err.Error(), Level: logrus.FatalLevel, Component: GetComponentName(s.AuthHandler.AddNewUser)})
		return &authServicePb.CreateNewUser{}, err
	}
	err = s.AuthHandler.NotifyProfileServiceAboutRegistration(in.User, &user)
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: err.Error(), Level: logrus.FatalLevel, Component: GetComponentName(s.AuthHandler.NotifyProfileServiceAboutRegistration)})
		return nil, err
	}
	s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User created. Username %s, Ip address: %s", user.Username, getRequestIpAddressFromContext(ctx)),
		Level: logrus.InfoLevel, Component: GetComponentName(s.AddNewUser)})
	return &authServicePb.CreateNewUser{User: mappers.MapUserToPb(&user)}, nil
}
func (s *Server) GetAll(ctx context.Context, in *authServicePb.GetAllRequest) (*authServicePb.GetAllResponse, error) {
	getRequestIpAddressFromContext(ctx)
	return mappers.MapUsersToPb(s.AuthHandler.GetAllUsers())
}
func (s *Server) LoginUser(ctx context.Context, user *authServicePb.CreateNewUser) (*authServicePb.Token, error) {
	jwt, err := s.AuthHandler.LoginUser(mappers.MapPbToUser(user.User))
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User authentication failed. %s. Ip address: %s ", err.Error(), getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
			Component: GetComponentName(s.AuthHandler.LoginUser)})
		return &authServicePb.Token{}, err
	}
	s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User %s logged in from %s", user.User.Username, getRequestIpAddressFromContext(ctx)),
		Level: logrus.InfoLevel, Component: GetComponentName(s.LoginUser)})
	return &authServicePb.Token{Token: jwt.Token}, err
}
func (s *Server) AuthorizeJWT(ctx context.Context, token *authServicePb.ValidateToken) (*authServicePb.CreateNewUser, error) {
	user, err := s.AuthHandler.ValidateToken(token.Token.Token)
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("%s", err.Error()), Level: logrus.WarnLevel, Component: GetComponentName(s.AuthHandler.ValidateToken)})
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
	if in.Token.Token == userDb.ApiToken {
		ownerPermission, err := s.PermissionStore.FindByUserRole(string(consts.COMPANY_OWNER))
		if err != nil {
			return nil, err
		}
		return mappers.MapPermissions(ownerPermission), nil
	}
	return mappers.MapPermissions(permission), nil
}

func (s *Server) UpdateUserPassword(ctx context.Context, in *authServicePb.ChangePasswordRequest) (*authServicePb.CreateNewUser, error) {
	user, err := s.AuthHandler.ValidateToken(s.getTokenFromContext(ctx))
	if err != nil {
		return &authServicePb.CreateNewUser{}, err
	}
	changePasswordData := mappers.MapMessToRequest(in)
	changePasswordData.Username = user.Username
	user, err = s.AuthHandler.ChangePassword(changePasswordData)
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("%s", err.Error()), Level: logrus.WarnLevel, Component: GetComponentName(s.AuthHandler.ChangePassword)})
		return &authServicePb.CreateNewUser{}, err
	}
	s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User changed password. Username: %s, Ip address: %s", user.Username, getRequestIpAddressFromContext(ctx)), Level: logrus.InfoLevel,
		Component: GetComponentName(s.UpdateUserPassword)})
	return &authServicePb.CreateNewUser{User: mappers.MapUserToPb(user)}, err
}

func (s *Server) ActivateUserAccount(ctx context.Context, in *authServicePb.ActivationToken) (*authServicePb.ActivationResponse, error) {
	username, err := s.AuthHandler.ActivateAccount(in.Token)
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: err.Error(), Level: logrus.WarnLevel, Component: GetComponentName(s.AuthHandler.ActivateAccount)})
		return &authServicePb.ActivationResponse{ResponseStatus: err.Error()}, err
	}
	s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User activated account. Username: %s, Ip address: %s", username, getRequestIpAddressFromContext(ctx)),
		Level: logrus.InfoLevel, Component: GetComponentName(s.ActivateUserAccount)})
	return &authServicePb.ActivationResponse{ResponseStatus: "Account successfully activated!"}, err
}

func (s *Server) ForgottenPassword(ctx context.Context, in *authServicePb.UserEmailMessage) (*authServicePb.ActivationResponse, error) {
	user, err := s.AuthHandler.AccountRecoveryEmail(in.Email.Email)
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: err.Error(), Level: logrus.WarnLevel, Component: GetComponentName(s.AuthHandler.AccountRecoveryEmail)})
		return &authServicePb.ActivationResponse{ResponseStatus: err.Error()}, err
	}
	s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("Recovery email sent to %s for %s", user.Email, user.Username), Level: logrus.InfoLevel,
		Component: GetComponentName(s.ForgottenPassword)})
	return &authServicePb.ActivationResponse{ResponseStatus: "Recovery email successfully sent!"}, err
}

func (s *Server) ResetPassword(ctx context.Context, in *authServicePb.ResetPasswordWithTokenMessage) (*authServicePb.ActivationResponse, error) {
	username, err := s.AuthHandler.ResetPassword(in.Details.Token, in.Details.NewPassword)
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("%s. Username: %s, Ip address: %s", err.Error(), username, getRequestIpAddressFromContext(ctx)),
			Level: logrus.WarnLevel, Component: GetComponentName(s.AuthHandler.ResetPassword)})
		return &authServicePb.ActivationResponse{ResponseStatus: err.Error()}, err
	}
	s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User reseted password with token"), Level: logrus.InfoLevel,
		Component: GetComponentName(s.ResetPassword)})
	return &authServicePb.ActivationResponse{ResponseStatus: "Password changed successfully!"}, err
}

func (s *Server) GeneratePasswordlessLoginToken(ctx context.Context, in *authServicePb.UserEmailMessage) (*authServicePb.ActivationResponse, error) {
	username, err := s.AuthHandler.SendPasswordlessLoginEmail(in.Email.Email)
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("%s. Username: %s, Ip address: %s", err.Error(), username, getRequestIpAddressFromContext(ctx)),
			Level: logrus.WarnLevel, Component: GetComponentName(s.AuthHandler.SendPasswordlessLoginEmail)})
		return &authServicePb.ActivationResponse{ResponseStatus: err.Error()}, err
	}
	s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User generated token for passwordles login. Username: %s, Ip address: %s", username, getRequestIpAddressFromContext(ctx)), Level: logrus.InfoLevel,
		Component: GetComponentName(s.GeneratePasswordlessLoginToken)})
	return &authServicePb.ActivationResponse{ResponseStatus: "Email successfully sent!"}, err
}
func (s *Server) PasswordlessLogin(ctx context.Context, in *authServicePb.ActivationTokenMessage) (*authServicePb.ActivationResponse, error) {
	jwt, err := s.AuthHandler.PasswordlessLogin(in.Token.Token)
	if err != nil {
		return &authServicePb.ActivationResponse{ResponseStatus: err.Error()}, err
	}

	return &authServicePb.ActivationResponse{ResponseStatus: jwt.Token}, err
}
func (s *Server) GenerateApiKey(ctx context.Context, in *authServicePb.GetAllRequest) (*authServicePb.Token, error) {
	loggedinUserToken := s.getTokenFromContext(ctx)
	tokenUser, _ := s.AuthHandler.ValidateToken(loggedinUserToken)
	apiKey, err := s.AuthHandler.GenerateJWT(handlers.JWTOptions{Username: tokenUser.Username, IsTokenNonExpired: true})
	if err != nil {
		return nil, err
	}
	err = s.AuthHandler.UserStore.UpdateApiKey(tokenUser.Username, apiKey)
	if err != nil {
		return nil, err
	}
	return &authServicePb.Token{Token: apiKey}, nil
}
func (s *Server) GetUserApiKey(ctx context.Context, in *authServicePb.GetAllRequest) (*authServicePb.Token, error) {
	loggedinUserToken := s.getTokenFromContext(ctx)
	tokenUser, _ := s.AuthHandler.ValidateToken(loggedinUserToken)
	dbUser, err := s.AuthHandler.UserStore.FindByUsername(tokenUser.Username)
	if err != nil {
		return nil, err
	}
	return &authServicePb.Token{Token: dbUser.ApiToken}, nil
}

func (s *Server) getUsernameFromContext(ctx context.Context) string {
	token := s.getTokenFromContext(ctx)
	if token != "" {
		return ""
	}
	user, err := s.AuthHandler.ValidateToken(token)
	if err != nil {
		return ""
	}
	return user.Username
}
func (s *Server) getTokenFromContext(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)
	if len(md.Get("authorization")) == 0 {
		s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("Unauthorized access from %s", getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel, Component: "auth_service.main.getTokenFromContext"})
		return ""
	}
	return strings.Split(md.Get("authorization")[0], " ")[1]
}

func getRequestIpAddressFromContext(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)
	userIp := md.Get("x-forwarded-for")[0]
	return userIp
}
func GetComponentName(methode interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(methode).Pointer()).Name()
}
