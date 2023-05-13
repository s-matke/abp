package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pricing struct {
	Id              primitive.ObjectID `bson:"_id"`
	AccommodationId string             `bson:"accommodation_id"`
	Price           float32            `bson:"price"`
	Season          Season             `bson:"season"`
	Week            Week               `bson:"week"`
	Holiday         float32            `bson:"holiday"`
	PriceType       PricingType        `bson:"pricing_type"`
}

type PricingType int

const (
	PER_PERSON PricingType = iota
	PER_HOUSEHOLD
)

type Season struct {
	SpringMultiplier float32 `bson:"spring_multiplier"`
	SummerMultiplier float32 `bson:"summer_multiplier"`
	FallMultiplier   float32 `bson:"fall_multiplier"`
	WinterMultiplier float32 `bson:"winter_multiplier"`
}

type Week struct {
	WorkdayMultiplier float32 `bson:"workday_multiplier"`
	WeekendMultiplier float32 `bson:"weekend_multiplier"`
}
