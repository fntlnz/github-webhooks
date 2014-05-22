package main

import (
    "flag"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "log"
    "os"
)

var configurationFilePath string
var configuration Configuration

func init() {
    flag.StringVar(&configurationFilePath, "configuration", "github-webhooks.json", "Configuration file path")
}

func main() {
    flag.Parse()

    if "" == configurationFilePath {
        log.Fatal("Configuration not provided")
    }

    _, err := os.Stat(configurationFilePath)

    if err != nil {
        log.Fatalf("Error: %s", err.Error())
    }

    configuration.Parse(configurationFilePath)
    m := martini.Classic()
    m.Use(render.Renderer())
    m.Map(configuration)
    Routes(m)
    m.Run()
}
