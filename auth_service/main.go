package main

import (
	"auth_service/handlers"
	"fmt"
	han "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := RegisterRouts()
	credentials := han.AllowCredentials()
	methods := han.AllowedMethods([]string{"POST", "PUT", "GET", "DELETE"})
	//ttl := han.MaxAge(3600)
	origins := han.AllowedOrigins([]string{"http://localhost:4200/**", "http://localhost:4200"})
	fmt.Println("START Listening")
	log.Fatal(http.ListenAndServe(":8080", han.CORS(credentials, methods, origins)(router)))
}

func RegisterRouts() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	rg := handlers.InitAuthHandler()
	router.HandleFunc("/api/auth/users", rg.GetAll).Methods("GET")
	router.HandleFunc("/api/auth/users", rg.AddNewUser).Methods("POST")
	router.HandleFunc("/api/auth/session", rg.LoginUser).Methods("PUT")
	router.HandleFunc("/api/auth/session/validations", rg.AuthorizeJWT).Methods("PUT")
	return router
}
