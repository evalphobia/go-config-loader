package config

import (
	"encoding/json"
	"io/ioutil"
)

const (
	extJSON = ".json"
)

// loadJSON loads config values in the JSON file
func loadJSON(file string) map[string]interface{} {
	c := make(map[string]interface{})
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return c
	}
	json.Unmarshal(f, &c)
	return c
}
