package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID         primitive.ObjectID `bson:"_id"`
	Text       string             `json:"text,omitempty" bson:"text,omitempty"`
	Title      string             `json:"title,omitempty" bson:"title,omitempty"`
	Created_at time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	Updated_at time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Note_id    string             `json:"note_id,omitempty" bson:"note_id,omitempty"`
}
