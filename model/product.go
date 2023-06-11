package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID  `bson:"_id"`
	Name        *string             `bson:"name"`
	Chapter     *uint16             `bson:"chapter"`
	Price       *float64            `bson:"price"`
	Description *string             `bson:"description"`
	Author 		*string             `bson:"author"`
	Genre 	 	*string             `bson:"genre"`
	Publication *string             `bson:"publication"`
	Photo       *string             `bson:"photo"`
	Vendor      *primitive.ObjectID `bson:"vendor"`
}

type RegisterUpdateProduct struct {
	Name        string             `json:"name" validate:"required,min=2,max=50,nameManga"`
	Chapter     uint16             `json:"chapter" validate:"required,numeric"`
	Price       float64            `json:"price" validate:"required,numeric"`
	Description string             `json:"description" validate:"required,min=8,max=100,description"`
	Author 		string             `json:"author" validate:"required,min=2,max=50,general"`
	Genre 	 	string             `json:"genre" validate:"required,min=2,max=50,general"`
	Publication string             `json:"publication" validate:"required,min=2,max=50,general"`
	Photo       string             `json:"image" validate:"required,url"`

}

type ProductView struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Chapter     uint16             `json:"chapter" bson:"chapter"`
	Price       float64            `json:"price" bson:"price"`
	Description string             `json:"description" bson:"description"`
	Author 		string             `json:"author" bson:"author"`
	Genre 	 	string             `json:"genre" bson:"genre"`
	Publication string             `json:"publication" bson:"publication"`
	Photo       string             `json:"image" bson:"photo"`
	Vendor      primitive.ObjectID `json:"idVendor" bson:"vendor"`

}

type Sold struct {
	Products []ProductView `json:"products"`
}
