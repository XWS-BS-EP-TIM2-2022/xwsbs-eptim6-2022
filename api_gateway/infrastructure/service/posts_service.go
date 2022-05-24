package service

import (
	"bytes"
	"context"
	"fmt"
	authGw "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	postsGw "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/posts_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"image"
	"image/jpeg"
	"net/http"
	"time"
)

type PostsHandler struct {
	postsClientAddress string
	imageHandler       *ImagesHandler
	authService        *AuthService
}

func NewPostsHandler(postsClientAddress string, authService *AuthService) Handler {
	//_, err := NewImageHandler()
	//if err != nil {
	//	return nil
	//}
	return &PostsHandler{postsClientAddress: postsClientAddress,
		authService: authService,
	}
}

func (handler *PostsHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/api/posts", handler.CreatePost)
	if err != nil {
		panic(err)
	}
}

func (p *PostsHandler) CreatePost(rw http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	token := p.authService.GetAuthToken(r)
	user, err := p.authService.client.AuthorizeJWT(context.TODO(), &authGw.ValidateToken{Token: &authGw.Token{Token: token}})
	if err != nil {
		return
	}
	username := user.User.Username
	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(rw, "Form parsing error", http.StatusBadRequest)
		return
	}
	postText := r.Form.Get("text")
	fmt.Println(postText)
	var url = ""
	f, _, err := r.FormFile("file")
	if err == nil {
		defer f.Close()
		imageData, _, err := image.Decode(f)
		buf := new(bytes.Buffer)
		err = jpeg.Encode(buf, imageData, nil)
		if err != nil {
			http.Error(rw, "ENCOIDNG IMAGE ERROR", http.StatusBadRequest)
			return
		}
		url, err = p.imageHandler.SaveImage(buf.Bytes())
	}
	currentTime := time.Now().Format("02.01.2006 15:04")
	connection, err := getConnection(p.postsClientAddress)
	if err != nil {
		return
	}
	postsServiceClient := postsGw.NewPostsServiceClient(connection)
	_, err = postsServiceClient.AddNewPost(context.TODO(), &postsGw.PostRequest{Post: &postsGw.Post{Username: username, ImageUrl: url, Text: postText, Liked: []string{}, Disliked: []string{}, CreatedOn: currentTime}})
	if err != nil {
		http.Error(rw, "Could not create post", http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Your post has been published."))
}
