package handler

import (
	"flight/service"
	"fmt"
	"net/http"
)

type Userhandler struct {
	UserService *service.UserService
}

func (handler *Userhandler) Hello(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello!")
}
