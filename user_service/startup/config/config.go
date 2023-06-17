package config

import "os"

type Config struct {
	Port              string
	UserDBHost        string
	UserDBPort        string
	BookingPort       string
	BookingHost       string
	AccommodationPort string
	AccommodationHost string
	// NatsHost	string
	// NatsPort	string
	// NatsUser	string
	// NatsPass	string
	// CreateUserCommandSubject string
	// CreateUserReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("USER_SERVICE_PORT"),
		UserDBHost:        os.Getenv("USER_DB_HOST"),
		UserDBPort:        os.Getenv("USER_DB_PORT"),
		BookingPort:       os.Getenv("BOOKING_SERVICE_PORT"),
		BookingHost:       os.Getenv("BOOKING_SERVICE_HOST"),
		AccommodationPort: os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationHost: os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		// NatsHost:                  os.Getenv("NATS_HOST"),
		// NatsPort:                  os.Getenv("NATS_PORT"),
		// NatsUser:                  os.Getenv("NATS_USER"),
		// NatsPass:                  os.Getenv("NATS_PASS"),
		// CreateOrderCommandSubject: os.Getenv("CREATE_ORDER_COMMAND_SUBJECT"),
		// CreateOrderReplySubject:   os.Getenv("CREATE_ORDER_REPLY_SUBJECT"),
	}
}
