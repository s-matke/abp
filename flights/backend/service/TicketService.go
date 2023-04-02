package service

import (
	"errors"
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
		return errors.New("greska nema mesta")
	}
	print("prosao prvi if")
	ticket.Flight.AvailableSeats = ticket.Flight.AvailableSeats - ticket.NumberOfTickets
	err := service.TicketRepository.BuyTicket(ticket)
	print("prosao buy ticket prvi if")
	if err != nil {
		return errors.New("greska prilikom kupovine")
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
