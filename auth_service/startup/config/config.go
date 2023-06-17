package config

import "os"

type Config struct {
	Port              string
	UserPort          string
	UserHost          string
	AccommodationPort string
	AccommodationHost string
	BookingPort       string
	BookingHost       string
	RatingPort        string
	RatingHost        string
	NotificationsHost string
	NotificationsPort string
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("AUTH_SERVICE_PORT"),
		UserPort:          os.Getenv("USER_SERVICE_PORT"),
		UserHost:          os.Getenv("USER_SERVICE_HOST"),
		AccommodationPort: os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationHost: os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		BookingPort:       os.Getenv("BOOKING_SERVICE_PORT"),
		BookingHost:       os.Getenv("BOOKING_SERVICE_HOST"),
		RatingPort:        os.Getenv("RATING_SERVICE_PORT"),
		RatingHost:        os.Getenv("RATING_SERVICE_HOST"),
		NotificationsHost: os.Getenv("NOTIFICATIONS_SERVICE_HOST"),
		NotificationsPort: os.Getenv("NOTIFICATIONS_SERVICE_PORT"),
	}
}
