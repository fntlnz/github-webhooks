package main

import (
	"flag"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"os"
	"net/http"
	"fmt"
	"log"
)

func main() {
	var configurationFilePath string
	var configuration Configuration
	var coloured bool

	flag.StringVar(&configurationFilePath, "configuration", "github-webhooks.json", "Configuration file path")
	flag.BoolVar(&coloured, "colors", true, "Coloured output")
	flag.Parse()

	logger := NewLogger()
	logger.Coloured = coloured

	if "" == configurationFilePath {
		log.Fatal("Configuration not provided")
	}

	_, err := os.Stat(configurationFilePath)

	if err != nil {
		log.Fatal("Fatal Error: ", err.Error())
	}

	configuration.Parse(configurationFilePath)

	port := "3091"

	if configuration.Port != "" {
		port = configuration.Port
	}

	if (configuration.Path != "") {
		err := os.Setenv("PATH", configuration.Path);

		if err != nil {
			logger.WriteError(fmt.Sprintf("An error occurred setting PATH: %s", err))
		}
	}

	m := martini.Classic()
	m.Use(render.Renderer())
	m.Map(configuration)
	m.Map(logger)
	m.Map(NewNullLogger()) // Disable default martini logging
	Routes(m)
	logger.WriteInfo(fmt.Sprintf("GitHub Web Hooks Running on port %s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), m))
}
