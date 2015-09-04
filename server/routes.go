package server

import(
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc AppHandlerFunc
}

var routes = []Route{
	Route {
		Name: "Repository",
		Method:"POST",
		Pattern: "/{vendor}/{repository}",
		HandlerFunc: Repository,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	context := new(Context)
	for _, route := range routes {
		handler := loggingMiddleware(route, AppHandler(context, route.HandlerFunc))
		r := router.Methods(route.Method)
		r.Path(route.Pattern)
		r.Name(route.Name)
		r.Handler(handler)
	}
	return router
}
