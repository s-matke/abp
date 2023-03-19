package model

type Location struct {
	Country string `bson:"country"`
	City    string `bson:"city"`
	Address string `bson:"address"`
	// StreetName   string `json:"streetName"`
	// StreetNumber string `json:"streetNumber"`
}
