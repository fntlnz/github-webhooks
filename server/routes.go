package server

import (
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc AppHandlerFunc
}

var routes = []Route{
	Route{
		Name:        "Repository",
		Method:      "POST",
		Pattern:     "/{vendor}/{repository}",
		HandlerFunc: Repository,
	},
	Route{
		Name:        "Repository",
		Method:      "POST",
		Pattern:     "/{vendor}/{repository}/{branch}",
		HandlerFunc: Repository,
	},
}

func NewRouter(context *Context) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		appHandler := AppHandler(context, route.HandlerFunc)
		r := router.Methods(route.Method)
		r.Path(route.Pattern)
		r.Name(route.Name)
		r.Handler(appHandler)
	}
	return router
}
