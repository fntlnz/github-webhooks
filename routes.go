package main

import (
    "encoding/json"
    //"fmt"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "net/http"
)

//Routes ...
func Routes(m *martini.ClassicMartini) {
    m.Post("/:username/:repository", func(r render.Render, req *http.Request, params martini.Params, c Configuration) {

        p := make([]byte, req.ContentLength)
        _, err := req.Body.Read(p)

        r.JSON(200, c)
        //repoName := fmt.Sprintf("%s/%s", params["username"], params["repository"])

        if err != nil {
            r.JSON(500, err)
        }

        str := map[string]string{}

        err = json.Unmarshal(p, &str)
        if err != nil {
            r.JSON(500, err)
        }
        r.JSON(200, str)
    })
}
