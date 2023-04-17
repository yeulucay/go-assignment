package route

import (
	"getir-assignment/src/di"
)

func ApiRoute(ctx *di.Context) *router {
	r := &router{}

	r.Get("in-memory", ctx.InMemoryHandler.Get)
	r.Post("in-memory", ctx.InMemoryHandler.Post)

	return r
}
