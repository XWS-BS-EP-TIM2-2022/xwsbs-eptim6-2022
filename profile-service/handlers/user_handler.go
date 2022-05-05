package handlers

import (
	"encoding/json"
	"net/http"
	"profile-service/store"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	UserStore *store.UsersStore
}

func InitUserHandler() *UserHandler {
	userStore := store.InitUsersStore()
	return &UserHandler{UserStore: userStore}
}

func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	user, _ := uh.UserStore.FindOne(id)
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

func (uh *UserHandler) AddExperience(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var experience store.Experience
	err := json.NewDecoder(r.Body).Decode(&experience)
	if err != nil {
		http.Error(w, "Error while adding new experience", http.StatusBadRequest)
	}

	err1 := uh.UserStore.InsertExperience(id, experience)
	if err1 != nil {
		http.Error(w, "Error commenting post", http.StatusBadRequest)
		return
	}
}

func (uh *UserHandler) AddEducation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var education store.Education
	err := json.NewDecoder(r.Body).Decode(&education)
	if err != nil {
		http.Error(w, "Error while adding new education", http.StatusBadRequest)
	}

	err1 := uh.UserStore.InsertEducation(id, education)
	if err1 != nil {
		http.Error(w, "Error adding eduction", http.StatusBadRequest)
		return
	}
}

func (uh *UserHandler) AddSkill(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var skill store.Skill
	err := json.NewDecoder(r.Body).Decode(&skill)
	if err != nil {
		http.Error(w, "Error while adding new skill", http.StatusBadRequest)
	}

	err1 := uh.UserStore.InsertSkill(id, skill)
	if err1 != nil {
		http.Error(w, "Error adding eduction", http.StatusBadRequest)
		return
	}
}

func (uh *UserHandler) AddInterest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var interest store.Interest
	err := json.NewDecoder(r.Body).Decode(&interest)
	if err != nil {
		http.Error(w, "Error while adding new interest", http.StatusBadRequest)
	}

	err1 := uh.UserStore.InsertInterest(id, interest)
	if err1 != nil {
		http.Error(w, "Error adding eduction", http.StatusBadRequest)
		return
	}
}
