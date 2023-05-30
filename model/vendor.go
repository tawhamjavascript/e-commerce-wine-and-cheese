package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vendor struct {
	ID          primitive.ObjectID   `bson:"_id"`
	Name        string               `bson:"name"`
	Rating      float64              `bson:"rating" default:"0.0"`
	Description string               `bson:"description"`
	Email       string               `bson:"email"`
	Password    []byte               `bson:"password"`
	Products    []primitive.ObjectID `bson:"products" default:"[]"`
	Sold        []Product            `bson:"sold" default:"[]"`
}

type RegisterVendor struct {
	Name            string `json:"name" validator:"required,min=8,max=50,name"`
	Description     string `json:"description" validator:"required,min=8,max=50,description"`
	Email           string `json:"email" validator:"required,email"`
	Password        string `json:"password" validator:"required,min=8,max=50"`
	ConfirmPassword string `json:"confirmPassword" validator:"required,min=8,max=50"`
}

type SignInVendor struct {
	Email    string `json:"Email" validator:"required,email"`
	Password string `json:"Password" validator:"required,min=8,max=50"`
}

type SignInVendoDatabase struct {
	Email    string             `bson:"email"`
	Password []byte             `bson:"password"`
	Id       primitive.ObjectID `bson:"_id"`
}
