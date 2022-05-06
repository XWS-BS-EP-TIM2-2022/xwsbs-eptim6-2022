package main

import (
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/api_gateway/startup"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/api_gateway/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
