package main

import (
	"github.com/mihajlo-ra92/XML/auth_service/startup"
	cfg "github.com/mihajlo-ra92/XML/auth_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
