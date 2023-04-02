package service

// Imports
import (
	"flight/model"
	"flight/repository"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlightService struct {
	FlightRepository *repository.FlightRepository
}

func randate() time.Time {
	min := time.Date(2024, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func (service *FlightService) Create(flight *model.Flight) error {

	// flight.Departure = randate()
	// flight.Departure = time.Now()
	// flight.ID = uuid.New()
	err := service.FlightRepository.Create(flight)

	if err != nil {
		return err
	}

	return nil
}

func (service *FlightService) DeleteFlight(id primitive.ObjectID) error {
	err := service.FlightRepository.DeleteFlight(id)
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
