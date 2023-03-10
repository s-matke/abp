package main

import (
	"context"
	"flight/handler"
	"flight/repository"
	"flight/service"
	"fmt"
	"log"

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

	return client.Database("letovi-database")
}

func startServer(handler *handler.FlightHandler) {
	//router := mux.NewRouter().StrictSlash(true)

	//router.HandleFunc()
	// println("Server starting")
	// log.Fatal(http.ListenAndServe(":8080", router))

}

func main() {

	// Inicijalna podesavanja
	database := initDB()

	flightRepository := &repository.FlightRepository{Database: database}
	service := &service.FlightService{FlightRepository: flightRepository}
	handler := &handler.FlightHandler{FlightService: service}

	// Pokretanje servera
	startServer(handler)

}
