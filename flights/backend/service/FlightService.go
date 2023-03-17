package service

// Imports
import (
	"flight/model"
	"flight/repository"
)

type FlightService struct {
	FlightRepository *repository.FlightRepository
}

// Funkcije
func (service *FlightService) GetAllFlights() ([]model.Flight, error) {
	flights, err := service.FlightRepository.GetAllFlights()
	if err != nil {
		return nil, err
	}
	return flights, nil
}