package handler

// Imports
import (
	"encoding/json"
	"flight/model"
	"flight/service"
	"fmt"
	"net/http"
)

type FlightHandler struct {
	FlightService *service.FlightService
}

// Requestovi/Funkcije
func (handler *FlightHandler) World(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("World!")
}

func (handler *FlightHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var flight model.Flight
	err := json.NewDecoder(req.Body).Decode(&flight)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.FlightService.Create(&flight)

	if err != nil {
		writer.WriteHeader(http.StatusConflict)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *FlightHandler) GetAllFlights(writer http.ResponseWriter, req *http.Request) {
	flights, err := handler.FlightService.GetAllFlights()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(&flights)
}
