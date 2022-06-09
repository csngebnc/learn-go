package routes

import (
	"github.com/csngebnc/schoolapp/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterStudentRoutes = func(router *mux.Router){
	router.HandleFunc("/student", controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/student/grade", controllers.AddGrade).Methods("POST")
	router.HandleFunc("/student", controllers.GetStudents).Methods("GET")
	router.HandleFunc("/student/{studentId}", controllers.GetStudentById).Methods("GET")
	router.HandleFunc("/student/{studentId}", controllers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student/{studentId}", controllers.DeleteStudent).Methods("DELETE")
}