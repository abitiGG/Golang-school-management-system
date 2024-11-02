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

var teacherCollection *mongo.Collection = database.OpenCollection(database.Client, "teacher")

func CreateTeacher() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var course models.Course
		if err := c.BindJSON(&course); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		course.ID = primitive.NewObjectID().Hex()
		course.CreatedAt = time.Now()
		course.UpdatedAt = time.Now()

		_, err := courseCollection.InsertOne(ctx, course)
		if err != nil {
			log.Printf("Error occurred while creating the teacher: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating the teacher"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Teacher created successfully"})
	}
}

func GetTeachers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var teacher models.Teacher
		teacherID := c.Param("teacher_id")

		err := teacherCollection.FindOne(ctx, bson.M{"teacher_id": teacherID}).Decode(&teacher)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while listing the courses"})
			return
		}
		c.JSON(http.StatusOK, teacher)

	}

}

func GetTeacher() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var teacher models.Teacher
		teacherID := c.Param("teacher_id")

		err := teacherCollection.FindOne(ctx, bson.M{"teacher_id": teacherID}).Decode(&teacher)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while fetching the course"})
			return
		}
		c.JSON(http.StatusOK, teacher)
	}
}

func UpdateTeacher() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var teacher models.Teacher
		teacherID := c.Param("teacher_id")

		if err := c.BindJSON(&teacher); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateObj primitive.D
		if teacher.Name != "" {
			updateObj = append(updateObj, bson.E{Key: "name", Value: teacher.Name})
		}
		if teacher.ID != "" {
			updateObj = append(updateObj, bson.E{Key: "id", Value: teacher.ID})
		}
		_, err := teacherCollection.UpdateOne(
			ctx,
			bson.M{"teacher_id": teacherID},
			bson.D{{Key: "$set", Value: updateObj}},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating the course"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Teacher updated successfully"})
	}
}

func DeleteTeacher() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		teacherID := c.Param("teacher_id")

		_, err := teacherCollection.DeleteOne(ctx, bson.M{"teacher_id": teacherID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting the teacher"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted successfully"})
	}
}
