package server

import (
	"net/http"

	"github.com/Sirupsen/logrus"
)

type LogResponseWriter struct {
	rw     http.ResponseWriter
	status int
}

func (r *LogResponseWriter) Write(p []byte) (int, error) {
	return r.rw.Write(p)
}

func (r *LogResponseWriter) WriteHeader(status int) {
	r.status = status
	r.rw.WriteHeader(status)
}

func (r *LogResponseWriter) Header() http.Header {
	return r.rw.Header()
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logResponseWriter := &LogResponseWriter{
			rw: w,
		}
		next.ServeHTTP(logResponseWriter, r)
		logrus.Printf(
			"Response: %d (%s)\tRemote Addr: %s\tRequest: %s %s%s",
			logResponseWriter.status,
			http.StatusText(logResponseWriter.status),
			r.RemoteAddr,
			r.Method,
			r.Host,
			r.RequestURI,
		)
	})
}
