package config

import "os"

type Config struct {
	Port		string
	UserDBHost	string
	UserDBPort	string
	NatsHost	string
	NatsPort	string
	NatsUser	string
	NatsPass	string
	// CreateUserCommandSubject string
	// CreateUserReplySubject   string
}

func NewConfig() *Config{
	return &Config{
		Port:                      os.Getenv("INVENTORY_SERVICE_PORT"),
		UserDBHost:           os.Getenv("INVENTORY_DB_HOST"),
		UserDBPort:           os.Getenv("INVENTORY_DB_PORT"),
		// NatsHost:                  os.Getenv("NATS_HOST"),
		// NatsPort:                  os.Getenv("NATS_PORT"),
		// NatsUser:                  os.Getenv("NATS_USER"),
		// NatsPass:                  os.Getenv("NATS_PASS"),
		// CreateOrderCommandSubject: os.Getenv("CREATE_ORDER_COMMAND_SUBJECT"),
		// CreateOrderReplySubject:   os.Getenv("CREATE_ORDER_REPLY_SUBJECT"),
	}
}