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
	//ph := handlers.InitProfileHandler()

	myRouter.HandleFunc("/user/{id}", uh.GetUser).Methods("GET")
	myRouter.HandleFunc("/user", uh.AddNewUser).Methods("POST")
	myRouter.HandleFunc("/users", uh.GetAll).Methods("GET")
	myRouter.HandleFunc("/user/{id}", uh.UpdateUser).Methods("PUT")

	myRouter.HandleFunc("/user/{id}/experience", uh.AddExperience).Methods("POST")
	myRouter.HandleFunc("/user/{id}/education", uh.AddEducation).Methods("POST")
	myRouter.HandleFunc("/user/{id}/skill", uh.AddSkill).Methods("POST")
	myRouter.HandleFunc("/user/{id}/interest", uh.AddInterest).Methods("POST")

	//u post telo ide ID usera koji zeli da zaprati ID usera {id}
	//{id} usera koga zelim da zapratim (meni se dodaje u following a njemu u followers)
	myRouter.HandleFunc("/user/follow/{id}", uh.FollowUser).Methods("POST")
	//u post telo ide ID usera koji zeli da otprati ID usera {id}
	myRouter.HandleFunc("/user/unfollow/{id}", uh.UnfollowUser).Methods("POST")
	//u post telo ide ID usera koji zeli da prihvati zahtev od usera sa id-jem {id}
	myRouter.HandleFunc("/user/accept-follow-request/{id}", uh.AcceptFollow).Methods("POST")
	//u post telo ide ID usera koji zeli da odbije zahtev od usera sa id-jem {id}
	myRouter.HandleFunc("/user/reject-follow-request/{id}", uh.RejectFollow).Methods("POST")

	//myRouter.HandleFunc("/profile", ph.getProfile).Methods("GET")
	return myRouter
}

func main() {
	router := RegisterRouts()
	fmt.Println("START Listening")
	log.Fatal(http.ListenAndServe(":8081", router))
}
