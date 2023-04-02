package model

type DTOTicket struct {
	IdFlight        string
	Flight          Flight
	IdUser          string
	NumberOfTickets int
}
