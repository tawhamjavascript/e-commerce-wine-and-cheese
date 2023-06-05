package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vendor struct {
	ID          primitive.ObjectID   `bson:"_id"`
	Name        string               `bson:"name"`
	Description string               `bson:"description"`
	Email       string               `bson:"email"`
	Password    []byte               `bson:"password"`
	Products    []primitive.ObjectID `bson:"products" default:"[]"`
	Sold        []primitive.ObjectID `bson:"sold" default:"[]"`
}

type RegisterVendor struct {
	Name            string `json:"name" validate:"required,min=8,max=50,name"`
	Description     string `json:"description" validate:"required,min=8,max=50,description"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,max=50"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8,max=50"`
}

type SignInVendor struct {
	Email    string `json:"Email" validate:"required,email"`
	Password string `json:"Password" validate:"required,min=8,max=50"`
}

type SignInVendoDatabase struct {
	Email    string             `bson:"email"`
	Password []byte             `bson:"password"`
	Id       primitive.ObjectID `bson:"_id"`
}
