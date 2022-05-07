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
	user.Experiences = []store.Experience{}
	user.Educations = []store.Education{}
	user.Skills = []store.Skill{}
	user.Interests = []store.Interest{}
	user.Followers = []store.Follower{}
	user.Followings = []store.Following{}
	user.FollowRequests = []string{}
	user1 := uh.UserStore.AddUser(user)
	json.NewEncoder(w).Encode(user1)
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var user store.User
	json.NewDecoder(r.Body).Decode(&user)

	err := uh.UserStore.UpdateUser(id, user)
	json.NewEncoder(w).Encode(err)
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
		http.Error(w, "Error adding experience", http.StatusBadRequest)
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
		http.Error(w, "Error adding skill", http.StatusBadRequest)
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
		http.Error(w, "Error adding interest", http.StatusBadRequest)
		return
	}
}

func (uh *UserHandler) FollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userToFollowID, _ := primitive.ObjectIDFromHex(params["idToFollow"])
	userID, _ := primitive.ObjectIDFromHex(params["id"])

	var follower store.Follower
	var following store.Following

	var userFollower store.User
	var userFollowing store.User

	userFollower, err1 := uh.UserStore.FindOne(userID)
	if err1 != nil {
		http.Error(w, "Error getting user with {id}", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(userFollower)

	userFollowing, err2 := uh.UserStore.FindOne(userToFollowID)
	if err2 != nil {
		http.Error(w, "Error getting user with {idToFollow}", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(userFollowing)

	follower.Username = userFollowing.Username
	following.Username = userFollower.Username

	if userFollowing.IsPublic {
		err := uh.UserStore.FollowUser(userToFollowID, userID, follower, following)
		if err != nil {
			http.Error(w, "Error following user", http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(err)
	} else {
		err := uh.UserStore.AddFollowRequest(userToFollowID, userFollower.Username)
		if err != nil {
			http.Error(w, "Error adding your follow request to the user you want to follow", http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(err)
	}
}

func (uh *UserHandler) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userToUnfollowID, _ := primitive.ObjectIDFromHex(params["idToUnfollow"])
	userID, _ := primitive.ObjectIDFromHex(params["id"])

	var follower store.Follower
	var following store.Following

	var userFollower store.User
	var userFollowing store.User

	userFollower, err1 := uh.UserStore.FindOne(userID)
	if err1 != nil {
		http.Error(w, "Error getting user with {id}", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(userFollower)

	userFollowing, err2 := uh.UserStore.FindOne(userToUnfollowID)
	if err2 != nil {
		http.Error(w, "Error getting user with {idToFollow}", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(userFollowing)

	follower.Username = userFollowing.Username
	following.Username = userFollower.Username

	err := uh.UserStore.UnfollowUser(userToUnfollowID, userID, follower, following)
	if err != nil {
		http.Error(w, "Error following user", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(err)
}

func (uh *UserHandler) AcceptFollow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userToAcceptID, _ := primitive.ObjectIDFromHex(params["idUserToAccept"])
	userID, _ := primitive.ObjectIDFromHex(params["id"])

	userFromFollowRequest, err2 := uh.UserStore.FindOne(userToAcceptID) //ovaj user zeli da zaprati
	if err2 != nil {
		http.Error(w, "Error getting user with {idUserToAccept}", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(userFromFollowRequest)

	user, err1 := uh.UserStore.FindOne(userID) //ovog usera
	if err1 != nil {
		http.Error(w, "Error getting user with {id}", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(user)

	var follower store.Follower
	var following store.Following

	following.Username = userFromFollowRequest.Username
	follower.Username = user.Username

	err := uh.UserStore.FollowUser(userID, userFromFollowRequest.ID, follower, following)
	if err != nil {
		http.Error(w, "Error following user", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(err)

	err3 := uh.UserStore.AcceptRejectFollow(userID, userFromFollowRequest.Username)
	if err3 != nil {
		http.Error(w, "Error accepting user follow", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(err3)
}

func (uh *UserHandler) RejectFollow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userToAcceptID, _ := primitive.ObjectIDFromHex(params["idUserToReject"])
	userID, _ := primitive.ObjectIDFromHex(params["id"])

	userFromFollowRequest, err2 := uh.UserStore.FindOne(userToAcceptID)
	if err2 != nil {
		http.Error(w, "Error getting user with {idUserToReject}", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(userFromFollowRequest)

	err3 := uh.UserStore.AcceptRejectFollow(userID, userFromFollowRequest.Username)
	if err3 != nil {
		http.Error(w, "Error rejecting user follow", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(err2)
}
