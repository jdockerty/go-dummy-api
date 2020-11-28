package logger

import (
	log "github.com/sirupsen/logrus"
)

func Run() {
	log.SetFormatter(&log.JSONFormatter{})

	fields := log.Fields{
		"ID":  "",
		"app": "dummy-api",
	}

	log.WithFields(fields).WithFields(log.Fields{"string": "foo"}).Info("Event from Logger.")
}
