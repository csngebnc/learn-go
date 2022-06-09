package models

import (
	"github.com/jinzhu/gorm"
)

type Grade struct {
	gorm.Model
	Value int `json:"value"`
	StudentID int64
}

type GradeDto struct {
	Value int
	StudentID int64
}