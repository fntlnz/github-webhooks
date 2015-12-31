package server

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
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
	branch := vars["branch"]

	repoName := fmt.Sprintf("%s/%s", vendor, repository)

	if len(branch) > 0 {
		repoName = fmt.Sprintf("%s/%s", repoName, branch)
	}

	config := c.Configuration

	// Retrieve repository data
	if _, ok := config.Repositories[repoName]; !ok {
		logrus.Warningf("Requested repository is not configured: %s", repoName)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	repo := config.Repositories[repoName]

	// Retrieve event data
	event := r.Header.Get("X-GitHub-Event")

	if len(event) <= 0 {
		logrus.Warning("X-GitHub-Event header not provided")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := repo.Events[event]; !ok {
		logrus.Warningf("Requested event not found for this hook: %s", event)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// TODO(fntlnz): do something with actions (in a service possibly)
	for _, cmdString := range repo.Events[event] {
		logrus.Printf(cmdString)
	}

	w.WriteHeader(http.StatusOK)
}
