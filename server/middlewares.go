package server

import (
	"net/http"
	"github.com/Sirupsen/logrus"
)

func loggingMiddleware(route Route, next http.Handler)  http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			route.Name,
		)
		next.ServeHTTP(w, r)
	})
}
