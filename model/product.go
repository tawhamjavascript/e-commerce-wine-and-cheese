package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID  `bson:"_id"`
	Name        *string             `bson:"name"`
	Description *string             `bson:"description"`
	Price       *float64            `bson:"price"`
	Photo       *string             `bson:"photo"`
	Rating      float64             `bson:"rating" default:"0.0"`
	Vendor      *primitive.ObjectID `bson:"vendor"`
	Comments    *[]*Comments        `bson:"comments"`
}

type RegisterProduct struct {
	Name        string             `json:"name" validate:"required,min=2,max=50,name"`
	Description string             `json:"description" validate:"required,min=8,max=100,description"`
	Price       float64            `json:"price" validate:"required,numeric"`
	Photo       string             `json:"photo" validate:"required,url"`
}

type UpdateProduct struct {
	Name        string             `json:"name" validate:"required,min=2,max=50,name"`
	Description string             `json:"description" validate:"required,min=8,max=100,description"`
	Price       float64            `json:"price" validate:"required,numeric"`
	Photo       string             `json:"photo" validate:"required,url"`
}
