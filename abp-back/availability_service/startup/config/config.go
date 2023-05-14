package config

import "os"

type Config struct {
	Port               string
	AvailabilityDBHost string
	AvailabilityDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:               os.Getenv("AVAILABILITY_SERVICE_PORT"),
		AvailabilityDBHost: os.Getenv("AVAILABILITY_DB_HOST"),
		AvailabilityDBPort: os.Getenv("AVAILABILITY_DB_PORT"),
	}
}
