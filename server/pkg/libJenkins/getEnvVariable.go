package libJenkins

import (
	"log"
	"os"
)

func getEnvVariable(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal("Cannot get environment variable named \"" + key + "\"")
	}
	return value
}
