package app

import (
	"fmt"
	"getir-assignment/src/di"
	"getir-assignment/src/route"
	"net/http"
	"os"
)

type application struct {
	context *di.Context
}

func NewApplication() *application {
	ctx := di.InitContext()

	return &application{
		context: ctx,
	}
}

func (a *application) Start() {

	router := route.ApiRoute(a.context)

	fmt.Println("server is running")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}

	panic(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
