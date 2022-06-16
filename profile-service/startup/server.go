package startup

import (
	"context"
	"fmt"
	"profile-service/handlers"
	"profile-service/infrastructure"
	"profile-service/mappers"
	"profile-service/startup/config"
	"profile-service/store"
	"reflect"
	"runtime"
	"strings"

	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/logger"
	authGw "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	usersServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/profile_service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	usersServicePb.UnsafeProfileServiceServer
	userHandler       *handlers.UserHandler
	authServiceClient authGw.AuthServiceClient
	log               *logger.LoggerWrapper
}

func NewServer(config *config.Config, log *logger.LoggerWrapper) (*Server, error) {
	initHandler, err := handlers.InitUserHandler(*config)
	if err != nil {
		return nil, err
	}
	return &Server{userHandler: initHandler, authServiceClient: infrastructure.InitAuthServiceClient(config), log: log}, nil
}

func getTokenFromContext(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Println(md.Get("authorization")[0])
	return strings.Split(md.Get("authorization")[0], " ")[1]
}
func (s *Server) validateLoggedinUser(token string) (string, error) {
	user, err := s.authServiceClient.AuthorizeJWT(context.TODO(), &authGw.ValidateToken{Token: &authGw.Token{Token: token}})
	return user.User.Username, err
}

func (s *Server) GetAllUsers(ctx context.Context, in *usersServicePb.EmptyRequest) (*usersServicePb.UsersResponse, error) {
	getRequestIpAddressFromContext(ctx)
	return mappers.MapToUsersResponse(s.userHandler.GetAll()), nil
}
func (s *Server) AddNewUser(ctx context.Context, in *usersServicePb.UserRequest) (*usersServicePb.UserResponse, error) {
	usr, err := s.userHandler.AddNewUser(mappers.MapToUser(in.User))
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: err.Error(), Level: logrus.FatalLevel, Component: GetComponentName(s.userHandler.AddNewUser)})
		return nil, err
	}
	s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User created. Username %s, Ip address: %s", usr.Username, getRequestIpAddressFromContext(ctx)),
		Level: logrus.InfoLevel, Component: GetComponentName(s.AddNewUser)})
	return mappers.MapToUserResponse(usr), nil
}
func (s *Server) AddSkill(ctx context.Context, in *usersServicePb.SkillRequest) (*usersServicePb.UserResponse, error) {
	username, err := s.validateLoggedinUser(getTokenFromContext(ctx))
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User validation failed. %s. Ip address: %s ", err.Error(), getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
			Component: "profile_service.startup.server.validateLoggedInUser"})
		return nil, err
	}
	user, err := s.userHandler.AddSkill(username, in.Skill.Text)
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User: %s adding skill failed. %s. Ip address: %s ", username, err.Error(), getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
			Component: GetComponentName(s.userHandler.AddSkill)})
		return mappers.MapToUserResponse(&user), err
	}
	s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User: %s  added skill. %s. Ip address: %s ", username, err.Error(), getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
		Component: GetComponentName(s.userHandler.AddSkill)})
	return mappers.MapToUserResponse(&user), nil
}
func (s *Server) AddInterest(ctx context.Context, in *usersServicePb.InterestRequest) (*usersServicePb.UserResponse, error) {
	username, err := s.validateLoggedinUser(getTokenFromContext(ctx))
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User validation failed. %s. Ip address: %s ", err.Error(), getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
			Component: "profile_service.startup.server.validateLoggedInUser"})
		return nil, err
	}
	user, err := s.userHandler.AddInterest(username, in.Interest.Text)
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User: %s adding interest failed. %s. Ip address: %s ", username, err.Error(), getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
			Component: GetComponentName(s.userHandler.AddInterest)})
		return mappers.MapToUserResponse(&user), err
	}
	s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User: %s  added interest. %s. Ip address: %s ", username, err.Error(), getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
		Component: GetComponentName(s.userHandler.AddInterest)})
	return mappers.MapToUserResponse(&user), nil

}
func (s *Server) AddExperience(ctx context.Context, in *usersServicePb.ExperienceRequest) (*usersServicePb.UserResponse, error) {
	username, err := s.validateLoggedinUser(getTokenFromContext(ctx))
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User validation failed. %s. Ip address: %s ", err.Error(), getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
			Component: "profile_service.startup.server.validateLoggedInUser"})
		return nil, err
	}
	user, err := s.userHandler.AddExperience(username, store.Experience{Text: in.Experience.Text})
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User: %s adding experience failed. %s. Ip address: %s ", username, err.Error(), getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
			Component: GetComponentName(s.userHandler.AddExperience)})
		return mappers.MapToUserResponse(&user), err
	}
	s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User: %s  added experience. %s. Ip address: %s ", username, err.Error(), getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
		Component: GetComponentName(s.userHandler.AddExperience)})
	return mappers.MapToUserResponse(&user), nil

}
func (s *Server) AddEducation(ctx context.Context, in *usersServicePb.EducationRequest) (*usersServicePb.UserResponse, error) {
	username, err := s.validateLoggedinUser(getTokenFromContext(ctx))
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User validation failed. %s. Ip address: %s ", err.Error(), getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
			Component: "profile_service.startup.server.validateLoggedInUser"})
		return nil, err
	}
	user, err := s.userHandler.AddEducation(username, in.Education.Text)
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User: %s adding education failed. %s. Ip address: %s ", username, err.Error(), getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
			Component: GetComponentName(s.userHandler.AddEducation)})
		return mappers.MapToUserResponse(&user), err
	}
	s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User: %s  added education. %s. Ip address: %s ", username, err.Error(), getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
		Component: GetComponentName(s.userHandler.AddEducation)})
	return mappers.MapToUserResponse(&user), nil

}
func (s *Server) FollowUser(ctx context.Context, in *usersServicePb.GetByIdRequest) (*usersServicePb.UserResponse, error) {
	username, err := s.validateLoggedinUser(getTokenFromContext(ctx))
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User validation failed. %s. Ip address: %s ", err.Error(), getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
			Component: "profile_service.startup.server.validateLoggedInUser"})
		return nil, err
	}
	err = s.userHandler.FollowUser(username, in.Id)
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User: %s following user with id: %s failed. %s. Ip address: %s ", username, in.Id, err.Error(), getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
			Component: GetComponentName(s.userHandler.FollowUser)})
		return nil, err
	}
	s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User: %s followed user with id: %s. %s. Ip address: %s ", username, in.Id, err.Error(), getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
		Component: GetComponentName(s.userHandler.FollowUser)})
	user, err := s.userHandler.GetUser(username)
	out := mappers.MapToUserResponse(&user)
	return out, nil
}

func (s *Server) GetUserById(ctx context.Context, in *usersServicePb.GetByIdRequest) (*usersServicePb.UserResponse, error) {
	return nil, nil
}
func (s *Server) FollowResponse(context.Context, *usersServicePb.GetByIdRequest) (*usersServicePb.UserResponse, error) {
	return nil, nil
}
func (s *Server) GetLoggedinUser(ctx context.Context, in *usersServicePb.EmptyRequest) (*usersServicePb.UserResponse, error) {
	username, _ := s.validateLoggedinUser(getTokenFromContext(ctx))
	user, err := s.userHandler.GetUser(username)
	if err != nil {
		return &usersServicePb.UserResponse{}, err
	}
	out := mappers.MapToUserResponse(&user)
	return out, nil
}

func (s *Server) UnFollowUser(ctx context.Context, in *usersServicePb.GetByIdRequest) (*usersServicePb.UserResponse, error) {
	username, err := s.validateLoggedinUser(getTokenFromContext(ctx))
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User validation failed. %s. Ip address: %s ", err.Error(), getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
			Component: "profile_service.startup.server.validateLoggedInUser"})
		return nil, err
	}
	err = s.userHandler.UnfollowUser(username, in.Id)
	if err != nil {
		s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User: %s unfollowing user with id: %s failed. %s. Ip address: %s ", username, in.Id, err.Error(), getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
			Component: GetComponentName(s.userHandler.UnfollowUser)})
		return nil, err
	}
	s.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("User: %s unfollowed user with id: %s. %s. Ip address: %s ", err.Error(), username, in.Id, getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel,
		Component: GetComponentName(s.userHandler.UnfollowUser)})
	user, err := s.userHandler.GetUser(username)
	out := mappers.MapToUserResponse(&user)
	return out, nil
}
func (s *Server) AcceptFollow(ctx context.Context, in *usersServicePb.GetByIdRequest) (*usersServicePb.UserResponse, error) {
	return nil, nil
}
func (s *Server) RejectFollow(ctx context.Context, in *usersServicePb.GetByIdRequest) (*usersServicePb.UserResponse, error) {
	return nil, nil
}
func getRequestIpAddressFromContext(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)
	userIp := md.Get("x-forwarded-for")[0]
	return userIp
}
func GetComponentName(methode interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(methode).Pointer()).Name()
}
