package main

import (
	"github.com/mihajlo-ra92/XML/ordering_service/startup"
	cfg "github.com/mihajlo-ra92/XML/ordering_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
