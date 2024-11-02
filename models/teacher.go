package models

type Teacher struct {
	ID      string `bson:"_id,omitempty"`
	Name    string `bson:"name"`
	Subject string `bson:"subject"`
}
