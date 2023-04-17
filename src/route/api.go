package route

import (
	"getir-assignment/src/di"
)

func ApiRoute(ctx *di.Context) *router {
	r := &router{}

	// In-Memory Routes
	r.Get("in-memory", ctx.InMemoryHandler.Get)
	r.Post("in-memory", ctx.InMemoryHandler.Post)

	// Record Routes
	r.Post("records", ctx.RecordHandler.List)

	return r
}
