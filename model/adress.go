package model

type Address struct {
	Country      string `bson:"country"`
	Street       string `bson:"street"`
	Number       string `bson:"number"`
	City         string `bson:"city"`
	Neighborhood string `bson:"neighborhood"`
}

type RegisterAddress struct {
	Country      string `json:"country"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
}
