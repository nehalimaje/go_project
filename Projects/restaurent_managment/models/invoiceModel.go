package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID               primitive.ObjectID `bson:"_id"`
	Invoice_id       string             `json:"invoice_id,omitempty" bson:"invoice_id,omitempty" validate:"required,min=2,max=100"`
	Order_id         int                `json:"order_id,omitempty" bson:"order_id,omitempty" validate:"required"`
	Payment_method   string             `json:"payment_method,omitempty" bson:"payment_method,omitempty" validate:"eq=CARD|eq=CASH|eq="`
	Payment_status   string             `json:"payment_status,omitempty" bson:"payment_status,omitempty" validate:"required,eq=PENDING|eq=PAID"`
	Payment_due_date time.Time          `json:"payment_due_date,omitempty" bson:"payment_due_date,omitempty"`
	Created_at       time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"udpated_at" validate:"required"`
}
