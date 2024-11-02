package main

import (
	middleware "golang-school-management-system/middleware"
	"golang-school-management-system/routes"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.Authentication())

	routes.CourseRoutes(router)
	routes.StudentRoutes(router)
	routes.TeacherRoutes(router)
	routes.EnrollmentRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)
}
