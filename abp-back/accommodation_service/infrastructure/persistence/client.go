package persistence

import (
	"context"
	"fmt"
	"log"

	availability "github.com/s-matke/abp/abp-back/common/proto/availability_service"
	pricing "github.com/s-matke/abp/abp-back/common/proto/pricing_service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetClient(host, port string) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s/", host, port)
	options := options.Client().ApplyURI(uri)
	return mongo.Connect(context.TODO(), options)
}

func NewPricingClient(address string) pricing.PricingServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("failed to start gRPC conn to pricing service: %v", err)
	}

	return pricing.NewPricingServiceClient(conn)
}

func NewAvailabilityClient(address string) availability.AvailabilityServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("failed to start gRPC conn to availability service: %v", err)
	}

	return availability.NewAvailabilityServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
