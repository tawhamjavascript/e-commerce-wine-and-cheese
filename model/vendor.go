package model

import (
	"github.com/google/uuid"
)

type Vendor struct {
	Uuid        uuid.UUID   `bson:"uuid"`
	Name        string      `bson:"name"`
	Cnpj        string      `bson:"cnpj"`
	Rating      float64     `bson:"rating"`
	Description string      `bson:"description"`
	Email       string      `bson:"email"`
	Password    []byte      `bson:"password"`
	Products    []uuid.UUID `bson:"products"`
	Address     Address     `bson:"address"`
}

type RegisterVendor struct {
	Uuid            uuid.UUID `json:"uuid" validator:"required"`
	Name            string    `json:"name" validator:"required,min=8,max=20"`
	Rating          float64   `json:"rating" validator:"required,numeric"`
	Description     string    `json:"description" validator:"required,min=8,max=50"`
	Email           string    `json:"email" validator:"required,email"`
	Password        string    `json:"password" validator:"required,min=8,max=50"`
	ConfirmPassword string    `json:"confirmPassword" validator:"required,min=8,max=50"`
	Address         Address   `json:"address"`
}

type SignInVendor struct {
	Email    string `json:"Email" validator:"required,email"`
	Password string `json:"Password" validator:"required,min=8,max=50"`
}
