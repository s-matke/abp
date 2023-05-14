package api

import (
	pb "github.com/s-matke/abp/abp-back/common/proto/pricing_service"
	"github.com/s-matke/abp/abp-back/pricing_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapPricing(pricing *domain.Pricing) *pb.Pricing {
	pricingPb := &pb.Pricing{
		Id:              pricing.Id.Hex(),
		AccommodationId: pricing.AccommodationId,
		Price:           pricing.Price,
		Season: &pb.Season{
			SpringMultiplier: pricing.Season.SpringMultiplier,
			SummerMultiplier: pricing.Season.SummerMultiplier,
			FallMultiplier:   pricing.Season.FallMultiplier,
			WinterMultiplier: pricing.Season.WinterMultiplier,
		},
		Week: &pb.Week{
			WorkdayMultiplier: pricing.Week.WorkdayMultiplier,
			WeekendMultiplier: pricing.Week.WeekendMultiplier,
		},
		Holiday:     pricing.Holiday,
		PricingType: mapPricingType(pricing.PricingType),
	}

	return pricingPb
}

func mapPricingType(pricingType domain.PricingType) pb.Pricing_PricingType {
	switch pricingType {
	case domain.PER_PERSON:
		return pb.Pricing_PER_PERSON
	}
	return pb.Pricing_PER_HOUSEHOLD
}

func mapNewPricing(request *pb.CreatePricingRequest) *domain.Pricing {
	pricing := &domain.Pricing{
		Id:              primitive.NewObjectID(),
		AccommodationId: request.Pricing.AccommodationId,
		Price:           request.Pricing.Price,
		Holiday:         request.Pricing.Holiday,
		PricingType:     domain.PricingType(request.Pricing.PricingType),
	}

	season := &domain.Season{
		SpringMultiplier: request.Pricing.Season.SpringMultiplier,
		SummerMultiplier: request.Pricing.Season.SummerMultiplier,
		FallMultiplier:   request.Pricing.Season.FallMultiplier,
		WinterMultiplier: request.Pricing.Season.WinterMultiplier,
	}

	week := &domain.Week{
		WorkdayMultiplier: request.Pricing.Week.WorkdayMultiplier,
		WeekendMultiplier: request.Pricing.Week.WeekendMultiplier,
	}

	pricing.Week = *week
	pricing.Season = *season

	return pricing
}
