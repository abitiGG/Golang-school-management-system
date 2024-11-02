package routes

import (
	"golang-school-management-system/controllers"

	"github.com/gin-gonic/gin"
)

func CourseRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/courses", controllers.CreateCourse())
	incomingRoutes.GET("/courses", controllers.GetCourses())
	incomingRoutes.GET("/courses/:course_id", controllers.GetCourse())
	incomingRoutes.PUT("/courses/:course_id", controllers.UpdateCourse())
	incomingRoutes.DELETE("/courses/:course_id", controllers.DeleteCourse())

}
