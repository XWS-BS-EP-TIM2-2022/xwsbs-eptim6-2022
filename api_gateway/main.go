package main

import (
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/api_gateway/startup"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/api_gateway/startup/config"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/logger"
	"github.com/jasonlvhit/gocron"
)

var log *logger.LoggerWrapper

func init() {
	s := gocron.NewScheduler()
	log = logger.InitLogger("api_gateway", s)
}

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config, log)
	server.Start()
}
