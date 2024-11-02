package routes

import (
	"golang-school-management-system/controllers"

	"github.com/gin-gonic/gin"
)

func EnrollmentRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/enrollments", controllers.CreateEnrollment())
	incomingRoutes.GET("/enrollments", controllers.GetEnrollments())
	incomingRoutes.GET("/enrollments/:enrollment_id", controllers.GetEnrollment())
	incomingRoutes.PUT("/enrollments/:enrollment_id", controllers.UpdateEnrollment())
	incomingRoutes.DELETE("/enrollments/:enrollment_id", controllers.DeleteEnrollment())

}
