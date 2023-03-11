package handler

// Imports
import (
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
