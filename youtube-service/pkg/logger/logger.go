package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Redirect formatted logs to a file if available else prints them
func InitLogger() {
	log.SetOutput(os.Stdout)

	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	f, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(f)
	} else {
		log.Info("Failed to log to file, using default stdout")
	}
}
