package routes

import (
	"github.com/csngebnc/schoolapp/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterGradeRoutes = func(router *gin.Engine){
	router.POST("/grade", controllers.AddGrade)
}