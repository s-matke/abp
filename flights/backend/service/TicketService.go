package service

import "flight/repository"

type TicketService struct {
	TicketRepository *repository.TicketRepository
}
