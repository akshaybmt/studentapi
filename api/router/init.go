package router

import (
	"proj1/api/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	v2 := r.Group("/api/v2")

	v1.POST("/student", handlers.CreateStudent)
	v1.GET("/student", handlers.GetStudents)
	v1.GET("/student/:id", handlers.GetStudentByID)
	v1.PUT("/student/:id", handlers.UpdateStudent)
	v1.DELETE("/student/:id", handlers.DeleteStudent)

	v2.POST("/login", handlers.Login)
	v2.POST("/signup", handlers.Signup)

	return r
}
