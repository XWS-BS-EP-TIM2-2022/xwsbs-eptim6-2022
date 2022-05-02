package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"xwsbs-eptim6-2022/posts-service/store"
)

type PostsHandler struct {
	l          *log.Logger
	postsStore *store.PostsStore
}

func NewPostsHandler(l *log.Logger) *PostsHandler {
	postsStore := store.InitPostsStore()
	return &PostsHandler{l, postsStore}
}

func (p *PostsHandler) GetAll(rw http.ResponseWriter, r *http.Request) {

	lp := p.postsStore.GetAll()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *PostsHandler) GetByUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["username"]

	lp := p.postsStore.GetByUser(user)

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *PostsHandler) CreatePost(rw http.ResponseWriter, r *http.Request) {

	post := store.Post{}

	err := post.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Error creating post", http.StatusBadRequest)
		return
	}
	err = p.postsStore.CreatePost(post)
	if err != nil {
		http.Error(rw, "Error creating post", http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("New post created."))
}

func (p *PostsHandler) GetOne(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	post := p.postsStore.GetById(id)

	err = post.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *PostsHandler) LikePost(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	post := p.postsStore.LikePost(id)

	err = post.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *PostsHandler) DislikePost(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	post := p.postsStore.LikePost(id)

	err = post.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
