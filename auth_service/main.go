package main

import (
	"auth_service/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := RegisterRouts()
	fmt.Println("START Listening")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func RegisterRouts() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	rg := handlers.InitAuthHandler()
	router.HandleFunc("/api/auth/users", rg.GetAll).Methods("GET")
	router.HandleFunc("/api/auth/users", rg.AddNewUser).Methods("POST")
	router.HandleFunc("/api/auth/session", rg.LoginUser).Methods("POST")
	return router
}
