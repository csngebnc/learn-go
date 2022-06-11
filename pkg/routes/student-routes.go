package routes

import (
	"github.com/csngebnc/schoolapp/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterStudentRoutes = func(router *gin.Engine){
	router.POST("/student", controllers.CreateStudent)
	router.GET("/student", controllers.GetStudents)
	router.GET("/student/:studentId", controllers.GetStudentById)
	router.PUT("/student/:studentId", controllers.UpdateStudent)
	router.DELETE("/student/:studentId", controllers.DeleteStudent)
}