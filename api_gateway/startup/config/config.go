package config

import "os"

type Config struct {
	Port              string
	CatalogueHost     string
	CataloguePort     string
	OrderingHost      string
	OrderingPort      string
	ShippingHost      string
	ShippingPort      string
	InventoryHost     string
	InventoryPort     string
	UserHost          string
	UserPort          string
	AccommodationHost string
	AccommodationPort string
	AuthHost          string
	AuthPort          string
	BookingHost       string
	BookingPort       string
	RatingHost        string
	RatingPort        string
	NotificationsHost string
	NotificationsPort string
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("GATEWAY_PORT"),
		CatalogueHost:     os.Getenv("CATALOGUE_SERVICE_HOST"),
		CataloguePort:     os.Getenv("CATALOGUE_SERVICE_PORT"),
		OrderingHost:      os.Getenv("ORDERING_SERVICE_HOST"),
		OrderingPort:      os.Getenv("ORDERING_SERVICE_PORT"),
		ShippingHost:      os.Getenv("SHIPPING_SERVICE_HOST"),
		ShippingPort:      os.Getenv("SHIPPING_SERVICE_PORT"),
		InventoryHost:     os.Getenv("INVENTORY_SERVICE_HOST"),
		InventoryPort:     os.Getenv("INVENTORY_SERVICE_PORT"),
		UserHost:          os.Getenv("USER_SERVICE_HOST"),
		UserPort:          os.Getenv("USER_SERVICE_PORT"),
		AccommodationHost: os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort: os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AuthHost:          os.Getenv("AUTH_SERVICE_HOST"),
		AuthPort:          os.Getenv("AUTH_SERVICE_PORT"),
		BookingHost:       os.Getenv("BOOKING_SERVICE_HOST"),
		BookingPort:       os.Getenv("BOOKING_SERVICE_PORT"),
		RatingHost:        os.Getenv("RATING_SERVICE_HOST"),
		RatingPort:        os.Getenv("RATING_SERVICE_PORT"),
		NotificationsHost: os.Getenv("NOTIFICATIONS_SERVICE_HOST"),
		NotificationsPort: os.Getenv("NOTIFICATIONS_SERVICE_PORT"),
	}
}
