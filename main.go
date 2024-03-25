package main

import (
	"Golang-crud/helper"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main()  {
	fmt.Printf("startt")
	routes := httprouter.New()
	routes.GET("/",func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w,"HEEEEEEEE")
	})

	server := http.Server{Addr: "localhost:8888", Handler: routes}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}	