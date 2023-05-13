package main

import (
	"fmt"

	"github.com/s-matke/abp/abp-back/api_gateway/startup"
	"github.com/s-matke/abp/abp-back/api_gateway/startup/config"
)

func main() {
	fmt.Print("api_gateway -> main")
	config := config.NewConfig()
	fmt.Print("api_gateway -> NewConfig Created.")
	server := startup.NewServer(config)
	fmt.Print("api_gateway -> NewServer Created.")
	fmt.Print("api_gateway -> Starting server...")
	server.Start()
}
