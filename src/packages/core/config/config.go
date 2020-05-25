package config

import (
	"os"
)

func GetValue(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(key + " config is not set.")
	}

	return value
}
