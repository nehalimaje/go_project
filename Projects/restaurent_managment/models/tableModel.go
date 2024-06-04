package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID              primitive.ObjectID `bson:"_id"`
	Number_of_guest int                `json:"number_of_guest,omitempty" bson:"Number_of_guest,omitempty" validate:"required"`
	Table_number    int                `json:"table_number,omitempty" bson:"table_number,omitempty" validate:"required"`
	Created_at      time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"udpated_at"`
	Table_id        string             `json:"table_id"`
}
