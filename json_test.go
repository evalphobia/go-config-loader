package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadJSON(t *testing.T) {
	assert := assert.New(t)

	data := loadJSON("example/example.json")
	assert.NotEmpty(data)

	setting, ok := data["setting"].(map[string]interface{})
	assert.True(ok)

	assert.Equal(float64(100), setting["userid"])
	assert.Equal("John Doe", setting["username"])
	assert.Equal("bar", setting["foo"])
	assert.Equal(true, setting["isProd"])
}
