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
    m.Post("/:username/:repository", func(r render.Render, req *http.Request, params martini.Params, c Configuration) {
        p := make([]byte, req.ContentLength)
        _, err := req.Body.Read(p)

        repoName := fmt.Sprintf("%s/%s", params["username"], params["repository"])

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
        actions := repo.Events[event]

        for _, cmdString := range actions {
            arguments := strings.Fields(cmdString)
            command := arguments[0]
            arguments = arguments[1:len(arguments)]
            cmd := exec.Command(command, arguments...)
            cmd.Run()
        }

        r.JSON(200, map[string]string{"status": "ok"})
        return
    })
}
