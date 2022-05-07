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

	myRouter.HandleFunc("/user/{id}", uh.GetUser).Methods("GET")
	myRouter.HandleFunc("/user", uh.AddNewUser).Methods("POST")
	myRouter.HandleFunc("/users", uh.GetAll).Methods("GET")
	myRouter.HandleFunc("/user/{id}", uh.UpdateUser).Methods("PUT")

	myRouter.HandleFunc("/user/{id}/experience", uh.AddExperience).Methods("POST")
	myRouter.HandleFunc("/user/{id}/education", uh.AddEducation).Methods("POST")
	myRouter.HandleFunc("/user/{id}/skill", uh.AddSkill).Methods("POST")
	myRouter.HandleFunc("/user/{id}/interest", uh.AddInterest).Methods("POST")

	myRouter.HandleFunc("/user/follow/{id}/{idToFollow}", uh.FollowUser).Methods("PUT")
	myRouter.HandleFunc("/user/unfollow/{id}/{idToUnfollow}", uh.UnfollowUser).Methods("PUT")
	myRouter.HandleFunc("/user/accept-follow-request/{id}/{idUserToAccept}", uh.AcceptFollow).Methods("PUT")
	//u post telo ide ID usera koji zeli da odbije zahtev od usera sa id-jem {id}
	myRouter.HandleFunc("/user/reject-follow-request/{id}", uh.RejectFollow).Methods("POST")

	return myRouter
}

func main() {
	router := RegisterRouts()
	fmt.Println("START Listening")
	log.Fatal(http.ListenAndServe(":8081", router))
}
