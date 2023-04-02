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
	"github.com/rs/cors"
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

	router.HandleFunc("/signup", handler.UserHandler.Create).Methods("POST")
	router.HandleFunc("/signin", handler.UserHandler.SignIn).Methods("POST")
	router.HandleFunc("/createFlight", handler.FlightHandler.Create).Methods("POST")
	router.HandleFunc("/deleteFlight", handler.FlightHandler.DeleteFlight).Methods("DELETE")
	router.HandleFunc("/showFlights", handler.FlightHandler.GetAllFlights).Methods("GET")
	router.HandleFunc("/searchFlights", handler.FlightHandler.GetFlightsBySearchCriteria).Methods("POST")
	router.HandleFunc("/buyTicket", handler.TicketHandler.BuyTicket).Methods("POST")
	router.HandleFunc("/showTickets/{idUser}", handler.TicketHandler.ShowUserTickets).Methods("GET")
	println("Server starting")

	corsSetup := SetupCors()

	http.Handle("/", corsSetup.Handler(router))
	err := http.ListenAndServe(":8084", corsSetup.Handler(router))
	if err != nil {
		log.Println(err)
	}
	// log.Fatal(http.ListenAndServe(":8084", router))
}

func SetupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
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
