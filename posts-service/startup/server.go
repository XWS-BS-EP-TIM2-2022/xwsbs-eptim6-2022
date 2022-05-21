package startup

import (
	"context"
	"fmt"
	authGw "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	postsServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/posts_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
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
}

func NewServer(config *config.Config) (*Server, error) {
	l := log.New(os.Stdout, "posts-service ", log.LstdFlags)
	return &Server{postsHandler: handlers.NewPostsHandler(l, config), config: config}, nil
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
	username, err := s.validateLoggedinUser(getTokenFromContext(ctx))
	if err != nil {
		return nil, err
	}
	post, err := s.postsHandler.LikePost(in.Id, username)
	if err != nil {
		return nil, err
	}
	return mappers.MapToPostResponse(post), nil
}
func (s *Server) DislikePost(ctx context.Context, in *postsServicePb.GetByIdRequest) (*postsServicePb.PostResponse, error) {
	username, err := s.validateLoggedinUser(getTokenFromContext(ctx))
	if err != nil {
		return nil, err
	}
	post, err := s.postsHandler.DislikePost(in.Id, username)
	if err != nil {
		return nil, err
	}
	return mappers.MapToPostResponse(post), nil
}
func (s *Server) AddNewPost(ctx context.Context, in *postsServicePb.PostRequest) (*postsServicePb.PostResponse, error) {
	post, err := s.postsHandler.CreatePost(in.Post)
	if err != nil {
		return &postsServicePb.PostResponse{}, err
	}
	return mappers.MapToPostResponse(post), nil
}
func (s *Server) AddNewComment(ctx context.Context, in *postsServicePb.CommentRequest) (*postsServicePb.PostResponse, error) {
	username, err := s.getLoggedinUsername(ctx)
	if err != nil {
		return &postsServicePb.PostResponse{}, err
	}
	err = s.postsHandler.CommentOnPost(in.Id, store.Comment{Username: username, Text: in.Comment.Text})
	return &postsServicePb.PostResponse{}, err
}

func (s *Server) getLoggedinUsername(ctx context.Context) (string, error) {
	return s.validateLoggedinUser(getTokenFromContext(ctx))
}

func getTokenFromContext(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Println(md.Get("authorization")[0])
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
