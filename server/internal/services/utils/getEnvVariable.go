package utils

import (
	"log"
	"os"
)

func GetEnvVariable(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal("Cannot get environment variable named \"" + key + "\"")
	}
	return value
}