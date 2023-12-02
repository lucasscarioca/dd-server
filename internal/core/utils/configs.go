package utils

import (
	"bytes"
	"encoding/json"
)

// equals to string "null"
var NULL_JSON = []byte{110, 117, 108, 108}

// Checks if b is an empty slice or if it represents an empty JSON object
func EmptyConfigs(b []byte) bool {
	if len(b) == 0 || b == nil || bytes.Equal(b, NULL_JSON) {
		return true
	}

	configsJson, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return false
	}

	if bytes.Equal(b, configsJson) {
		return true
	}

	return false
}
