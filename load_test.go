package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigs(t *testing.T) {
	assert := assert.New(t)

	conf := NewConfig()
	assert.Empty(conf.data)
	conf.LoadConfigs("example", "json")
	assert.NotEmpty(conf.data)
	assert.Equal("John Doe", conf.ValueString("setting.username"))
	assert.Equal(100, conf.ValueInt("setting.userid"))
	assert.Equal(true, conf.ValueBool("setting.isProd"))
}
