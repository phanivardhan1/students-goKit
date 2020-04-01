package main

import (
	"fmt"
	"net/http"

	"./services"
)

func main() {
	fmt.Println("this is main function")
	fmt.Println("this is test print statement")

	service := services.NewStudentservice()

	handler := services.MakeHandlers(service)

	http.ListenAndServe(":8081", handler)
}
