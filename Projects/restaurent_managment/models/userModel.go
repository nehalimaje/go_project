package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    string             `json:"first_name,omitempty" bson:"first_name,omitempty" validate:"required,min=2,max=100"`
	Last_name     string             `json:"last_name,omitempty" bson:"last_name,omitempty" validate:"required,min=2,max=100"`
	Password      string             `json:"password" validate:"required,min=6"`
	Email         string             `json:"email" validate:"email,required"`
	Avatar        string             `json:"avatar"`
	Phone         string             `json:"phone,omitempty" bson:"phone,omitempty" validate:"required"`
	Token         string             `json:"token,omitempty" bson:"token,omitempty"`
	Refresh_token string             `json:"refresh_token,omitempty" bson:"refresh_token,omitempty" validate:"required"`
	Created_at    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty" validate:"required"`
	Updated_at    time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty" validate:"required"`
	User_id       string             `json:"user_id,omitempty" bson:"user_id,omitempty" validate:"required"`
}
