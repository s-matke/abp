package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/s-matke/abp/abp-back/availability_service/application"
	"github.com/s-matke/abp/abp-back/availability_service/domain"
	"github.com/s-matke/abp/abp-back/availability_service/infrastructure/api"
	"github.com/s-matke/abp/abp-back/availability_service/infrastructure/persistence"
	"github.com/s-matke/abp/abp-back/availability_service/startup/config"
	availability "github.com/s-matke/abp/abp-back/common/proto/availability_service"
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
	availabilityStore := server.initAvailabilityStore(mongoClient)
	availabilityService := server.initAvailabilityService(availabilityStore)
	availabilityHandler := server.initAvailabilityHandler(availabilityService)

	server.startGrpcServer(availabilityHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AvailabilityDBHost, server.config.AvailabilityDBPort)

	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initAvailabilityStore(client *mongo.Client) domain.AvailabilityStore {
	store := persistence.NewAvailabilityMongoDBStore(client)
	store.DeleteAll()
	for _, availability := range availabilities {
		err := store.Insert(availability)
		if err != nil {
			log.Fatal(err)
		}
	}

	return store
}

func (server *Server) initAvailabilityService(store domain.AvailabilityStore) *application.AvailabilityService {
	return application.NewAvailabilityService(store)
}

func (server *Server) initAvailabilityHandler(service *application.AvailabilityService) *api.AvailabilityHandler {
	return api.NewAvailabilityHandler(service)
}

func (server *Server) startGrpcServer(availabilityHandler *api.AvailabilityHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	availability.RegisterAvailabilityServiceServer(grpcServer, availabilityHandler)
	fmt.Println("Serving...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
