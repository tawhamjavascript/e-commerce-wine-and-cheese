package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	ID        primitive.ObjectID `bson:"_id"`
	Password  string             `bson:"password"`
	Photo     string             `bson:"photo"`
	Phone     string             `bson:"phone"`
	Uuid      uuid.UUID          `bson:"uuid"`
	Email     string             `bson:"email"`
	Purchases []Product          `bson:"purchases"`
	Comments  []Comments         `bson:"comments"`
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
