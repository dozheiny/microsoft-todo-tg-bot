package config

import (
	"os"
	"strings"
)

// Get will get key of environment and return value of key.
// if value is not exist, return error
func Get(key string) (string, error) {
	var value string

	value = os.Getenv(key)
	if strings.Compare(value, "") == 0 {
		return "", notFound
	}

	return value, nil
}
