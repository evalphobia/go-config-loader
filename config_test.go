package config

import (
	"fmt"
	"os"
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

func TestSetExpandEnvOn(t *testing.T) {
	assert := assert.New(t)

	c := NewConfig()
	assert.False(c.useExpandEnv)

	c.SetExpandEnvOn()
	assert.True(c.useExpandEnv)
}

func TestSetExpandEnvOff(t *testing.T) {
	assert := assert.New(t)

	c := NewConfig()
	c.useExpandEnv = true
	assert.True(c.useExpandEnv)

	c.SetExpandEnvOff()
	assert.False(c.useExpandEnv)
}

func TestExpandEnv(t *testing.T) {
	assert := assert.New(t)

	os.Clearenv()
	defer os.Clearenv()
	os.Setenv("TEST_STRING", "ENV_TEST_STRING")
	os.Setenv("test_string", "env_test_string")
	os.Setenv("TEST_INT", "9")
	os.Setenv("TEST_BOOL", "true")

	tests := []struct {
		key   interface{}
		value interface{}
	}{
		{"$TEST_STRING", "ENV_TEST_STRING"},
		{"$test_string", "env_test_string"},
		{"$TEST_INT", "9"},
		{"$TEST_BOOL", "true"},
		{9, 9},
		{true, true},
		{false, false},
	}

	c := NewConfig()
	for _, tt := range tests {
		target := fmt.Sprintf("%+v", tt)

		c.SetExpandEnvOff()
		assert.Equal(tt.key, c.expandEnv(tt.key), target)

		c.SetExpandEnvOn()
		assert.Equal(tt.value, c.expandEnv(tt.key), target)
	}
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
