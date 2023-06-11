package model

import (

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name 	  *string             `bson:"name"`
	Password  *[]byte             `bson:"password"`
	Email     *string             `bson:"email"`
	ShoppingCart []ProductView  	  `bson:"shoppingCart"`
	Purchased	[]ProductView		  `bson:"purchased"`


}

type SignUpClient struct {
	Email           string `json:"email" validate:"required,email"`
	Name            string `json:"name" validate:"required,min=3,max=50,general"`
	Password        string `json:"password" validate:"required,min=8,max=50"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8,max=50"`
}

type SignInClient struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

type SignInClientDatabase struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string 			`bson:"email"`
	Password []byte 			`bson:"password"`
}