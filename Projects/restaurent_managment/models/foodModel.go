package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=2,max=100"`
	Price      int                `json:"price,omitempty" bson:"price,omitempty" validate:"required"`
	Food_image string             `json:"food_image,omitempty" bson:"food_image,omitempty" validate:"required"`
	CreatedAt  time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty" validate:"required"`
	UpdatedAt  time.Time          `json:"updatedat,omitempty" bson:"updatedat,omitempty"`
	Food_id    string             `json:"food_id"`
	Menu_id    string             `json:"menu_id" validate:"required"`
}
