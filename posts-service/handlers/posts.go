package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"xwsbs-eptim6-2022/posts-service/store"
)

type PostsHandler struct {
	l            *log.Logger
	postsStore   *store.PostsStore
	imageHandler *ImagesHandler
}

func NewPostsHandler(l *log.Logger) *PostsHandler {
	postsStore := store.InitPostsStore()
	imagesHandler, err := InitImageHandler()
	if err != nil {
		l.Fatalln("Firebase storage error")
	}
	return &PostsHandler{l, postsStore, imagesHandler}
}

func (p *PostsHandler) GetAll(rw http.ResponseWriter, r *http.Request) {

	lp, err := p.postsStore.GetAll()
	if err != nil {
		http.Error(rw, "Error while fetching posts", http.StatusBadRequest)
	}

	// serialize the list to JSON
	err = lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *PostsHandler) GetByUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["username"]

	lp, err := p.postsStore.GetByUser(user)
	if err != nil {
		http.Error(rw, "Error while fetching posts", http.StatusBadRequest)
	}

	// serialize the list to JSON
	err = lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *PostsHandler) CreatePost(rw http.ResponseWriter, r *http.Request) {
	username, err := validateLoggedinUser(r)
	if err != nil {
		http.Error(rw, "Login error", http.StatusBadRequest)
		return
	}
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
	post := store.Post{Username: username, ImageUrl: url, Text: postText}
	err = p.postsStore.CreatePost(post)
	if err != nil {
		http.Error(rw, "Could not create post", http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(post)
}

func (p *PostsHandler) GetOne(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	post, err := p.postsStore.GetById(getObjectId(id))
	if err != nil {
		http.Error(rw, "Could not get post", http.StatusBadRequest)
		return
	}

	err = post.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *PostsHandler) LikePost(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	curUser, err := validateLoggedinUser(r)
	if err != nil {
		http.Error(rw, "Login error", http.StatusBadRequest)
		return
	}

	liked := p.postsStore.IsAlreadyLiked(getObjectId(id), curUser)
	if liked == false {
		post, err := p.postsStore.LikePost(getObjectId(id), curUser)
		if err != nil {
			http.Error(rw, "Could not like post", http.StatusBadRequest)
			return
		}

		err = post.ToJSON(rw)
		if err != nil {
			http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		}
	}
	if liked == true {
		post, err := p.postsStore.UnlikePost(getObjectId(id), curUser)
		if err != nil {
			http.Error(rw, "Could not unlike post", http.StatusBadRequest)
			return
		}

		err = post.ToJSON(rw)
		if err != nil {
			http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		}
	}
}

func (p *PostsHandler) DislikePost(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	curUser, err := validateLoggedinUser(r)
	if err != nil {
		http.Error(rw, "Login error", http.StatusBadRequest)
		return
	}

	disliked := p.postsStore.IsAlreadyDisliked(getObjectId(id), curUser)

	if disliked == false {
		post, err := p.postsStore.DislikePost(getObjectId(id), curUser)
		if err != nil {
			http.Error(rw, "Could not dislike post", http.StatusBadRequest)
			return
		}

		err = post.ToJSON(rw)
		if err != nil {
			http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		}
	}
	if disliked == true {
		post, err := p.postsStore.UndislikePost(getObjectId(id), curUser)
		if err != nil {
			http.Error(rw, "Could not undislike post", http.StatusBadRequest)
			return
		}

		err = post.ToJSON(rw)
		if err != nil {
			http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		}
	}
}

func (p *PostsHandler) CommentOnPost(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	curUser, err := validateLoggedinUser(r)
	if err != nil {
		http.Error(rw, "Login error", http.StatusBadRequest)
		return
	}

	comment := store.Comment{}
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(rw, "Error while commenting post", http.StatusBadRequest)
	}
	comment.Username = curUser
	err = p.postsStore.InsertComment(getObjectId(id), comment)
	if err != nil {
		http.Error(rw, "Error commenting post", http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Your comment has been submitted."))
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

func validateLoggedinUser(r *http.Request) (string, error) {
	client := &http.Client{}
	var authServiceHost = "http://localhost:8080/api/auth/session/validations"
	jsonBody, err := json.Marshal(map[string]string{
		"token": r.Header["Authorization"][0],
	})
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(http.MethodPut, authServiceHost, bytes.NewBuffer(jsonBody))
	req.Header.Set("Authorization", r.Header["Authorization"][0])
	resp, err := client.Do(req)
	var user store.User
	json.NewDecoder(resp.Body).Decode(&user)
	fmt.Println(user)
	if err != nil {
		return "", err
	}
	return user.Username, nil
}
