package app

import (
	"fmt"
	"getir-assignment/src/di"
	"getir-assignment/src/route"
	"net/http"
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
	panic(http.ListenAndServe(":8000", router))
}
