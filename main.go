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

	restHandler := services.MakeHandlers(service)
	restServer := &http.Server{
		Addr:    ":8081",
		Handler: restHandler,
	}
	err := restServer.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
	//gRPCServer := http.Server{}

	//http.ListenAndServe(":8081", handler)
}
