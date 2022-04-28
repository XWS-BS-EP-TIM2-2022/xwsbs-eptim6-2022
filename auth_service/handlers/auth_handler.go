package handlers

import (
	"auth_service/store"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	UserStore *store.UsersStore
}

func InitAuthHandler() *AuthHandler {
	userStore := store.InitUsersStore()
	return &AuthHandler{UserStore: userStore}
}

func DecodeUser(req *http.Request) (store.User, error) {
	var user store.User
	err := json.NewDecoder(req.Body).Decode(&user)
	return user, err
}

func (ag *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	user, err := DecodeUser(r)
	if err != nil {
		println("Error while parsing json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ag.UserStore.FindByUsername(user.Username)

}

func (ag *AuthHandler) AddNewUser(w http.ResponseWriter, r *http.Request) {
	user, err := DecodeUser(r)
	if err != nil {
		println("Error while parsing json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ag.UserStore.AddNew(user)
	fmt.Println("Post request")
}

func (ag *AuthHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users := ag.UserStore.FindAll()
	json.NewEncoder(w).Encode(users)
}
