package di

import (
	"getir-assignment/src/handler"
	"getir-assignment/src/service"
)

type Context struct {
	InMemoryHandler handler.InMemoryHandler
}

func InitContext() *Context {

	//services
	ims := service.NewInMemoryService()

	//handlers
	imh := handler.NewInMemoryHandler(ims)

	return &Context{
		InMemoryHandler: imh,
	}
}
