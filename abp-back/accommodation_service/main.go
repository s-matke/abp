package main

import (
	"github.com/s-matke/abp/abp-back/accommodation_service/startup"
	cfg "github.com/s-matke/abp/abp-back/accommodation_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()

}
