package handlers

import (
	"errors"
	postsServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/posts_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"xwsbs-eptim6-2022/posts-service/startup/config"
	"xwsbs-eptim6-2022/posts-service/store"
)

type PostsHandler struct {
	postsStore   *store.PostsStore
	imageHandler *ImagesHandler
}

func NewPostsHandler(l *log.Logger, config *config.Config) *PostsHandler {
	postsStore := store.InitPostsStore(config.MongoDbUri)
	//imagesHandler, err := InitImageHandler()
	//if err != nil {
	//	l.Fatalln("Firebase storage error")
	//}
	return &PostsHandler{postsStore, nil}
}

func (p *PostsHandler) GetAll() (store.Posts, error) {
	return p.postsStore.GetAll()
}

func (p *PostsHandler) GetByUser(username string) (store.Posts, error) {
	return p.postsStore.GetByUser(username)
}

func (p *PostsHandler) CreatePost(in *postsServicePb.Post) (store.Post, error) {
	post := store.Post{Username: in.Username, ImageUrl: in.ImageUrl, Text: in.Text, Liked: []string{}, Disliked: []string{}, Comments: []store.Comment{}, CreatedOn: in.CreatedOn}
	err, id := p.postsStore.CreatePost(post)
	if err != nil {
		return store.Post{}, errors.New("Could not create")
	}
	post.ID = id
	return post, nil
}

func (p *PostsHandler) GetOne(id string) (store.Post, error) {
	return p.postsStore.GetById(getObjectId(id))
}

func (p *PostsHandler) LikePost(id string, username string) (store.Post, error) {
	liked := p.postsStore.IsAlreadyLiked(getObjectId(id), username)
	var post store.Post
	var err error
	if liked == false {
		post, err = p.postsStore.LikePost(getObjectId(id), username)
	} else {
		post, err = p.postsStore.UnlikePost(getObjectId(id), username)
	}
	if err != nil {
		return store.Post{}, err
	}
	return post, nil
}

func (p *PostsHandler) DislikePost(id string, curUser string) (store.Post, error) {
	disliked := p.postsStore.IsAlreadyDisliked(getObjectId(id), curUser)
	var post store.Post
	var err error
	if disliked == false {
		post, err = p.postsStore.DislikePost(getObjectId(id), curUser)
	} else {
		post, err = p.postsStore.UndislikePost(getObjectId(id), curUser)
	}
	if err != nil {
		return store.Post{}, err
	}
	return post, nil
}

func (p *PostsHandler) CommentOnPost(id string, comment store.Comment) error {
	return p.postsStore.InsertComment(getObjectId(id), comment)
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
