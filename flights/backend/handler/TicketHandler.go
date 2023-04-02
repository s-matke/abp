package handler

import (
	"encoding/json"
	"flight/model"
	"flight/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
		fmt.Println("Greska")
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")

}

func (handler *TicketHandler) ShowUserTickets(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["idUser"]
	ticketPresenet, err := handler.TicketService.ShowUserTickets(key)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(ticketPresenet)

}
