package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"xwsbs-eptim6-2022/posts-service/store"
)

type PostsHandler struct {
	l *log.Logger
}

func NewPostsHandler(l *log.Logger) *PostsHandler {
	return &PostsHandler{l}
}

func (p *PostsHandler) GetAll(rw http.ResponseWriter, r *http.Request) {

	lp := store.GetPosts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *PostsHandler) CreatePost(rw http.ResponseWriter, r *http.Request) {

	post := r.Context().Value(KeyProduct{}).(store.Post)
	store.CreatePost(&post)
}

func (p *PostsHandler) GetOne(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	post, _, err := store.FindPost(id)

	if err == store.ErrProductNotFound {
		http.Error(rw, "Post not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Post not found", http.StatusInternalServerError)
		return
	}

	err = post.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

type KeyProduct struct{}
