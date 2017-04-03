package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	VALUE_INT    = 9
	VALUE_STRING = "string"
	VALUE_BOOL   = true
)

func TestNewConfig(t *testing.T) {
	assert := assert.New(t)

	c := NewConfig()
	assert.NotNil(c)
	assert.NotNil(c.data)
	assert.Empty(c.data)
}

func TestGetConfigValues(t *testing.T) {
	assert := assert.New(t)

	c := testConfig()
	assert.Equal(c.data, c.GetConfigValues())
}

func TestValue(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		key      string
		hasValue bool
		value    interface{}
	}{
		{"int", true, VALUE_INT},
		{"string", true, VALUE_STRING},
		{"bool", true, VALUE_BOOL},
		{"no_key", false, nil},
	}

	c := testConfig()
	for _, tt := range tests {
		target := fmt.Sprintf("%+v", tt)
		result, err := c.Value(tt.key)
		if !tt.hasValue {
			assert.Error(err, target)
			assert.Nil(result, target)
			continue
		}

		assert.NoError(err, target)
		assert.Equal(tt.value, result, target)
	}
}

func testConfig() *Config {
	c := NewConfig()
	c.data = map[string]interface{}{
		"int":    VALUE_INT,
		"string": VALUE_STRING,
		"bool":   VALUE_BOOL,
	}
	return c
}
