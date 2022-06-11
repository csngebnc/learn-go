package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/csngebnc/schoolapp/pkg/models"
	"github.com/gin-gonic/gin"
)

var NewStudent models.Student

func GetStudents(c *gin.Context){
	students := models.GetAllStudents()

	c.JSON(http.StatusOK, students)
}

func GetStudentById(c *gin.Context){
	studentId := c.Param("studentId")
	parsedId, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("hiba történt")
	}

	student, _ := models.GetStudentById(parsedId)

	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context){
	var student models.Student
	err := c.Bind(&student)
	
	if err != nil {
		panic(err)
	}

	insertResult := student.CreateStudent()
	c.JSON(http.StatusOK, insertResult)
}

func AddGrade(c *gin.Context){
	var gradeDto models.GradeDto
	c.Bind(&gradeDto)

	student, _ := models.GetStudentById(gradeDto.StudentID)

	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context){
	studentId := c.Param("studentId")
	parsedId, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("hiba történt")
	}

	student := models.DeleteStudent(parsedId)
	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context){
	var updateData models.Student
	c.Bind(&updateData)

	studentId := c.Param("studentId")
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
	c.JSON(http.StatusOK, student)
}