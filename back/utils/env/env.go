package env

import (
	"os"
	"strings"
)

func GetEnv(key string, fallbackValue string) string {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == key {
			return pair[1]
		}
	}
	return fallbackValue
}
