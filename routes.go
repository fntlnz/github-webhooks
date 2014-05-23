package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"os/exec"
	"strings"
)

//Routes ...
func Routes(m *martini.ClassicMartini) {
	m.Post("/:username/:repository/:branch", func(r render.Render, req *http.Request, params martini.Params) {
			repositoryAction(r, req, params)
		})
	m.Post("/:username/:repository", func(r render.Render, req *http.Request, params martini.Params) {
			repositoryAction(r, req, params)
		})
}

func repositoryAction(r render.Render, req *http.Request, params martini.Params) {
	p := make([]byte, req.ContentLength)
	_, err := req.Body.Read(p)


	repoName := fmt.Sprintf("%s/%s", params["username"], params["repository"])

	if _, ok := params["branch"]; ok {
		repoName = fmt.Sprintf("%s/%s/%s", params["username"], params["repository"], params["branch"])
	}

	if err != nil {
		r.JSON(500, err)
		return
	}

	if _, ok := configuration.Repositories[repoName]; !ok {
		r.JSON(404, map[string]interface{}{"status": "error", "errors": []string{"Repository not configured"}})
		return
	}

	repo := configuration.Repositories[repoName]
	event := req.Header.Get("X-GitHub-Event")

	actions, ok := repo.Events[event]

	if false == ok {
		r.JSON(404, map[string]interface{}{"status": "error", "errors": []string{fmt.Sprintf("%s is not configured for this hook", event)}})
		return
	}

	var errs []error
	for _, cmdString := range actions {
		arguments := strings.Fields(cmdString)
		command := arguments[0]
		arguments = arguments[1:len(arguments)]
		cmd := exec.Command(command, arguments...)
		err := cmd.Run()
		if err != nil {
			errs = append(errs, err)

		}
	}

	if len(errs) > 0 {
		r.JSON(500, map[string]interface{}{"status": "error", "errors": errs})
		return
	}

	r.JSON(200, map[string]string{"status": "ok"})
	return
}
