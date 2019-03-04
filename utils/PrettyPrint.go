// Package utils provides various helper functions for commonly tasks
package utils

import (
	"encoding/json"
)

// PrettyPrint print a struct with pretty indentation
func PrettyPrint(v interface{}) (s string, err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		return string(b), nil
	}
	return "", err
}
