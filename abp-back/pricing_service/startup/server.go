package startup

import (
	"fmt"
	"log"
	"net"

	pricing "github.com/s-matke/abp/abp-back/common/proto/pricing_service"
	"github.com/s-matke/abp/abp-back/pricing_service/application"
	"github.com/s-matke/abp/abp-back/pricing_service/domain"
	"github.com/s-matke/abp/abp-back/pricing_service/infrastructure/api"
	"github.com/s-matke/abp/abp-back/pricing_service/infrastructure/persistence"
	"github.com/s-matke/abp/abp-back/pricing_service/startup/config"
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
	pricingStore := server.initPricingStore(mongoClient)
	pricingServie := server.initPricingService(pricingStore)
	pricingHandler := server.initPricingHandler(pricingServie)

	server.startGrpcServer(pricingHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.PricingDBHost, server.config.PricingDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initPricingStore(client *mongo.Client) domain.PricingStore {
	store := persistence.NewPricingMongoDBStore(client)
	store.DeleteAll()
	for _, Pricing := range pricings {
		err := store.Insert(Pricing)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initPricingService(store domain.PricingStore) *application.PricingService {
	return application.NewPricingService(store)
}

// func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscribe

func (server *Server) initPricingHandler(service *application.PricingService) *api.PricingHandler {
	return api.NewPricingHandler(service)
}

func (server *Server) startGrpcServer(pricingHandler *api.PricingHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pricing.RegisterPricingServiceServer(grpcServer, pricingHandler)
	fmt.Println("Serving...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
