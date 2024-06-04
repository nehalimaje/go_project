package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Category   string             `json:"category,omitempty" bson:"category,omitempty" validate:"required"`
	Start_date time.Time          `json:"start_date,omitempty" bson:"start_date,omitempty"`
	End_date   time.Time          `json:"end_date,omitempty" bson:"end_date,omitempty"`
	Created_at time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"udpated_at"`
	Menu_id    string             `json:"menu_id" validate:"required"`
}
