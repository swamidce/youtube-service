package utils

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Converts string environment variable to int64. Returns default value if not found
func GetEnvInt(key string, defaultVal int64) int64 {
	s := os.Getenv(key)
	if s == "" {
		log.Errorf("Utils: environment variable %v not found, using default value %v.", key, defaultVal)
		return defaultVal
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Errorf("Utils: environment variable %v not found, using default value %v.", key, defaultVal)
		return defaultVal
	}
	return int64(v)
}
