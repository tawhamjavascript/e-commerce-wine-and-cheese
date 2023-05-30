package model

import (
	"github.com/google/uuid"
)

type Product struct {
	Uuid        uuid.UUID  `bson:"id"`
	Name        string     `bson:"name"`
	Description string     `bson:"description"`
	Price       string     `bson:"price"`
	Photo       string     `bson:"photo"`
	Rating      float64    `bson:"rating"`
	Vendor      uuid.UUID  `bson:"vendor"`
	Comments    []Comments `bson:"comments"`
}

type RegisterProduct struct {
	Uuid        uuid.UUID `json:"uuid" validator:"required"`
	Name        string    `json:"name" validator:"required,min=8,max=50"`
	Description string    `json:"description" validator:"required,min=10,max=100"`
	Price       string    `bson:"price" validator:"required,numeric"`
	Photo       string    `bson:"photo" validator:"required"`
	Rating      float64   `bson:"rating" validator:"required,numeric"`
	Vendor      uuid.UUID `bson:"vendor" validator:"required"`
}
