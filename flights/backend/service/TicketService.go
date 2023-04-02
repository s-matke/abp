package service

import (
	"flight/dto"
	"flight/model"
	"flight/repository"
)

type TicketService struct {
	TicketRepository *repository.TicketRepository
}

func (service *TicketService) BuyTicket(ticket *model.DTOTicket) error {

	flight := ticket.Flight
	if (flight.AvailableSeats-ticket.NumberOfTickets) < 0 || ticket.NumberOfTickets <= 0 {
		return nil
	}
	ticket.Flight.AvailableSeats = ticket.Flight.AvailableSeats - ticket.NumberOfTickets
	err := service.TicketRepository.BuyTicket(ticket)

	if err != nil {
		return err
	}

	return nil
}

func (service *TicketService) ShowUserTickets(idUser string) ([]dto.TicketPresenetDTO, error) {

	var ticketPresenet []dto.TicketPresenetDTO
	ticketPresenet, err := service.TicketRepository.ShowUserTickets(idUser)
	if err != nil {
		return nil, nil
	}
	return ticketPresenet, nil

}
