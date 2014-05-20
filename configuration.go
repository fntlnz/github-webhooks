package main

import (
    "encoding/json"
    "io/ioutil"
)

// Configuration configuration struct
type Configuration struct {
    Repositories []Repository `json:"repositories"`
}

// Repository ...
type Repository struct {
    Name string `json:"name"`
    Key  string `json:"key"`
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

// GetRepository ...
func (c *Configuration) GetRepository(repository string) {

}
