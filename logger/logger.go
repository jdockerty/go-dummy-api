package logger

import (
	"io"
	"log"
	"os"
	"github.com/sirupsen/logrus"
)

// var (
// 	logwrapper LogWrap
// )

// type Event struct {
// 	id      string
// 	message string
// }

type Logger struct {
	*logrus.Logger
}

func(l *Logger) DebugAPIMessage(apiId, msg string) {
	l.WithField("id", apiId).Debug(msg)
} 

func New() *Logger {

	baseLogrus := logrus.New()

	var logger = &Logger{baseLogrus}

	f, err := os.OpenFile("/var/log/dummy-api.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("unable to interact with log file: %s", err)
	}

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "02-01-2006 15:04:05", // DD-MM-YYYY HH:MM:SS

	})

	outputs := io.MultiWriter(os.Stderr, f) // Write to both standard error and the log file.
	logger.Out = outputs


	return logger

}


