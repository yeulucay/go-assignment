package di

import (
	"getir-assignment/src/db"
	"getir-assignment/src/handler"
	"getir-assignment/src/repository"
	"getir-assignment/src/service"
)

type Context struct {
	InMemoryHandler handler.InMemoryHandler
	RecordHandler   handler.RecordHandler
}

func InitContext() *Context {

	//database
	mc := db.InitializeMongo()

	//repository
	mr := repository.NewMongoRepository(mc)

	//services
	ims := service.NewInMemoryService()
	rs := service.NewRecordService(mr)

	//handlers
	imh := handler.NewInMemoryHandler(ims)
	rh := handler.NewRecordHandler(rs)

	return &Context{
		InMemoryHandler: imh,
		RecordHandler:   rh,
	}
}
