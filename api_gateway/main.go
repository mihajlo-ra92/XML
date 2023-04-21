package main

import (
	"github.com/mihajlo-ra92/XML/api_gateway/startup"
	"github.com/mihajlo-ra92/XML/api_gateway/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
