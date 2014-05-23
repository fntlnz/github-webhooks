package main

import (
	"encoding/json"
	"io/ioutil"
)

// Configuration configuration struct
type Configuration struct {
	Port string `json:"port"`
	Repositories map[string]*Repository `json:"repositories"`
}

// Repository ...
type Repository struct {
	Events map[string][]string `json:"events"`
}

// Parse ...
func (c *Configuration) Parse(filePath string) error {
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &c)

	if err != nil {
		return err
	}

	return nil
}
