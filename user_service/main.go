package main

import (
	"mihajlo-ra92/microservices_demo/user_service/startup"
	cfg "mihajlo-ra92/microservices_demo/user_service/startup/config"
)

func main(){
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
