package model

import "time"

type Employee struct {
	ID         string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string    `json:"name,omitempty" bson:"name,omitempty"`
	Age        int       `json:"age,omitempty" bson:"age,omitempty"`
	Department string    `json:"department,omitempty" bson:"department,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
