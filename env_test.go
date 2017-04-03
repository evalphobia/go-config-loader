package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValueFromEnv(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		key      string
		env      string
		hasValue bool
	}{
		{"env_value", "env_value", true},
		{"env_value", "ENV_VALUE", true},
		{"env.value", "env_value", true},
		{"env..value", "env__value", true},
		{"env...value", "env___value", true},
		{".e.n.v.v.a.l.u.e.", "_e_n_v_v_a_l_u_e_", true},
		{"env_value", "envvalue", false},
		{"env_value", "env.value", false},
	}

	for _, tt := range tests {
		target := fmt.Sprintf("%+v", tt)
		os.Clearenv()
		os.Setenv(tt.env, VALUE_STRING)

		result, ok := GetValueFromEnv(tt.key)
		if !tt.hasValue {
			assert.False(ok, target)
			assert.Empty(result, target)
			continue
		}

		assert.True(ok, target)
		assert.Equal(VALUE_STRING, result, target)
	}
}
