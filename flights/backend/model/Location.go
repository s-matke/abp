package model

type Location struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Address string `json:"address"`
	// StreetName   string `json:"streetName"`
	// StreetNumber string `json:"streetNumber"`
}
