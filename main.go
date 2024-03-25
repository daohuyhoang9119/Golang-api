package main

import (
	"Golang-crud/config"
	"Golang-crud/controller"
	"Golang-crud/helper"
	"Golang-crud/repository"
	"Golang-crud/router"
	"Golang-crud/service"
	"fmt"
	"net/http"
)

func main()  {
	fmt.Printf("startt")
	db := config.DatabaseConnection()

	// repository
	bookRepository := repository.NewBookRepository(db)

	// service
	bookService := service.NewBookServiceImpl(bookRepository)

	// controller
	bookController := controller.NewBookController(bookService)

	// router
	routes := router.NewRouter(bookController)
	

	server := http.Server{Addr: "localhost:8888", Handler: routes}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}	