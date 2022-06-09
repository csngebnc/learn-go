package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/csngebnc/schoolapp/pkg/models"
	"github.com/csngebnc/schoolapp/pkg/utils"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

var NewStudent models.Student

func GetStudents(w http.ResponseWriter, r *http.Request){
	students := models.GetAllStudents()

	res, _ := json.Marshal(students)

	w.Write(res)
}

func GetStudentById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	studentId := vars["studentId"]

	parsedId, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("hiba történt")
	}

	student, _ := models.GetStudentById(parsedId)

	res, _ := json.Marshal(student)
	w.Write(res)
}

func CreateStudent(w http.ResponseWriter, r *http.Request){
	student := &models.Student{}
	utils.ParseBody(r, student)

	v := validator.New()
	err := v.Struct(student)
	if err != nil {
		panic(err)
	}

	insertResult := student.CreateStudent()

	res, _ := json.Marshal(insertResult)
	w.Write(res)
}

func AddGrade(w http.ResponseWriter, r *http.Request){
	gradeDto := &models.GradeDto{}
	utils.ParseBody(r, gradeDto)

	student, _ := models.GetStudentById(gradeDto.StudentID)

	res, _ := json.Marshal(student.AddGrade(gradeDto.Value))
	w.Write(res)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	studentId := vars["studentId"]

	parsedId, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("hiba történt")
	}

	student := models.DeleteStudent(parsedId)
	res, _ := json.Marshal(student)
	w.Write(res)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request){
	updateData := &models.Student{}
	utils.ParseBody(r, updateData)

	vars := mux.Vars(r)
	studentId := vars["studentId"]

	parsedId, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("hiba történt")
	}

	student, db := models.GetStudentById(parsedId)
	if(updateData.Name != ""){
		student.Name = updateData.Name
	}
	if(updateData.Age > 0){
		student.Age = updateData.Age
	}

	db.Save(&student)
	res, _ := json.Marshal(student)

	w.Write(res)
}