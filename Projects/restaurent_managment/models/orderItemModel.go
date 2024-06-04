package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID            primitive.ObjectID `bson:"_id"`
	Quantity      string             `json:"quantity,omitempty" bson:"quantity,omitempty" validate:"required"`
	Unit_price    float64            `json:"unit_price,omitempty" bson:"unit_price,omitempty" validate:"required"`
	Created_at    time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"udpated_at"`
	Food_id       string             `json:"food_id"`
	Order_item_id string             `json:"order_item_id,omitempty" bson:"order_item_id,omitempty" validate:"required"`
	Order_id      string             `json:"order_id,omitempty" bson:"order_id,omitempty" validate:"required"`
}
