package config

import (
	"log"
	"os"
	"strconv"
)

func GetString(key string) string {
	return os.Getenv(key)
}

func GetInt(key string) int {
	v, err := strconv.Atoi(GetString(key))
	mustParseKey(err, key)
	return v
}

func mustParseKey(err error, key string) {
	if err != nil {
		log.Fatalf("Could not parse key: %s, error: %s", key, err)
	}
}
