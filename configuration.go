package main

import (
	"encoding/json"
	"io/ioutil"
)

// Configuration configuration struct
type Configuration struct {
	Repositories []string
}

// AddRepository ...
func (c *Configuration) AddRepository(repository string) {
	c.Repositories = append(c.Repositories, repository)
}

// Parse ...
func (c *Configuration) Parse(filePath string) {
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		//TODO handle error
	}

	err = json.Unmarshal(file, c)

	if err != nil {
		//TODO handle error
	}
}
