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
func Routes(m *martini.ClassicMartini,) {
	m.Post("/:username/:repository/:branch", func(r render.Render, req *http.Request, params martini.Params, c Configuration, l *Logger) {
			repositoryAction(r, req, params, c, l)
		})

	m.Post("/:username/:repository", func(r render.Render, req *http.Request, params martini.Params, c Configuration, l *Logger) {
			repositoryAction(r, req, params, c, l)
		})
}

func repositoryAction(r render.Render, req *http.Request, params martini.Params, c Configuration, l *Logger) {

	repoName := fmt.Sprintf("%s/%s", params["username"], params["repository"])

	if _, ok := params["branch"]; ok {
		repoName = fmt.Sprintf("%s/%s/%s", params["username"], params["repository"], params["branch"])
	}

	if _, ok := c.Repositories[repoName]; !ok {
		l.WriteAlert(fmt.Sprintf("Repository not configured: `%s`", repoName))
		r.JSON(404, map[string]interface{}{"status": "error", "errors": []string{"Repository not configured"}})
		return
	}

	repo := c.Repositories[repoName]
	event := req.Header.Get("X-GitHub-Event")

	l.WriteInfo(fmt.Sprintf("Event `%s` called on repository: `%s`", event, repoName))
	if event == "" {
		l.WriteAlert("Event not passed")
		r.JSON(400, map[string]interface{}{"status": "error", "errors": []string{"Event not provided"}})
		return
	}

	actions, ok := repo.Events[event]

	if false == ok {
		l.WriteAlert(fmt.Sprintf("Event not found for hook: `%s`", event))
		r.JSON(404, map[string]interface{}{"status": "error", "errors": []string{fmt.Sprintf("`%s` event is not configured for this hook", event)}})
		return
	}

	for _, cmdString := range actions {
		arguments := strings.Fields(cmdString)

		command := arguments[0]

		arguments = arguments[1:len(arguments)]
		l.WriteInfo(fmt.Sprintf("Executing command: `%s`", cmdString))
		cmd := exec.Command(command, arguments...)
		err := cmd.Run()
		if err != nil {
			l.WriteError(fmt.Sprintf("command: `%s` - Error: `%s`", cmdString, err))
			r.JSON(500, map[string]interface{}{"status": "error", "command": cmdString, "errors": []string{fmt.Sprintf("%s", err)}})
			return
		}
		l.WriteSuccess(fmt.Sprintf("Command executed"))
	}

	l.WriteSuccess("Hook executed successfully")
	r.JSON(200, map[string]string{"status": "ok"})
	return
}
