package startup

import (
	"context"
	"fmt"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/logger"
	authGw "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	postsServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/posts_service"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
	"xwsbs-eptim6-2022/posts-service/handlers"
	"xwsbs-eptim6-2022/posts-service/mappers"
	"xwsbs-eptim6-2022/posts-service/startup/config"
	"xwsbs-eptim6-2022/posts-service/store"
)

type Server struct {
	config *config.Config
	postsServicePb.UnimplementedPostsServiceServer
	postsHandler *handlers.PostsHandler
	logging      *logger.LoggerWrapper
}

func NewServer(config *config.Config, logging *logger.LoggerWrapper) (*Server, error) {
	l := log.New(os.Stdout, "posts-service ", log.LstdFlags)
	return &Server{postsHandler: handlers.NewPostsHandler(l, config), config: config, logging: logging}, nil
}

func (s *Server) GetAllPosts(ctx context.Context, in *postsServicePb.EmptyRequest) (*postsServicePb.PostsResponse, error) {
	all, err := s.postsHandler.GetAll()
	if err != nil {
		return &postsServicePb.PostsResponse{}, err
	}
	return mappers.MapToPostsResponse(all), nil
}
func (s *Server) GetAllPostsByUser(ctx context.Context, in *postsServicePb.GetByUsernameRequest) (*postsServicePb.PostsResponse, error) {
	all, err := s.postsHandler.GetByUser(in.Username)
	if err != nil {
		return &postsServicePb.PostsResponse{}, err
	}
	return mappers.MapToPostsResponse(all), nil
}
func (s *Server) GetPostById(ctx context.Context, in *postsServicePb.GetByIdRequest) (*postsServicePb.PostResponse, error) {
	post, err := s.postsHandler.GetOne(in.Id)
	if err != nil {
		return &postsServicePb.PostResponse{}, err
	}
	return mappers.MapToPostResponse(post), nil
}
func (s *Server) LikePost(ctx context.Context, in *postsServicePb.GetByIdRequest) (*postsServicePb.PostResponse, error) {
	username, err := s.validateLoggedinUser(s.getTokenFromContext(ctx))
	if err != nil {
		return nil, err
	}
	post, err := s.postsHandler.LikePost(in.Id, username)
	if err != nil {
		s.logging.Writeln(logger.LogMessage{Message: err.Error(), Level: logrus.FatalLevel, Component: GetComponentName(s.postsHandler.LikePost)})
		return nil, err
	}
	s.logging.Writeln(logger.LogMessage{Message: fmt.Sprintf("User: %s, liked post: %s, from IP address: %s.", username, getObjectId(in.Id), getRequestIpAddressFromContext(ctx)),
		Level: logrus.InfoLevel, Component: GetComponentName(s.postsHandler.LikePost)})
	return mappers.MapToPostResponse(post), nil
}
func (s *Server) DislikePost(ctx context.Context, in *postsServicePb.GetByIdRequest) (*postsServicePb.PostResponse, error) {
	username, err := s.validateLoggedinUser(s.getTokenFromContext(ctx))
	if err != nil {
		return nil, err
	}
	post, err := s.postsHandler.DislikePost(in.Id, username)
	if err != nil {
		s.logging.Writeln(logger.LogMessage{Message: err.Error(), Level: logrus.FatalLevel, Component: GetComponentName(s.postsHandler.DislikePost)})
		return nil, err
	}
	s.logging.Writeln(logger.LogMessage{Message: fmt.Sprintf("User: %s, disliked post: %s, from IP address: %s.", username, getObjectId(in.Id), getRequestIpAddressFromContext(ctx)),
		Level: logrus.InfoLevel, Component: GetComponentName(s.postsHandler.DislikePost)})
	return mappers.MapToPostResponse(post), nil
}
func (s *Server) AddNewPost(ctx context.Context, in *postsServicePb.PostRequest) (*postsServicePb.PostResponse, error) {
	post, err := s.postsHandler.CreatePost(in.Post)
	if err != nil {
		s.logging.Writeln(logger.LogMessage{Message: err.Error(), Level: logrus.FatalLevel, Component: GetComponentName(s.postsHandler.CreatePost)})
		return &postsServicePb.PostResponse{}, err
	}
	s.logging.Writeln(logger.LogMessage{Message: fmt.Sprintf("User: %s, created post: %s.", post.Username, post.ID),
		Level: logrus.InfoLevel, Component: GetComponentName(s.postsHandler.CreatePost)})
	return mappers.MapToPostResponse(post), nil
}
func (s *Server) AddNewComment(ctx context.Context, in *postsServicePb.CommentRequest) (*postsServicePb.PostResponse, error) {
	username, err := s.getLoggedinUsername(ctx)
	if err != nil {
		return &postsServicePb.PostResponse{}, err
	}
	err = s.postsHandler.CommentOnPost(in.Id, store.Comment{Username: username, Text: in.Comment.Text})
	if err != nil {
		s.logging.Writeln(logger.LogMessage{Message: err.Error(), Level: logrus.FatalLevel, Component: GetComponentName(s.postsHandler.CommentOnPost)})
	}
	s.logging.Writeln(logger.LogMessage{Message: fmt.Sprintf("User: %s, commented on post: %s, from IP address: %s.", username, getObjectId(in.Id), getRequestIpAddressFromContext(ctx)),
		Level: logrus.InfoLevel, Component: GetComponentName(s.postsHandler.CommentOnPost)})
	return &postsServicePb.PostResponse{}, err
}

func (s *Server) getLoggedinUsername(ctx context.Context) (string, error) {
	return s.validateLoggedinUser(s.getTokenFromContext(ctx))
}

func (s *Server) getTokenFromContext(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)
	if len(md.Get("authorization")) == 0 {
		s.logging.Writeln(logger.LogMessage{Message: fmt.Sprintf("Unauthorized access from %s", getRequestIpAddressFromContext(ctx)), Level: logrus.WarnLevel, Component: "auth_service.main.getTokenFromContext"})
		return ""
	}
	return strings.Split(md.Get("authorization")[0], " ")[1]
}

func (s *Server) validateLoggedinUser(token string) (string, error) {
	authServiceClient := InitAuthService(s.config)
	user, err := authServiceClient.AuthorizeJWT(context.TODO(), &authGw.ValidateToken{Token: &authGw.Token{Token: token}})
	return user.User.Username, err
}
func InitAuthService(config *config.Config) authGw.AuthServiceClient {
	authEndpoint := fmt.Sprintf("%s:%s", config.AuthServiceGrpcHost, config.AuthServiceGrpcPort)
	conn, err := getConnection(authEndpoint)
	if err != nil {
		fmt.Println("Fatal error init auth service connection!")
		return nil
	}
	return authGw.NewAuthServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func GetComponentName(methode interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(methode).Pointer()).Name()
}

func getRequestIpAddressFromContext(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)
	userIp := md.Get("x-forwarded-for")[0]
	return userIp
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
