package routes

import (
	"golang-school-management-system/controllers"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/students", controllers.CreateStudent())
	incomingRoutes.GET("/students", controllers.GetStudents())
	incomingRoutes.GET("/students/:student_id", controllers.GetStudent())
	incomingRoutes.PUT("/students/:student_id", controllers.UpdateStudent())
	incomingRoutes.DELETE("/students/:student_id", controllers.DeleteStudent())
}
