package handler

// Imports
import (
	"flight/service"
	"fmt"
	"net/http"
	"encoding/json"
	

)

type FlightHandler struct {
	FlightService *service.FlightService
}

// Requestovi/Funkcije
func (handler *FlightHandler) World(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("World!")
}

func (handler *FlightHandler) GetAllFlights(writer http.ResponseWriter, req *http.Request) {
    flights, err := handler.FlightService.GetAllFlights()
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }

    writer.Header().Set("Content-Type", "application/json")
    json.NewEncoder(writer).Encode(flights)
}
