package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Configuration struct {
	Host         string                 `json:"host"`
	Port         string                 `json:"port"`
	Path         string                 `json:"path"`
	Repositories map[string]*Repository `json:"repositories"`
}

type Repository struct {
	Events map[string][]string `json:"events"`
}

func NewConfiguration() *Configuration {
	config := new(Configuration)
	config.Host = "0.0.0.0"
	config.Port = "3091"
	return config
}

func (c *Configuration) Parse(configuration []byte) error {
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

func (c Configuration) GetAddress() string {
	if len(c.Host) > 0 {
		return fmt.Sprintf("%s:%s", c.Host, c.Port)
	}
	return fmt.Sprintf(":%s", c.Port)
}
