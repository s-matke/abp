package config

import "os"

type Config struct {
	Port          string
	PricingDBHost string
	PricingDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:          os.Getenv("PRICING_SERVICE_PORT"),
		PricingDBHost: os.Getenv("PRICING_DB_HOST"),
		PricingDBPort: os.Getenv("PRICING_DB_PORT"),
	}
}
