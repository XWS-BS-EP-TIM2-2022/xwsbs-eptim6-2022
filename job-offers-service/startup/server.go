package startup

import (
	"context"
	"fmt"
	authGw "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	jobOffersPb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/job_offers_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"job-offers-service/handlers"
	"job-offers-service/mappers"
	"job-offers-service/startup/config"
	"strings"
)

type Server struct {
	config *config.Config
	jobOffersPb.UnimplementedJobOffersServiceServer
	jobOffersHandler *handlers.JobOffersHandler
}

func NewServer(config config.Config) (*Server, error) {
	return &Server{config: &config, jobOffersHandler: handlers.NewOffersHandler(config)}, nil
}
func (s *Server) GetAllJobOffers(ctx context.Context, in *jobOffersPb.EmptyRequest) (*jobOffersPb.JobOffersResponse, error) {
	all, err := s.jobOffersHandler.GetAll()
	if err != nil {
		return nil, err
	}
	return mappers.MapToResponses(all), nil
}
func (s *Server) GetAllJobOffersByUsername(ctx context.Context, in *jobOffersPb.UsernameMessage) (*jobOffersPb.JobOffersResponse, error) {
	return nil, nil
}
func (s *Server) CreateNewJobOffer(ctx context.Context, in *jobOffersPb.CreateJobOfferMessage) (*jobOffersPb.JobOfferResponse, error) {
	offer := mappers.MapToStore(in.Offer)
	offer.User, _ = s.validateLoggedinUser(getTokenFromContext(ctx))
	jobOffer, err := s.jobOffersHandler.CreateJobOffer(&offer)
	if err != nil {
		return nil, err
	}
	return mappers.MapToResponse(jobOffer), nil
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
