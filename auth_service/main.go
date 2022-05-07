package main

import (
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/startup"
	cfg "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}

/*router := RegisterRouts()
fmt.Println("START Listening")
log.Fatal(http.ListenAndServe(":8080", router))

router.HandleFunc("/api/auth/users", rg.GetAll).Methods("GET")
	router.HandleFunc("/api/auth/users", rg.AddNewUser).Methods("POST")
	router.HandleFunc("/api/auth/session", rg.LoginUser).Methods("PUT")
	router.HandleFunc("/api/auth/session/validations", rg.AuthorizeJWT).Methods("PUT")

*/
