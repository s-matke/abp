package startup

import "user_service/domain"

var location_novi_sad = domain.Location{
	City:    "Novi Sad",
	Country: "Serbia",
}

var location_beograd = domain.Location{
	City:    "Beograd",
	Country: "Serbia",
}

var users = []*domain.User{
	{
		Username:    "pera",
		Name:        "Petar",
		Surname:     "Petrovic",
		PhoneNumber: "+3812312313",
		Email:       "pera@gmail.com",
		Password:    "123",
		Role:        domain.GUEST,
		Location:    location_novi_sad,
	},
	{
		Username:    "mirko",
		Name:        "Mirko",
		Surname:     "Mirkovic",
		PhoneNumber: "+371823121",
		Email:       "mirko@gmail.com",
		Password:    "123",
		Role:        domain.GUEST,
		Location:    location_novi_sad,
	},
	{
		Username:    "maja",
		Name:        "Maja",
		Surname:     "Majic",
		PhoneNumber: "+3812313123",
		Email:       "maja@gmail.com",
		Password:    "123",
		Role:        domain.HOST,
		Location:    location_beograd,
	},
	{
		Username:    "dzoni",
		Name:        "Nikola",
		Surname:     "Nikolic",
		PhoneNumber: "+381290903",
		Email:       "dzoni@gmail.com",
		Password:    "123",
		Role:        domain.HOST,
		Location:    location_novi_sad,
	},
}
