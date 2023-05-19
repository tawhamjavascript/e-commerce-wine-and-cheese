package model

import (
	"github.com/google/uuid"
	"time"
)

type Comments struct {
	Uuid        uuid.UUID `bson:"Uuid"`
	DateCreated time.Time `bson:"dateCreated"`
	Title       string    `bson:"title"`
	Photo       string    `bson:"photo"`
	Client      uuid.UUID `bson:"client"`
	Product     uuid.UUID `bson:"product"`
}

type RegisterComments struct {
	Uuid        uuid.UUID `bson:"Uuid"`
	DateCreated time.Time `json:"dateCreated"`
	Title       string    `json:"title"`
	Photo       string    `json:"photo"`
	Client      uuid.UUID `json:"client"`
	Product     uuid.UUID `json:"product"`
}
