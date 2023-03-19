package main

import (
	"context"
	"flight/handler"
	"flight/repository"
	"flight/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDB() *mongo.Database {

	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Uspesno uspostavljena veza sa MongoDB!")

	return client.Database("ain-xml")
}

func startServer(handler *Handler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/hello", handler.UserHandler.Hello).Methods("GET")
	router.HandleFunc("/world", handler.FlightHandler.World).Methods("GET")
	router.HandleFunc("/hi", handler.TicketHandler.Hi).Methods("GET")

	router.HandleFunc("/signup", handler.UserHandler.Create).Methods("POST")
	router.HandleFunc("/signin", handler.UserHandler.SignIn).Methods("POST")

	println("Server starting")
	log.Fatal(http.ListenAndServe(":8084", router))
}

type Handler struct {
	FlightHandler handler.FlightHandler
	UserHandler   handler.UserHandler
	TicketHandler handler.TicketHandler
}

func main() {

	// Inicijalna podesavanja
	database := initDB()

	// Repos
	flightRepo := &repository.FlightRepository{Database: database}
	userRepo := &repository.UserRepository{Database: database}
	ticketRepo := &repository.TicketRepository{Database: database}

	// Services
	flightService := &service.FlightService{FlightRepository: flightRepo}
	userService := &service.UserService{UserRepository: userRepo}
	ticketService := &service.TicketService{TicketRepository: ticketRepo}

	// Handlers
	flightHandler := &handler.FlightHandler{FlightService: flightService}
	userHandler := &handler.UserHandler{UserService: userService}
	ticketHandler := &handler.TicketHandler{TicketService: ticketService}

	handler := new(Handler)
	handler.FlightHandler = *flightHandler
	handler.UserHandler = *userHandler
	handler.TicketHandler = *ticketHandler

	// Pokretanje servera
	startServer(handler)

}
