package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/csngebnc/schoolapp/pkg/routes"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	router := mux.NewRouter()

	router.Use(func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.URL.Path)
				defer func() {
					if err := recover(); err != nil {
						switch v := err.(type){
							case validator.ValidationErrors:
								errors := make(map[string][]string)

								for _, fe := range err.(validator.ValidationErrors) {
										errors[fe.Field()] = append(errors[fe.Field()], fe.ActualTag())
								}
								
								res, _ := json.Marshal(errors)
								w.WriteHeader(http.StatusBadRequest)
								w.Write(res)

								return
							default:
								fmt.Printf("Unexpected type: %v", v)
						}
					}
				}()
        next.ServeHTTP(w, r)
    })
	})

	routes.RegisterStudentRoutes(router)
	routes.RegisterGradeRoutes(router)

	http.Handle("/", router)
	fmt.Println("Listening on: 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}