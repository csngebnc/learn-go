package routes

import (
	"github.com/csngebnc/schoolapp/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterGradeRoutes = func(router *mux.Router){
	router.HandleFunc("/grade", controllers.AddGrade).Methods("POST")
}