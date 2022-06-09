package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	db *gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/go?parseTime=true")

	if err != nil {
		panic("Cound not connect to db")
	}

	db = d
}

func GetDB() *gorm.DB{
	return db
}