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
	//myRouter.HandleFunc("/profile", ph.getProfile).Methods("GET")
	return myRouter
}

func main() {
	router := RegisterRouts()
	fmt.Println("START Listening")
	log.Fatal(http.ListenAndServe(":8081", router))
}
