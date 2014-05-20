package main

import (
    "encoding/json"
    "fmt"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "net/http"
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

        jsonData := map[string]string{}
        err = json.Unmarshal(p, &jsonData)

        if err != nil {
            r.JSON(500, err)
            return
        }

        if _, ok := configuration.Repositories[repoName]; !ok {
            r.JSON(404, map[string]interface{}{"status": "error", "errors": []string{"Repository not found"}})
            return
        }

        repo := configuration.Repositories[repoName]

        if repo.Token != jsonData["token"] {
            r.JSON(401, map[string]interface{}{"status": "error", "errors": []string{"Invalid token"}})
            return
        }

        r.JSON(200, map[string]string{"cippa": repoName})
    })
}
