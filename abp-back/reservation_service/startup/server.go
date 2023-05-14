package startup

import (
	"fmt"
	"log"
	"net"

	reservation "github.com/s-matke/abp/abp-back/common/proto/reservation_service"
	"github.com/s-matke/abp/abp-back/reservation_service/application"
	"github.com/s-matke/abp/abp-back/reservation_service/domain"
	"github.com/s-matke/abp/abp-back/reservation_service/infrastructure/api"
	"github.com/s-matke/abp/abp-back/reservation_service/infrastructure/persistence"
	"github.com/s-matke/abp/abp-back/reservation_service/startup/config"
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
	reservationStore := server.initReservationStore(mongoClient)
	reservationService := server.initReservationService(reservationStore)
	reservationHandler := server.initReservationHandler(reservationService)

	server.startGrpcServer(reservationHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.ReservationDBHost, server.config.ReservationDBPort)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func (server *Server) initReservationStore(client *mongo.Client) domain.ReservationStore {
	store := persistence.NewReservationMongoDBStore(client)
	store.DeleteAll()
	for _, reservation := range reservations {
		err := store.Insert(reservation)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initReservationService(store domain.ReservationStore) *application.ReservationService {
	return application.NewReservationService(store)
}

func (server *Server) initReservationHandler(service *application.ReservationService) *api.ReservationHandler {
	return api.NewReservationHandler(service)
}

func (server *Server) startGrpcServer(resevationHandler *api.ReservationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reservation.RegisterReservationServiceServer(grpcServer, resevationHandler)
	fmt.Println("Serving...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
