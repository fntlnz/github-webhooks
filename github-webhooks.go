package main

import (
	"flag"
	"fmt"
)

var configurationFilePath string
var configuration Configuration

func init() {
	flag.StringVar(&configurationFilePath, "configuration", "", "Configuration file path")
}

func main() {
	flag.Parse()
	configuration.Parse(configurationFilePath)
	fmt.Println(configuration.Repositories)
}
