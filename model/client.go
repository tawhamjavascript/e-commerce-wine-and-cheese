package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	ID        primitive.ObjectID `bson:"_id"`
	Password  string             `bson:"password"`
	Photo     string             `bson:"photo"`
	Phone     string             `bson:"phone"`
	Email     string             `bson:"email"`
}

type SignUpClient struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,max=50"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8,max=50"`
}

type SignInClient struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}
