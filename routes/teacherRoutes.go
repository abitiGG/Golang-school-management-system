package routes

import (
	"golang-school-management-system/controllers"

	"github.com/gin-gonic/gin"
)

func TeacherRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/teachers", controllers.CreateTeacher())
	incomingRoutes.GET("/teachers", controllers.GetTeachers())
	incomingRoutes.GET("/teachers/:teacher_id", controllers.GetTeacher())
	incomingRoutes.PUT("/teachers/:teacher_id", controllers.UpdateTeacher())
	incomingRoutes.DELETE("/teachers/:teacher_id", controllers.DeleteTeacher())
}
