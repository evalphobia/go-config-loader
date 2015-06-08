package config

import (
	"errors"
	"strings"
)

// Config is cofig data struct
type Config struct {
	data map[string]interface{}
}

// NewConfig creates empty Config struct
func NewConfig() *Config {
	return &Config{
		data: make(map[string]interface{}),
	}
}

// Value returns the config value from key
func (c *Config) Value(key string) (interface{}, error) {
	keys := strings.Split(key, ".")
	maxIndex := len(keys) - 1

	for i, k := range keys {
		v, ok := c.data[k]
		if !ok {
			return nil, errors.New("no data on config, key=" + key)
		}
		if i >= maxIndex {
			return v, nil
		}

		data, ok := v.(map[string]interface{})
		if ok {
			c.data = data
		}
	}
	return c.data[key], nil
}

// ValueString returns the config value with string type
func (c *Config) ValueString(keys string) string {
	v, err := c.Value(keys)
	if err != nil {
		return ""
	}
	return v.(string)
}

// ValueStringDefault returns the config value with string type or returns defaut data when missing
func (c *Config) ValueStringDefault(keys string, def string) string {
	v := c.ValueString(keys)
	if v == "" {
		return def
	}
	return v
}

// ValueInt returns the config value with int type
func (c *Config) ValueInt(keys string) int {
	v, err := c.Value(keys)
	if err != nil {
		return 0
	}
	return int(v.(int64))
}

// ValueIntDefault returns the config value with int type or returns defaut data when missing
func (c *Config) ValueIntDefault(keys string, def int) int {
	v := c.ValueInt(keys)
	if v == 0 {
		return def
	}
	return int(v)
}

// ValueBool returns the config value with bool type
func (c *Config) ValueBool(keys string) bool {
	v, err := c.Value(keys)
	if err != nil {
		return false
	}
	return v.(bool)
}
