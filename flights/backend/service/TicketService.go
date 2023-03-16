package service

import (
	"flight/model"
	"flight/repository"
)

type TicketService struct {
	TicketRepository *repository.TicketRepository
}

func (service *TicketService) BuyTicket(ticket *model.DTOTicket) error {

	var flight model.Flight
	flight = ticket.Flight
	if (flight.AvailableSeats - ticket.NumberOfTickets) < 0 {
		return nil
	}
	ticket.Flight.AvailableSeats = ticket.Flight.AvailableSeats - ticket.NumberOfTickets
	err := service.TicketRepository.BuyTicket(ticket)

	if err != nil {
		return err
	}

	return nil
}
