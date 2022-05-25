package startup

import (
	"context"
	"fmt"
	authGw "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	usersServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/profile_service"
	"google.golang.org/grpc/metadata"
	"profile-service/handlers"
	"profile-service/infrastructure"
	"profile-service/mappers"
	"profile-service/startup/config"
	"profile-service/store"
	"strings"
)

type Server struct {
	usersServicePb.UnsafeProfileServiceServer
	userHandler       *handlers.UserHandler
	authServiceClient authGw.AuthServiceClient
}

func NewServer(config *config.Config) (*Server, error) {
	return &Server{userHandler: handlers.InitUserHandler(*config), authServiceClient: infrastructure.InitAuthServiceClient(config)}, nil
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
	return mappers.MapToUsersResponse(s.userHandler.GetAll()), nil
}
func (s *Server) AddNewUser(ctx context.Context, in *usersServicePb.UserRequest) (*usersServicePb.UserResponse, error) {
	usr := s.userHandler.AddNewUser(mappers.MapToUser(in.User))
	return mappers.MapToUserResponse(usr), nil
}
func (s *Server) AddSkill(ctx context.Context, in *usersServicePb.SkillRequest) (*usersServicePb.UserResponse, error) {
	username, err := s.validateLoggedinUser(getTokenFromContext(ctx))
	if err != nil {
		return nil, err
	}
	user, err := s.userHandler.AddSkill(username, in.Skill.Text)
	if err != nil {
		return mappers.MapToUserResponse(&user), err
	}
	return mappers.MapToUserResponse(&user), nil
}
func (s *Server) AddInterest(ctx context.Context, in *usersServicePb.InterestRequest) (*usersServicePb.UserResponse, error) {
	username, err := s.validateLoggedinUser(getTokenFromContext(ctx))
	if err != nil {
		return nil, err
	}
	user, err := s.userHandler.AddInterest(username, in.Interest.Text)
	if err != nil {
		return mappers.MapToUserResponse(&user), err
	}
	return mappers.MapToUserResponse(&user), nil

}
func (s *Server) AddExperience(ctx context.Context, in *usersServicePb.ExperienceRequest) (*usersServicePb.UserResponse, error) {
	username, err := s.validateLoggedinUser(getTokenFromContext(ctx))
	if err != nil {
		return nil, err
	}
	user, err := s.userHandler.AddExperience(username, store.Experience{Text: in.Experience.Text})
	if err != nil {
		return mappers.MapToUserResponse(&user), err
	}
	return mappers.MapToUserResponse(&user), nil

}
func (s *Server) AddEducation(ctx context.Context, in *usersServicePb.EducationRequest) (*usersServicePb.UserResponse, error) {
	username, err := s.validateLoggedinUser(getTokenFromContext(ctx))
	if err != nil {
		return nil, err
	}
	user, err := s.userHandler.AddEducation(username, in.Education.Text)
	if err != nil {
		return mappers.MapToUserResponse(&user), err
	}
	return mappers.MapToUserResponse(&user), nil

}
func (s *Server) FollowUser(ctx context.Context, in *usersServicePb.GetByIdRequest) (*usersServicePb.UserResponse, error) {
	username, err := s.validateLoggedinUser(getTokenFromContext(ctx))
	if err != nil {
		return nil, err
	}
	err = s.userHandler.FollowUser(username, in.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
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
	return nil, nil
}
func (s *Server) AcceptFollow(ctx context.Context, in *usersServicePb.GetByIdRequest) (*usersServicePb.UserResponse, error) {
	return nil, nil
}
func (s *Server) RejectFollow(ctx context.Context, in *usersServicePb.GetByIdRequest) (*usersServicePb.UserResponse, error) {
	return nil, nil
}
