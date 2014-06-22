package main

import (
	"flag"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
	"os"
	"net/http"
	"fmt"
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

	port := ":3091"

	if configuration.Port != "" {
		port = fmt.Sprintf(":%s", configuration.Port)
	}

	m := martini.Classic()
	logger := log.New(os.Stdout, "[github-webhooks] ", 0)
	m.Use(render.Renderer())
	m.Map(configuration)
	m.Map(logger)
	Routes(m)
	logger.Printf("Running on port %s", port)
	log.Fatal(http.ListenAndServe(port, m))
}



