package main

import (
	//"encoding/json"

	"fmt"
	"net/http"

	"github.com/csngebnc/schoolapp/pkg/models"
	"github.com/csngebnc/schoolapp/pkg/routes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/samber/lo"
	//"github.com/gorilla/mux"
	//	"github.com/samber/lo"
	//"gopkg.in/go-playground/validator.v9"
)

func main() {

	students := []models.Student{
		{
			Name: "Bob",
			Age: 9,
		},
		{
			Name: "Alice",
			Age: 11,
		},
		{
			Name: "Bob",
			Age: 10,
		},
	}

	filtered := models.FilterStudentsUniqueByName(students)
	fmt.Println(filtered)

	router := gin.New()
	router.Use(gin.Logger(), func(c *gin.Context){
		defer func() {
			if err := recover(); err != nil {
				switch v := err.(type){
					case validator.ValidationErrors:
						errors := make(map[string]string)

						lo.ForEach(err.(validator.ValidationErrors), func(fe validator.FieldError, _ int){
							errors[fe.Field()] = fe.ActualTag()
						})

						c.JSON(http.StatusBadRequest, errors)
						return
					default:
						fmt.Printf("Unexpected type: %v", v)
				}
			}
		}()
	})

	routes.RegisterStudentRoutes(router)
	routes.RegisterGradeRoutes(router)

	router.Run(":8000")
}