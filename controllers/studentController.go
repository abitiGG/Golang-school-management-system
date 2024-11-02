package controllers

import (
	"context"
	"golang-school-management-system/database"
	"golang-school-management-system/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var studentCollection *mongo.Collection = database.OpenCollection(database.Client, "students")

func CreateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var student models.Student
		if err := c.BindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		student.ID = primitive.NewObjectID().Hex()
		student.CreatedAt = time.Now()
		student.UpdatedAt = time.Now()

		_, err := studentCollection.InsertOne(ctx, student)
		if err != nil {
			log.Printf("Error occurred while creating the student: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating the student"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Student created successfully"})
	}
}

func GetStudents() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var students []models.Student
		cursor, err := studentCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while listing the students"})
			return
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var student models.Student
			if err = cursor.Decode(&student); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while decoding student data"})
				return
			}
			students = append(students, student)
		}

		c.JSON(http.StatusOK, students)
	}
}

func GetStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var student models.Student
		studentID := c.Param("student_id")

		err := studentCollection.FindOne(ctx, bson.M{"student_id": studentID}).Decode(&student)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while fetching the student"})
			return
		}
		c.JSON(http.StatusOK, student)
	}
}

func UpdateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var student models.Student
		studentID := c.Param("student_id")

		if err := c.BindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateObj primitive.D
		if student.Name != "" {
			updateObj = append(updateObj, bson.E{Key: "name", Value: student.Name})
		}
		if student.ID != "" {
			updateObj = append(updateObj, bson.E{Key: "id", Value: student.ID})
		}

		_, err := studentCollection.UpdateOne(
			ctx,
			bson.M{"student_id": studentID},
			bson.D{{Key: "$set", Value: updateObj}},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating the student"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Student updated successfully"})
	}
}

func DeleteStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		studentID := c.Param("student_id")

		_, err := studentCollection.DeleteOne(ctx, bson.M{"student_id": studentID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting the student"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
	}
}
