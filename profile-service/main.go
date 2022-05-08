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

	myRouter.HandleFunc("/user", uh.GetUser).Methods("GET")
	myRouter.HandleFunc("/users", uh.AddNewUser).Methods("POST")
	myRouter.HandleFunc("/users", uh.GetAll).Methods("GET")
	myRouter.HandleFunc("/users", uh.UpdateUser).Methods("PUT")

	myRouter.HandleFunc("/users/experience", uh.AddExperience).Methods("POST")
	myRouter.HandleFunc("/users/education", uh.AddEducation).Methods("POST")
	myRouter.HandleFunc("/users/skill", uh.AddSkill).Methods("POST")
	myRouter.HandleFunc("/users/interest", uh.AddInterest).Methods("POST")

	myRouter.HandleFunc("/users/follow/{id}", uh.FollowUser).Methods("PUT")
	myRouter.HandleFunc("/users/unfollow/{id}", uh.UnfollowUser).Methods("PUT")
	myRouter.HandleFunc("/users/accept-follow-request/{id}", uh.AcceptFollow).Methods("PUT")
	myRouter.HandleFunc("/users/reject-follow-request/{id}", uh.RejectFollow).Methods("PUT")

	return myRouter
}

func main() {
	router := RegisterRouts()
	fmt.Println("START Listening")
	log.Fatal(http.ListenAndServe(":8081", router))
}
