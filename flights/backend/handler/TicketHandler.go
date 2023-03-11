package handler

import (
	"flight/service"
	"fmt"
	"net/http"
)

type TicketHandler struct {
	TicketService *service.TicketService
}

func (handler *TicketHandler) Hi(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("Hi!")
}
