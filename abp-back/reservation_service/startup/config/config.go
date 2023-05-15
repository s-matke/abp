package config

import "os"

type Config struct {
	Port              string
	ReservationDBHost string
	ReservationDBPort string
	AccommodationHost string
	AccommodationPort string
	PricingHost       string
	PricingPort       string
	AvailabilityHost  string
	AvailabilityPort  string
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("RESERVATION_SERVICE_PORT"),
		ReservationDBHost: os.Getenv("RESERVATION_DB_HOST"),
		ReservationDBPort: os.Getenv("RESERVATION_DB_PORT"),
		AccommodationHost: os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort: os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		PricingHost:       os.Getenv("PRICING_SERVICE_HOST"),
		PricingPort:       os.Getenv("PRICING_SERVICE_PORT"),
		AvailabilityHost:  os.Getenv("AVAILABILITY_SERVICE_HOST"),
		AvailabilityPort:  os.Getenv("AVAILABILITY_SERVICE_PORT"),
	}
}
