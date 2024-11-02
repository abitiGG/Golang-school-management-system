package models

import "time"

type Enrollment struct {
	Name           string    `bson:"name"`
	ID             string    `bson:"_id,omitempty"`
	StudentID      string    `bson:"student_id"`
	CourseID       string    `bson:"course_id"`
	Responsibility string    `bson:"responsibility"`
	CreatedAt      time.Time `bson:"created_at"`
	UpdatedAt      time.Time `bson:"updated_at"`
}
