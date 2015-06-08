package config

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// the string of Directory separator [/, \]
var Sep = string(filepath.Separator)

// LoadConfigs loads the config files from the path
func (c *Config) LoadConfigs(dir string, typ string) error {
	// check config type
	loadFunc, extName, err := prepareLoadFunc(typ)
	if err != nil {
		return err
	}

	// load file
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		fName := f.Name()
		if filepath.Ext(fName) != extName {
			continue
		}
		c.loadAndUpdateConfig(dir+Sep+fName, loadFunc)
	}
	return nil
}

// loadAndUpdateConfig adds data from loaded config to the unified config
func (c *Config) loadAndUpdateConfig(file string, loadFunc func(string) map[string]interface{}) {
	data := loadFunc(file)
	for key, val := range data {
		// skip if already set
		if _, ok := c.data[key]; ok {
			continue
		}
		c.data[key] = val
	}
}

// prepareLoadFunc returns func for file format of typ
func prepareLoadFunc(typ string) (func(string) map[string]interface{}, string, error) {
	switch typ {
	case "toml":
		return loadTOML, extTOML, nil
	default:
		return loadFuncError, "", errors.New("no matched load function")
	}
}

// empty func
func loadFuncError(s string) map[string]interface{} {
	return make(map[string]interface{})
}
