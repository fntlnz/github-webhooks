package server

import(
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type AppHandlerFunc func(c *Context, w http.ResponseWriter, r *http.Request)

func AppHandler(c *Context, hf AppHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hf(c, w, r)
	}
}

func Repository(c *Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vendor := vars["vendor"]
	repository := vars["repository"]
	json.NewEncoder(w).Encode(map[string]string{"vendor": vendor, "repository": repository})

}

