package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// @title Person API
// @description This is an OpenAPI specification of Person API.
// @contact.name API Support
// @contact.url https://github.com/kucinghitam13
// @contact.email dika.adhitama@gmail.com
// @host localhost:8080
// @schemes http
func main() {
	router := httprouter.New()

	router.GET("/persons", GetPersons)
	router.POST("/persons", AddPerson)
	router.PUT("/persons/id/:id", EditPerson)
	router.DELETE("/persons/id/:id", DeletePerson)

	initAPISpec(router)
	initSwagger(router)

	fmt.Println("server run on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
	}
	fmt.Println("closed")
}
