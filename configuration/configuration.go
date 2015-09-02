package configuration

import (
	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	Port         string                 `json:"port"`
	Path         string                 `json:"path"`
	Repositories map[string]*Repository `json:"repositories"`
}

type Repository struct {
	Events map[string][]string `json:"events"`
}

func (c *Configuration) Parse(configuration []byte)  error {
	return json.Unmarshal(configuration, &c)
}

func (c *Configuration) ParseFile(filePath string) error {
	configuration, err := ioutil.ReadFile(filePath)

	if err != nil {
		return err
	}

	err = c.Parse(configuration)

	if err != nil {
		return err
	}

	return nil
}
