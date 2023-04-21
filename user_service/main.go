package main

import (
	cfg "github.com/mihajlo-ra92/XML/user_service/startup/config"

	"github.com/mihajlo-ra92/XML/user_service/startup"
)

func main(){
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
