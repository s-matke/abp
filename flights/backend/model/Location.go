package model

type Location struct {
	Country string `bson:"country" json:"country"`
	City    string `bson:"city" json:"city"`
	Address string `bson:"address" json:"address"`
	// StreetName   string `json:"streetName"`
	// StreetNumber string `json:"streetNumber"`
}
