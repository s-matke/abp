package handler

// Imports
import (
	"encoding/json"
	"flight/model"
	"flight/service"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo"
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

func (handler *FlightHandler) GetFlightsBySearchCriteria(writer http.ResponseWriter, req *http.Request) {
	var requestBody struct {
		Departure      time.Time `bson:"departure" json:"departure"`
		Origin         string    `bson:"origin" json:"origin"`
		Destination    string    `bson:"destination" json:"destination"`
		AvailableSeats int       `bson:"availableSeats" json:"availableSeats"`
	}
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	flights, err := handler.FlightService.GetFlightsBySearchCriteria(requestBody.Departure, requestBody.Origin, requestBody.Destination, requestBody.AvailableSeats)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(&flights)
}

func (handler *FlightHandler) DeleteFlight(writer http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	err := handler.FlightService.DeleteFlight(objId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode("flight deleted successfully")

}
