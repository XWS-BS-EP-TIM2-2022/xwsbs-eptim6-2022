package main

import (
	"context"
	postsServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/posts_service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"xwsbs-eptim6-2022/posts-service/handlers"
)

type Server struct {
	postsServicePb.UnimplementedPostsServiceServer
	postsHandler *handlers.PostsHandler
}

func (s *Server) GetAllPosts(ctx context.Context, in *postsServicePb.EmptyRequest) (*postsServicePb.PostsResponse, error) {
	return nil, nil
}
func (s *Server) GetAllPostsByUser(ctx context.Context, in *postsServicePb.GetByUsernameRequest) (*postsServicePb.PostsResponse, error) {
	return nil, nil
}
func (s *Server) GetPostById(ctx context.Context, in *postsServicePb.GetByIdRequest) (*postsServicePb.PostResponse, error) {
	return nil, nil
}
func (s *Server) LikePost(ctx context.Context, in *postsServicePb.GetByIdRequest) (*postsServicePb.PostResponse, error) {
	return nil, nil
}
func (s *Server) DislikePost(ctx context.Context, in *postsServicePb.GetByIdRequest) (*postsServicePb.PostResponse, error) {
	return nil, nil
}
func (s *Server) AddNewPost(ctx context.Context, in *postsServicePb.PostRequest) (*postsServicePb.PostResponse, error) {
	return nil, nil
}
func (s *Server) AddNewComment(ctx context.Context, in *postsServicePb.CommentRequest) (*postsServicePb.PostResponse, error) {
	return nil, nil
}
func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	service, err := NewServer()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	postsServicePb.RegisterPostsServiceServer(s, service)
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = postsServicePb.RegisterPostsServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: tracingWrapper(gwmux),
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
func NewServer() (*Server, error) {
	l := log.New(os.Stdout, "posts-service ", log.LstdFlags)
	return &Server{postsHandler: handlers.NewPostsHandler(l)}, nil
}

var grpcGatewayTag = otgo.Tag{Key: string(ext.Component), Value: "grpc-gateway"}

func tracingWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parentSpanContext, err := otgo.GlobalTracer().Extract(
			otgo.HTTPHeaders,
			otgo.HTTPHeadersCarrier(r.Header))
		if err == nil || err == otgo.ErrSpanContextNotFound {
			serverSpan := otgo.GlobalTracer().StartSpan(
				"ServeHTTP",
				// this is magical, it attaches the new span to the parent parentSpanContext, and creates an unparented one if empty.
				ext.RPCServerOption(parentSpanContext),
				grpcGatewayTag,
			)
			r = r.WithContext(otgo.ContextWithSpan(r.Context(), serverSpan))
			defer serverSpan.Finish()
		}
		h.ServeHTTP(w, r)
	})
}

/*func main() {
	l := log.New(os.Stdout, "posts-service ", log.LstdFlags)
	ph := handlers.NewPostsHandler(l)
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/posts", ph.GetAll)
	getRouter.HandleFunc("/posts/user/{username:[a-zA-Z0-9]+}", ph.GetByUser)
	getRouter.HandleFunc("/posts/{id:[a-zA-Z0-9]+}", ph.GetOne)

	getRouter.HandleFunc("/posts/like/{id:[a-zA-Z0-9]+}", ph.LikePost)
	getRouter.HandleFunc("/posts/dislike/{id:[a-zA-Z0-9]+}", ph.DislikePost)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/posts/new-post", ph.CreatePost)
	postRouter.HandleFunc("/posts/new-comment/{id:[a-zA-Z0-9]+}", ph.CommentOnPost)

	s := http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
*/
//**
//
//	f, err := os.Open("./test.jpg")
//	imageData, _, err := image.Decode(f)
//	buf := new(bytes.Buffer)
//	err = jpeg.Encode(buf, imageData, nil)
//	if err != nil {
//		log.Fatalln("ENCOIDNG IMAGE ERROR")
//	}
//	send_s3 := buf.Bytes()
//	imgHandler.SaveImage(send_s3)
//
//*//
