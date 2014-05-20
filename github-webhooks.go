package main

import (
    "flag"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "log"
)

var configurationFilePath string
var configuration Configuration

func init() {
    flag.StringVar(&configurationFilePath, "configuration", "", "Configuration file path")
}

func main() {
    flag.Parse()

    if "" == configurationFilePath {
        log.Fatal("Configuration not provided")
    }

    configuration.Parse(configurationFilePath)
    m := martini.Classic()
    m.Use(render.Renderer())
    m.Map(configuration)
    Routes(m)
    m.Run()
}
