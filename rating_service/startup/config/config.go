package config

import "os"

type Config struct {
	Port              string
	RatingDBHost      string
	RatingDBPort      string
	BookingPort       string
	BookingHost       string
	AccommodationPort string
	AccommodationHost string
	UserPort          string
	UserHost          string
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("RATING_SERVICE_PORT"),
		RatingDBHost:      os.Getenv("RATING_DB_HOST"),
		RatingDBPort:      os.Getenv("RATING_DB_PORT"),
		UserHost:          os.Getenv("USER_SERVICE_HOST"),
		UserPort:          os.Getenv("USER_SERVICE_PORT"),
		BookingPort:       os.Getenv("BOOKING_SERVICE_PORT"),
		BookingHost:       os.Getenv("BOOKING_SERVICE_HOST"),
		AccommodationPort: os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationHost: os.Getenv("ACCOMMODATION_SERVICE_HOST"),
	}
}
