package config

import "os"

type Config struct {
	Port          string
	BookingDBHost string
	BookingDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:          os.Getenv("BOOKING_SERVICE_PORT"),
		BookingDBHost: os.Getenv("BOOKING_DB_HOST"),
		BookingDBPort: os.Getenv("BOOKING_DB_PORT"),
	}
}
