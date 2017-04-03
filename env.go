package config

import (
	"os"
	"strings"
)

// GetValueFromEnv returns the value from environment variable.
// Try three keys from the given key;
// (1) raw key, (2) dot to underscore, (3) 2 + Upper Case
// e.g. If you give `aws.access_key_id`, seek these order,
// (1) aws.access_key_id, (2) aws_access_key_id, (3) AWS_ACCESS_KEY_ID
func GetValueFromEnv(key string) (value string, ok bool) {
	if v := os.Getenv(key); v != "" {
		return v, true
	}

	noDot := strings.Replace(key, ".", "_", -1)
	if v := os.Getenv(noDot); v != "" {
		return v, true
	}
	if v := os.Getenv(strings.ToUpper(noDot)); v != "" {
		return v, true
	}

	return "", false
}
