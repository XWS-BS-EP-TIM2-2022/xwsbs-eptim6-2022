package main

import (
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/api_gateway/startup"
	cfg "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}

/*router := RegisterRouts()
fmt.Println("START Listening")
log.Fatal(http.ListenAndServe(":8080", router))*/
