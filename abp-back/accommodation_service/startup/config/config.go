package config

import "os"

type Config struct {
	Port                    string
	AccommodationDBHost     string
	AccommodationDBPort     string
	PricingServiceHost      string
	PricingServicePort      string
	AvailabilityServiceHost string
	AvailabilityServicePort string
}

func NewConfig() *Config {
	return &Config{
		Port:                    os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDBHost:     os.Getenv("ACCOMMODATION_DB_HOST"),
		AccommodationDBPort:     os.Getenv("ACCOMMODATION_DB_PORT"),
		PricingServiceHost:      os.Getenv("PRICING_SERVICE_HOST"),
		PricingServicePort:      os.Getenv("PRICING_SERVICE_PORT"),
		AvailabilityServiceHost: os.Getenv("AVAILABILITY_SERVICE_HOST"),
		AvailabilityServicePort: os.Getenv("AVAILABILITY_SERVICE_PORT"),
	}
}
