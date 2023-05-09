package main

import (
	"github.com/mihajlo-ra92/XML/booking_service/startup"
	cfg "github.com/mihajlo-ra92/XML/booking_service/startup/config"
)
func main(){
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}