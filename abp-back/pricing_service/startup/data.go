package startup

import (
	"github.com/s-matke/abp/abp-back/pricing_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var pricings = []*domain.Pricing{
	{
		Id:              getObjectId("1"),
		AccommodationId: "1",
		Price:           250.25,
		Season: domain.Season{
			SpringMultiplier: 1.0,
			SummerMultiplier: 1.1,
			FallMultiplier:   1.3,
			WinterMultiplier: 1.4,
		},
		Week: domain.Week{
			WorkdayMultiplier: 1.0,
			WeekendMultiplier: 1.1,
		},
		Holiday:     1.2,
		PricingType: domain.PER_HOUSEHOLD,
	},
	{
		Id:              getObjectId("2"),
		AccommodationId: "2",
		Price:           150.25,
		Season: domain.Season{
			SpringMultiplier: 1.5,
			SummerMultiplier: 1.15,
			FallMultiplier:   1.25,
			WinterMultiplier: 1.15,
		},
		Week: domain.Week{
			WorkdayMultiplier: 1.25,
			WeekendMultiplier: 1.25,
		},
		Holiday:     1.3,
		PricingType: domain.PER_PERSON,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
