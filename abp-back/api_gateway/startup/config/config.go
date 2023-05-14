package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port              string
	UserHost          string
	UserPort          string
	AccommodationHost string
	AccommodationPort string
	PricingHost       string
	PricingPort       string
	AvailabilityHost  string
	AvailabilityPort  string
}

func NewConfig() *Config {
	fmt.Print("api_gateway -> NewConfig.")
	return &Config{
		Port:              os.Getenv("GATEWAY_PORT"),
		UserHost:          os.Getenv("USER_SERVICE_HOST"),
		UserPort:          os.Getenv("USER_SERVICE_PORT"),
		AccommodationHost: os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort: os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		PricingHost:       os.Getenv("PRICING_SERVICE_HOST"),
		PricingPort:       os.Getenv("PRICING_SERVICE_PORT"),
		AvailabilityHost:  os.Getenv("AVAILABILITY_SERVICE_HOST"),
		AvailabilityPort:  os.Getenv("AVAILABILITY_SERVICE_PORT"),
	}
}
