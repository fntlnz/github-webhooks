package main

import (
    "crypto/hmac"
    "crypto/sha1"
    "encoding/json"
    "io/ioutil"
)

// Configuration configuration struct
type Configuration struct {
    Repositories map[string]*Repository `json:"repositories"`
}

// Repository ...
type Repository struct {
    Secret string `json:"secret"`
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

// CheckSecret ...
func (r *Repository) CheckSecret(expectedSecret []byte) bool {
    mac := hmac.New(sha1.New, []byte(r.Secret))
    return hmac.Equal(mac.Sum(nil), expectedSecret)
}
