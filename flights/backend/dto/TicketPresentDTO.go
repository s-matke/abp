package dto

import (
	"flight/model"
)

type TicketPresenetDTO struct {
	Origin          model.Location
	Destination     model.Location
	Price           float64
	NumberOfTickets int
	Departure       string
}
