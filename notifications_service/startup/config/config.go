package config

import "os"

type Config struct {
	Port                string
	NotificationsDBHost string
	NotificationsDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:                os.Getenv("NOTIFICATIONS_SERVICE_PORT"),
		NotificationsDBHost: os.Getenv("NOTIFICATIONS_DB_HOST"),
		NotificationsDBPort: os.Getenv("NOTIFICATIONS_DB_PORT"),
	}
}
