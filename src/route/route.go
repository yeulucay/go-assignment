package route

import (
	"getir-assignment/src/handler"
	"net/http"
	"regexp"
)

type route struct {
	pattern string
	method  string
	handler http.Handler
}

type router struct {
	routes []route
}

// to access from out of the package
func InitRouter() *router {
	return &router{}
}

func (r *router) Post(pattern string, handler handler.Handler) {
	rt := route{
		pattern: pattern,
		method:  "POST",
		handler: handler,
	}
	r.addRoute(rt)
}

func (r *router) Get(pattern string, handler handler.Handler) {
	rt := route{
		pattern: pattern,
		method:  "GET",
		handler: handler,
	}
	r.addRoute(rt)
}

//p.s. PUT and DELETE is not required in the project

func (r *router) addRoute(rt route) {
	r.routes = append(r.routes, rt)
}

func (r *router) getHandler(method string, path string) http.Handler {
	for _, rt := range r.routes {
		re := regexp.MustCompile(rt.pattern)
		if rt.method == method && re.MatchString(path) {
			return rt.handler
		}
	}

	return http.NotFoundHandler()
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	handler := r.getHandler(method, path)

	handler.ServeHTTP(w, req)
}
