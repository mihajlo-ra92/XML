package config

import "os"

type Config struct {
	Port              string
	UserPort          string
	UserHost          string
	AccommodationPort string
	AccommodationHost string
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("AUTH_SERVICE_PORT"),
		UserPort:          os.Getenv("USER_SERVICE_PORT"),
		UserHost:          os.Getenv("USER_SERVICE_HOST"),
		AccommodationPort: os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationHost: os.Getenv("ACCOMMODATION_SERVICE_HOST"),
	}
}
