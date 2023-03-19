package handler

import (
	"encoding/json"
	"flight/model"
	"flight/service"
	"net/http"
)

type TicketHandler struct {
	TicketService *service.TicketService
}

func (handler *TicketHandler) BuyTicket(writer http.ResponseWriter, req *http.Request) {

	var ticket model.DTOTicket
	err := json.NewDecoder(req.Body).Decode(&ticket)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TicketService.BuyTicket(&ticket)

	if err != nil {
		writer.WriteHeader(http.StatusConflict)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")

}
