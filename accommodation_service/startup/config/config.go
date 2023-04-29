package config

import "os"

type Config struct {
	Port		string
	AccommodationDBHost	string
	AccommodationDBPort	string
	// NatsHost	string
	// NatsPort	string
	// NatsUser	string
	// NatsPass	string
	// CreateUserCommandSubject string
	// CreateUserReplySubject   string
}

func NewConfig() *Config{
	return &Config{
		Port:                      os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDBHost:           os.Getenv("ACCOMMODATION_DB_HOST"),
		AccommodationDBPort:           os.Getenv("ACCOMMODATION_DB_PORT"),
		// NatsHost:                  os.Getenv("NATS_HOST"),
		// NatsPort:                  os.Getenv("NATS_PORT"),
		// NatsUser:                  os.Getenv("NATS_USER"),
		// NatsPass:                  os.Getenv("NATS_PASS"),
		// CreateOrderCommandSubject: os.Getenv("CREATE_ORDER_COMMAND_SUBJECT"),
		// CreateOrderReplySubject:   os.Getenv("CREATE_ORDER_REPLY_SUBJECT"),
	}
}