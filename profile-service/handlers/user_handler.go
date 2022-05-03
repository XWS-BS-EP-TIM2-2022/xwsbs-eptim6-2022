package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"profile-service/store"
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
	username, _ := DecodeUsername(r)
	user, _ := uh.UserStore.FindOne(username)
	json.NewEncoder(w).Encode(user)
}

func (uh *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users := uh.UserStore.FindAll()
	json.NewEncoder(w).Encode(users)
}

func (uh *UserHandler) AddNewUser(w http.ResponseWriter, r *http.Request) {
	user, err := DecodeUser(r)
	if _, err := uh.UserStore.FindOne(user.Username); err == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Message: "Username already in use"})
		return
	}
	if err != nil {
		println("Error while parsing json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uh.UserStore.AddUser(user)
	fmt.Println("Post request")
}

func DecodeUser(req *http.Request) (store.User, error) {
	var user store.User
	err := json.NewDecoder(req.Body).Decode(&user)
	return user, err
}

func DecodeUsername(req *http.Request) (string, error) {
	var username string
	err := json.NewDecoder(req.Body).Decode(&username)
	return username, err
}
