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

var courseCollection *mongo.Collection = database.OpenCollection(database.Client, "courses")

func CreateCourse() gin.HandlerFunc {
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
			log.Printf("Error occurred while creating the course: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating the course"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Course created successfully"})
	}
}

func GetCourses() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var course models.Course
		CourseID := c.Param("course_id")

		err := courseCollection.FindOne(ctx, bson.M{"course_id": CourseID}).Decode(&course)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while listing the courses"})
			return
		}
		c.JSON(http.StatusOK, course)

	}

}

func GetCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var course models.Course
		CourseID := c.Param("course_id")

		err := courseCollection.FindOne(ctx, bson.M{"course_id": CourseID}).Decode(&course)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while fetching the course"})
			return
		}
		c.JSON(http.StatusOK, course)
	}
}

func UpdateCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var course models.Course
		courseID := c.Param("course_id")

		if err := c.BindJSON(&course); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateObj primitive.D
		if course.CourseName != "" {
			updateObj = append(updateObj, bson.E{Key: "course_name", Value: course.CourseName})
		}
		if course.ID != "" {
			updateObj = append(updateObj, bson.E{Key: "course_id", Value: course.ID})
		}

		// Use courseID to update the specific course
		_, err := courseCollection.UpdateOne(
			ctx,
			bson.M{"course_id": courseID},
			bson.D{{Key: "$set", Value: updateObj}},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating the course"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
	}
}

func DeleteCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		courseID := c.Param("course_id")

		_, err := courseCollection.DeleteOne(ctx, bson.M{"course_id": courseID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting the course"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
	}
}
