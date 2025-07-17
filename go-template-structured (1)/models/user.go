package models

import "time"

type User struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	BirthDate time.Time `json:"birth_date" bson:"birth_date"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
