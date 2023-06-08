package config

import "os"

type Config struct {
	Port              string
	BookingDBHost     string
	BookingDBPort     string
	AccommodationHost string
	AccommodationPort string
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("BOOKING_SERVICE_PORT"),
		BookingDBHost:     os.Getenv("BOOKING_DB_HOST"),
		BookingDBPort:     os.Getenv("BOOKING_DB_PORT"),
		AccommodationHost: os.Getenv("ACCOMMODATIOIN_SERVICE_HOST"),
		AccommodationPort: os.Getenv("ACCOMMODATION_SERVICE_PORT"),
	}
}
