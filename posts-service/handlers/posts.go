package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
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

func (p *PostsHandler) Drop(rw http.ResponseWriter, r *http.Request) {
	p.postsStore.Drop()
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

	post := store.Post{}

	err := post.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Error creating post", http.StatusBadRequest)
		return
	}
	//post.ID = getObjectId(post.ID)
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

	var curUser string // dobaviti trenutno ulogovanog korisnika
	curUser = "korisnik"

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

	var curUser string // dobaviti trenutno ulogovanog korisnika
	curUser = "korisnik"

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

	comment := store.Comment{}
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(rw, "Error while commenting post", http.StatusBadRequest)
	}

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