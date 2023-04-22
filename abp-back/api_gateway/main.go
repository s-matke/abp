package main

import (
	"github.com/s-matke/abp/abp-back/api_gateway/startup"
	"github.com/s-matke/abp/abp-back/api_gateway/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
