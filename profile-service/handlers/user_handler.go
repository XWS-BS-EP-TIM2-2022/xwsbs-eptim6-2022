package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"profile-service/store"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	UserStore *store.UsersStore
}
type ErrorMessage struct {
	Message string `json:"message"`
}

func InitUserHandler() *UserHandler {
	userStore := store.InitUsersStore()
	return &UserHandler{UserStore: userStore}
}

func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	//user, _ := uh.UserStore.FindOne(id)
	var user store.User
	filter := bson.D{{Key: "id", Value: id}}
	err := uh.UserStore.UsersCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(user)
}

func (uh *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users := uh.UserStore.FindAll()
	json.NewEncoder(w).Encode(users)
}

func (uh *UserHandler) AddNewUser(w http.ResponseWriter, r *http.Request) {
	var user store.User
	json.NewDecoder(r.Body).Decode(&user)
	user1 := uh.UserStore.AddUser(user)
	json.NewEncoder(w).Encode(user1)
}
