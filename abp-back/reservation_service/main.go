package main

import (
	"github.com/s-matke/abp/abp-back/reservation_service/startup"
	cfg "github.com/s-matke/abp/abp-back/reservation_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}