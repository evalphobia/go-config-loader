package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config is cofig data struct
type Config struct {
	data         map[string]interface{}
	useExpandEnv bool
}

// NewConfig creates empty Config struct
func NewConfig() *Config {
	return &Config{
		data: make(map[string]interface{}),
	}
}

// SetExpandEnvOn enables to expand env variable in the result value string.
func (c *Config) SetExpandEnvOn() {
	c.useExpandEnv = true
}

// SetExpandEnvOff disables to expand env variable in the result value string.
func (c *Config) SetExpandEnvOff() {
	c.useExpandEnv = false
}

func (c *Config) expandEnv(value interface{}) interface{} {
	if !c.useExpandEnv {
		return value
	}

	if vv, ok := value.(string); ok {
		return os.ExpandEnv(vv)
	}
	return value
}

// GetConfigValues gets config values
func (c *Config) GetConfigValues() map[string]interface{} {
	return c.data
}

// Value returns the config value from key
func (c *Config) Value(key string) (interface{}, error) {
	if v, ok := GetValueFromEnv(key); ok {
		return v, nil
	}

	keys := strings.Split(key, ".")
	maxIndex := len(keys) - 1
	conf := c.data
	for i, k := range keys {
		v, ok := conf[k]
		if !ok {
			return nil, errors.New("no data on config, key=" + key)
		}
		if i >= maxIndex {
			return v, nil
		}

		data, ok := v.(map[string]interface{})
		if ok {
			conf = data
		}
	}
	return c.expandEnv(conf[key]), nil
}

// ValueString returns the config value with string type
func (c *Config) ValueString(keys string) string {
	v, err := c.Value(keys)
	if err != nil {
		return ""
	}
	switch t := v.(type) {
	case string:
		return t
	default:
		return fmt.Sprint(t)
	}
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
	switch iv := v.(type) {
	case int:
		return iv
	case int8:
		return int(iv)
	case int16:
		return int(iv)
	case int32:
		return int(iv)
	case int64:
		return int(iv)
	case uint:
		return int(iv)
	case uint8:
		return int(iv)
	case uint16:
		return int(iv)
	case uint32:
		return int(iv)
	case uint64:
		return int(iv)
	case float32:
		return int(iv)
	case float64:
		return int(iv)
	case string:
		i, _ := strconv.Atoi(iv)
		return i
	}
	return 0
}

// ValueIntDefault returns the config value with int type or returns defaut data when missing
func (c *Config) ValueIntDefault(keys string, def int) int {
	v := c.ValueInt(keys)
	if v == 0 {
		return def
	}
	return v
}

// ValueFloat returns the config value with float64 type
func (c *Config) ValueFloat(keys string) float64 {
	v, err := c.Value(keys)
	if err != nil {
		return 0
	}
	switch iv := v.(type) {
	case int:
		return float64(iv)
	case int8:
		return float64(iv)
	case int16:
		return float64(iv)
	case int32:
		return float64(iv)
	case int64:
		return float64(iv)
	case uint:
		return float64(iv)
	case uint8:
		return float64(iv)
	case uint16:
		return float64(iv)
	case uint32:
		return float64(iv)
	case uint64:
		return float64(iv)
	case float32:
		return float64(iv)
	case float64:
		return iv
	case string:
		f, _ := strconv.ParseFloat(iv, 64)
		return f
	}
	return 0
}

// ValueFloatDefault returns the config value with float64 type or returns defaut data when missing
func (c *Config) ValueFloatDefault(keys string, def float64) float64 {
	v := c.ValueFloat(keys)
	if v == 0 {
		return def
	}
	return v
}

// ValueBool returns the config value with bool type
func (c *Config) ValueBool(keys string) bool {
	v, err := c.Value(keys)
	if err != nil {
		return false
	}
	switch iv := v.(type) {
	case bool:
		return iv
	case string:
		b, _ := strconv.ParseBool(iv)
		return b
	}
	return false
}
