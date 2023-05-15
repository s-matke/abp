package application

import (
	"errors"
	"math"
	"time"

	"github.com/s-matke/abp/abp-back/pricing_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/ru"
	"github.com/rickar/cal/v2/us"
)

type PricingService struct {
	store domain.PricingStore
}

func NewPricingService(store domain.PricingStore) *PricingService {
	return &PricingService{
		store: store,
	}
}

func (service *PricingService) Get(id primitive.ObjectID) (*domain.Pricing, error) {
	return service.store.Get(id)
}

func (service *PricingService) GetAll() ([]*domain.Pricing, error) {
	return service.store.GetAll()
}

func (service *PricingService) GetByAccommodation(id string) (*domain.Pricing, error) {
	return service.store.GetByAccommodation(id)
}

func (service *PricingService) CreatePricing(pricing *domain.Pricing) error {
	return service.store.Insert(pricing)
}

func (service *PricingService) CalculatePrice(id string, numPeople int, startDate, endDate time.Time) (float32, error) {
	pricing, err := service.GetByAccommodation(id)

	if err != nil {
		return -1, err
	}

	holidayCalendar := cal.NewBusinessCalendar()

	holidayCalendar.AddHoliday(
		us.ChristmasDay,
		us.NewYear,
		ru.WomensDay,
		ru.LabourDay,
		ru.OrthodoxChristmas,
	)

	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, endDate.Location())

	var totalPrice float32 = 0.0
	var price float32 = 0.0

	for date := startDate; date.Before(endDate) || date.Equal(endDate); date = date.AddDate(0, 0, 1) {
		price = 0.0
		if IsWeekend(date) {
			price = pricing.Price * pricing.Week.WeekendMultiplier
		} else {
			price = pricing.Price * pricing.Week.WorkdayMultiplier
		}

		price = CalculateSeason(price, date, *pricing)

		if price == -1 {
			return -1, errors.New("something went wrong with CalculateSeason")
		}

		holiday, _, _ := holidayCalendar.IsHoliday(date)

		if holiday {
			price = price * pricing.Holiday
		}
		totalPrice += price
	}

	if pricing.PricingType == domain.PER_PERSON {
		totalPrice *= float32(numPeople)
	}

	totalPrice = float32(math.Round(float64(totalPrice)*100) / 100)

	return totalPrice, nil
}

func IsWeekend(date time.Time) bool {
	if date.Weekday() >= time.Monday && date.Weekday() <= time.Friday {
		return false
	}
	return true
}

func CalculateSeason(price float32, date time.Time, pricing domain.Pricing) float32 {
	yearDay := date.YearDay()

	if yearDay < 80 || yearDay >= 355 {
		return price * pricing.Season.WinterMultiplier
	} else if yearDay >= 80 && yearDay < 172 {
		return price * pricing.Season.SpringMultiplier
	} else if yearDay >= 172 && yearDay < 266 {
		return price * pricing.Season.SummerMultiplier
	} else if yearDay >= 266 && yearDay < 355 {
		return price * pricing.Season.FallMultiplier
	}
	return -1
}
