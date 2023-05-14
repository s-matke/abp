package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/s-matke/abp/abp-back/accommodation_service/application"
	"github.com/s-matke/abp/abp-back/accommodation_service/domain"
	"github.com/s-matke/abp/abp-back/accommodation_service/infrastructure/api"
	"github.com/s-matke/abp/abp-back/accommodation_service/infrastructure/persistence"
	"github.com/s-matke/abp/abp-back/accommodation_service/startup/config"
	accommodation "github.com/s-matke/abp/abp-back/common/proto/accommodation_service"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	accommodationStore := server.initAccommodationStore(mongoClient)

	pricingServiceAddr := fmt.Sprintf("%s:%s", server.config.PricingServiceHost, server.config.PricingServicePort)
	availabilityServiceAddr := fmt.Sprintf("%s:%s", server.config.AvailabilityServiceHost, server.config.AvailabilityServicePort)

	accommodationServie := server.initAccommodationService(accommodationStore, pricingServiceAddr, availabilityServiceAddr)
	accommodationHandler := server.initAccommodationHandler(accommodationServie)

	server.startGrpcServer(accommodationHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AccommodationDBHost, server.config.AccommodationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initAccommodationStore(client *mongo.Client) domain.AccommodationStore {
	store := persistence.NewAccommodationMongoDBStore(client)
	store.DeleteAll()
	for _, Accommodation := range accommodations {
		err := store.Insert(Accommodation)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initAccommodationService(store domain.AccommodationStore, pricingServiceAddr string, availabilityServiceAddr string) *application.AccommodationService {
	return application.NewAccommodationService(store, pricingServiceAddr, availabilityServiceAddr)
}

// func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscribe

func (server *Server) initAccommodationHandler(service *application.AccommodationService) *api.AccommodationHandler {
	return api.NewAccommodationHandler(service)
}

func (server *Server) startGrpcServer(accommodationHandler *api.AccommodationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	accommodation.RegisterAccommodationServiceServer(grpcServer, accommodationHandler)
	fmt.Println("Serving...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
