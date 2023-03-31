package service

// Imports
import (
	"flight/model"
	"flight/repository"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlightService struct {
	FlightRepository *repository.FlightRepository
}

func (service *FlightService) Create(flight *model.Flight) error {
	err := service.FlightRepository.Create(flight)

	if err != nil {
		return err
	}

	return nil
}

func (service *FlightService) DeleteFlight(ID uuid.UUID) error {
	err := service.FlightRepository.DeleteFlight(ID)
	if err != nil {
		return err
	}
	return nil
}

func (service *FlightService) GetAllFlights() ([]primitive.M, error) {
	var flights []primitive.M
	flights, err := service.FlightRepository.GetAllFlights()
	if err != nil {
		return nil, err
	}
	return flights, nil
}

func (service *FlightService) GetFlightsBySearchCriteria(departure time.Time, origin, destination string, availableSeats int) ([]primitive.M, error) {
	flights, err := service.FlightRepository.SearchFlights(availableSeats, departure, origin, destination)
	if err != nil {
		return nil, err
	}
	return flights, nil
}
