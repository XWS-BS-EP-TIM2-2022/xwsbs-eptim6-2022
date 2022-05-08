package main

import (
	"fmt"
	"log"
	"net/http"
	"profile-service/handlers"

	"github.com/gorilla/mux"
)

func RegisterRouts() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)

	uh := handlers.InitUserHandler()

	myRouter.HandleFunc("/users/{id}", uh.GetUser).Methods("GET")
	myRouter.HandleFunc("/users", uh.AddNewUser).Methods("POST")
	myRouter.HandleFunc("/users", uh.GetAll).Methods("GET")
	myRouter.HandleFunc("/users/{id}", uh.UpdateUser).Methods("PUT")

	myRouter.HandleFunc("/users/{id}/experience", uh.AddExperience).Methods("POST")
	myRouter.HandleFunc("/users/{id}/education", uh.AddEducation).Methods("POST")
	myRouter.HandleFunc("/users/{id}/skill", uh.AddSkill).Methods("POST")
	myRouter.HandleFunc("/users/{id}/interest", uh.AddInterest).Methods("POST")

	myRouter.HandleFunc("/users/follow/{id}/{idToFollow}", uh.FollowUser).Methods("PUT")
	myRouter.HandleFunc("/users/unfollow/{id}/{idToUnfollow}", uh.UnfollowUser).Methods("PUT")
	myRouter.HandleFunc("/users/accept-follow-request/{id}/{idUserToAccept}", uh.AcceptFollow).Methods("PUT")
	myRouter.HandleFunc("/users/reject-follow-request/{id}/{idUserToReject}", uh.RejectFollow).Methods("PUT")

	return myRouter
}

func main() {
	router := RegisterRouts()
	fmt.Println("START Listening")
	log.Fatal(http.ListenAndServe(":8081", router))
}
