package models

import (
	"github.com/csngebnc/schoolapp/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Student struct {
	gorm.Model
	Name string `json:"name" validate:"required"`
	Age uint8 `json:"age" validate:"required"`
	Grades []Grade `gorm:"foreignKey:StudentID" json:"grades"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Student{}, &Grade{})
}

func (s *Student) CreateStudent() *Student{
	db.NewRecord(s)
	db.Create(&s)

	return s
}

func GetAllStudents() []Student{
	var Students []Student
	db.Find(&Students)
	
	return Students
}

func GetStudentById(id int64) (*Student, *gorm.DB){
	var student Student
	db := db.First(&student, id)

	return &student, db
}

func (s *Student) AddGrade(value int) *Student{
	var student Student
	db.First(&student, s.ID)

	student.Grades = append(student.Grades, Grade{Value: value})
	db.Save(&student)

	return &student
}

func DeleteStudent(id int64) Student{
	var student Student
	db.Where("ID=?", id).Delete(&student)

	return student
}