package config

import (
	"github.com/BurntSushi/toml"
)

const (
	extTOML = ".toml"
)

// loadTOML loads config values in the TOML file
func loadTOML(file string) map[string]interface{} {
	c := make(map[string]interface{})
	if _, err := toml.DecodeFile(file, &c); err != nil {
		return c
	}
	return c
}
