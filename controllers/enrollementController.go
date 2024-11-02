package controllers

import (
	"context"
	"golang-school-management-system/database"
	"golang-school-management-system/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var enrollmentCollection *mongo.Collection = database.OpenCollection(database.Client, "enrollments")

func CreateEnrollment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var enrollment models.Enrollment
		if err := c.BindJSON(&enrollment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		enrollment.ID = primitive.NewObjectID().Hex()
		enrollment.CreatedAt = time.Now()
		enrollment.UpdatedAt = time.Now()

		_, err := enrollmentCollection.InsertOne(ctx, enrollment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating the enrollment"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Enrollment created successfully"})
	}
}

func GetEnrollments() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, err := enrollmentCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while listing the enrollments"})
			return
		}
		defer cursor.Close(ctx)

		var enrollments []models.Enrollment
		if err = cursor.All(ctx, &enrollments); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while decoding enrollments"})
			return
		}

		c.JSON(http.StatusOK, enrollments)
	}
}

func GetEnrollment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var enrollment models.Enrollment
		enrollmentID := c.Param("enrollment_id")

		err := enrollmentCollection.FindOne(ctx, bson.M{"_id": enrollmentID}).Decode(&enrollment)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Enrollment not found"})
			return
		}
		c.JSON(http.StatusOK, enrollment)
	}
}

func UpdateEnrollment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var enrollment models.Enrollment
		enrollmentID := c.Param("enrollment_id")

		if err := c.BindJSON(&enrollment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateObj primitive.D
		if enrollment.Name != "" {
			updateObj = append(updateObj, bson.E{Key: "name", Value: enrollment.Name})
		}

		_, err := enrollmentCollection.UpdateOne(
			ctx,
			bson.M{"enrollment_id": enrollmentID},
			bson.D{{Key: "$set", Value: updateObj}},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating the enrollment"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Enrollment updated successfully"})
	}
}

func DeleteEnrollment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		enrollmentID := c.Param("enrollment_id")

		_, err := enrollmentCollection.DeleteOne(ctx, bson.M{"_id": enrollmentID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting the enrollment"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Enrollment deleted successfully"})
	}
}
