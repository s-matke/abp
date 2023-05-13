package api

import (
	"context"

	pb "github.com/s-matke/abp/abp-back/common/proto/pricing_service"
	"github.com/s-matke/abp/abp-back/pricing_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PricingHandler struct {
	pb.UnimplementedPricingServiceServer
	service *application.PricingService
}

func NewPricingHandler(service *application.PricingService) *PricingHandler {
	return &PricingHandler{
		service: service,
	}
}

func (handler *PricingHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	pricing, err := handler.service.Get(objectId)

	if err != nil {
		return nil, err
	}

	pricingPb := mapPricing(pricing)
	response := &pb.GetResponse{
		Pricing: pricingPb,
	}

	return response, nil
}

func (handler *PricingHandler) GetByAccommodation(ctx context.Context, request *pb.GetByAccommodationRequest) (*pb.GetByAccommodationResponse, error) {
	accommodation_id := request.AccommodationId

	pricing, err := handler.service.GetByAccommodation(accommodation_id)

	if err != nil {
		return nil, err
	}

	pricingPb := mapPricing(pricing)
	response := &pb.GetByAccommodationResponse{
		Pricing: pricingPb,
	}

	return response, nil
}

func (handler *PricingHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	pricings, err := handler.service.GetAll()

	if err != nil {
		return nil, err
	}

	response := &pb.GetAllResponse{
		Pricings: []*pb.Pricing{},
	}

	for _, pricing := range pricings {
		current := mapPricing(pricing)
		response.Pricings = append(response.Pricings, current)
	}

	return response, nil
}
