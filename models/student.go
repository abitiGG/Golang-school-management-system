package models

import "time"

type Student struct {
	ID        string    `bson:"_id,omitempty"`
	Name      string    `bson:"name"`
	Age       int       `bson:"age"`
	Grade     string    `bson:"grade"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
