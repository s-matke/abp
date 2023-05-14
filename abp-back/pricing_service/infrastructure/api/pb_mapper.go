package api

import (
	pb "github.com/s-matke/abp/abp-back/common/proto/pricing_service"
	"github.com/s-matke/abp/abp-back/pricing_service/domain"
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
