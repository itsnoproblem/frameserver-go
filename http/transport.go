package http

import (
	"github.com/gin-gonic/gin"
)

type Transporter interface {
	MakeRoutes() []*Route
}

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

func NewRoute(method, path string, handler gin.HandlerFunc) *Route {
	return &Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}

func NewRouter(transporters ...Transporter) *gin.Engine {
	router := gin.Default()
	router.HTMLRender = DefaultRenderer

	for _, t := range transporters {
		routes := t.MakeRoutes()

		for _, route := range routes {
			router.Handle(route.Method, route.Path, route.Handler)
		}
	}

	return router
}
