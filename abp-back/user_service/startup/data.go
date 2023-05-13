package startup

import (
	"github.com/google/uuid"
	"github.com/s-matke/abp/abp-back/user_service/domain"
)

var location_novi_sad = domain.Location{
	Address: "Mileta Jaksica 2a",
	City:    "Novi Sad",
	Country: "Serbia",
}

var location_beograd = domain.Location{
	Address: "Dunavski kej 47",
	City:    "Beograd",
	Country: "Serbia",
}

var users = []*domain.User{
	{
		Id:          uuid.New(), //uuid.Must(uuid.New(), context.TODO().Err()),
		Username:    "pera",
		Name:        "Petar",
		Surname:     "Petrovic",
		PhoneNumber: "+3812312313",
		Email:       "pera@gmail.com",
		Password:    "123",
		Role:        domain.GUEST,
		Address:     location_novi_sad.Address,
		City:        location_novi_sad.City,
		Country:     location_novi_sad.Country,
	},
	{
		Id:          uuid.New(),
		Username:    "mirko",
		Name:        "Mirko",
		Surname:     "Mirkovic",
		PhoneNumber: "+371823121",
		Email:       "mirko@gmail.com",
		Password:    "123",
		Role:        domain.HOST,
		Address:     location_novi_sad.Address,
		City:        location_novi_sad.City,
		Country:     location_novi_sad.Country,
	},
	{
		Id:          uuid.New(),
		Username:    "maja",
		Name:        "Maja",
		Surname:     "Majic",
		PhoneNumber: "+3812313123",
		Email:       "maja@gmail.com",
		Password:    "123",
		Role:        domain.GUEST,
		Address:     location_beograd.Address,
		City:        location_beograd.City,
		Country:     location_beograd.Country,
	},
	{
		Id:          uuid.New(),
		Username:    "dzoni",
		Name:        "Nikola",
		Surname:     "Nikolic",
		PhoneNumber: "+381290903",
		Email:       "dzoni@gmail.com",
		Password:    "123",
		Role:        domain.HOST,
		Address:     location_beograd.Address,
		City:        location_beograd.City,
		Country:     location_beograd.Country,
	},
}
