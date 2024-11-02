package models

import "time"

type Course struct {
	ID         string    `bson:"_id,omitempty"`
	CourseName string    `bson:"course_name"`
	CourseID   string    `bson:"course_id"`
	TeacherID  string    `bson:"teacher_id"`
	CreatedAt  time.Time `bson:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at"`
}
