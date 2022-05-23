package handlers

import (
	"errors"
	"net/http"
	"profile-service/store"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	UserStore *store.UsersStore
}

func InitUserHandler() *UserHandler {
	userStore := store.InitUsersStore()
	return &UserHandler{UserStore: userStore}
}

func (uh *UserHandler) GetUser(username string) (store.User, error) {
	user, err := uh.UserStore.FindOneByUsername(username)
	if err != nil {
		return store.User{}, errors.New("Error while finding user with specified username")
	}
	return user, nil
}

func (uh *UserHandler) GetAll() []store.User {
	return uh.UserStore.FindAll()
}

func (uh *UserHandler) AddNewUser(user *store.User) *store.User {
	user.Experiences = []store.Experience{}
	user.Educations = []store.Education{}
	user.Skills = []store.Skill{}
	user.Interests = []store.Interest{}
	user.Followers = []store.Follower{}
	user.Followings = []store.Following{}
	user.FollowRequests = []string{}
	uh.UserStore.AddUser(user)
	return user
}

func (uh *UserHandler) UpdateUser(username string, newUserInfo store.User) (store.User, error) {
	user, err := uh.GetUser(username)
	id := user.ID
	err = uh.UserStore.UpdateUser(id, newUserInfo)
	if err != nil {
		return store.User{}, err
	}
	return newUserInfo, nil
}

func (uh *UserHandler) AddExperience(username string, experience store.Experience) (store.User, error) {
	user, _ := uh.GetUser(username)
	id := user.ID
	err := uh.UserStore.InsertExperience(id, experience)
	return user, err
}

func (uh *UserHandler) AddEducation(username string, text string) (store.User, error) {
	user, _ := uh.GetUser(username)
	id := user.ID
	err1 := uh.UserStore.InsertEducation(id, store.Education{Text: text})
	return user, err1
}

func (uh *UserHandler) AddSkill(username string, text string) (store.User, error) {
	user, _ := uh.GetUser(username)
	id := user.ID
	err1 := uh.UserStore.InsertSkill(id, store.Skill{Text: text})
	return user, err1
}

func (uh *UserHandler) AddInterest(username string, text string) (store.User, error) {
	user, _ := uh.GetUser(username)
	id := user.ID
	err1 := uh.UserStore.InsertInterest(id, store.Interest{Text: text})
	return user, err1
}

//TODO: TESTIRATI
func (uh *UserHandler) FollowUser(username string, userToFollowID string) error {
	user, _ := uh.GetUser(username)
	userID := user.ID
	var follower store.Follower
	var following store.Following

	var userFollower store.User
	var userFollowing store.User

	userFollower, err1 := uh.UserStore.FindOne(userID)
	if err1 != nil {
		return err1
	}
	hex, err := primitive.ObjectIDFromHex(userToFollowID)
	if err != nil {
		return err
	}
	userFollowing, err = uh.UserStore.FindOne(hex)
	if err != nil {
		return errors.New("Error getting user with {idToFollow}")
	}
	follower.Username = userFollowing.Username
	following.Username = userFollower.Username

	if userFollowing.IsPublic {
		err := uh.UserStore.FollowUser(hex, userID, follower, following)
		if err != nil {
			return errors.New("Error following user")
		}
	} else {
		err := uh.UserStore.AddFollowRequest(hex, userFollower.Username)
		if err != nil {
			return errors.New("Error adding your follow request to the user you want to follow")
		}
	}
	return nil
}

func (uh *UserHandler) UnfollowUser(username string, userToFollowID string) error {
	userToUnfollowID, _ := primitive.ObjectIDFromHex(userToFollowID)
	user, err := uh.GetUser(username)
	userID := user.ID

	var follower store.Follower
	var following store.Following

	var userFollower store.User
	var userFollowing store.User

	userFollower, err = uh.UserStore.FindOne(userID)
	if err != nil {
		return errors.New("Error getting user with {id}")
	}

	userFollowing, err = uh.UserStore.FindOne(userToUnfollowID)
	if err != nil {
		return errors.New("Error following user")
	}
	follower.Username = userFollowing.Username
	following.Username = userFollower.Username

	err = uh.UserStore.UnfollowUser(userToUnfollowID, userID, follower, following)
	if err != nil {
		return errors.New("Error following user")
	}
	return nil
}

func (uh *UserHandler) AcceptFollow(username string, id string) error {
	userToAcceptID, _ := primitive.ObjectIDFromHex(id)
	user, err := uh.GetUser(username)
	userID := user.ID
	userFromFollowRequest, err2 := uh.UserStore.FindOne(userToAcceptID) //ovaj user zeli da zaprati
	if err2 != nil {
		return errors.New("Error getting user with {idUserToAccept}")
	}
	user, err1 := uh.UserStore.FindOne(userID) //ovog usera
	if err1 != nil {
		return errors.New("Error getting user with {id}")
	}
	var follower store.Follower
	var following store.Following

	following.Username = userFromFollowRequest.Username
	follower.Username = user.Username

	err = uh.UserStore.FollowUser(userID, userFromFollowRequest.ID, follower, following)
	if err != nil {
		return errors.New("Error following user")
	}
	err = uh.UserStore.AcceptRejectFollow(userID, userFromFollowRequest.Username)
	if err != nil {
		return errors.New("Error accepting user follow")
	}
	return nil
}

func (uh *UserHandler) RejectFollow(w http.ResponseWriter, r *http.Request) error {
	return errors.New("NIJE MI SE DALO")
	/*params := mux.Vars(r)
	userToAcceptID, _ := primitive.ObjectIDFromHex(params["id"])
	username, err2 := validateLoggedinUser(r)
	if err2 != nil {
		http.Error(w, "Error while validating user", http.StatusBadRequest)
	}
	user, err3 := uh.UserStore.FindOneByUsername(username)
	if err3 != nil {
		http.Error(w, "Error while finding one user", http.StatusBadRequest)
	}
	userID := user.ID

	userFromFollowRequest, err2 := uh.UserStore.FindOne(userToAcceptID)
	if err2 != nil {
		http.Error(w, "Error getting user with {idUserToReject}", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(userFromFollowRequest)

	err3 = uh.UserStore.AcceptRejectFollow(userID, userFromFollowRequest.Username)
	if err3 != nil {
		http.Error(w, "Error rejecting user follow", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(err2)*/
}

//GET USER, EXPERIENCE, EDUCATION, SKILL, INTEREST
