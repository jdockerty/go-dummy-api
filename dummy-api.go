package main

import (
	"github.com/jdockerty/go-dummy-api/logger"
)

var (
	myLogger *logger.Logger
)

// HealthResponse provides a simple struct for providing a health check.
type HealthResponse struct {
	ID         string
	Message    string
	StatusCode int
}

// User struct is a minimal representation of simple data from the JSONPlaceholder API, only partial data is retreived.
type User struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
}

// Users gives an intuitive way to reference a slice of User.
type Users []User

func main() {

	myLogger = logger.New()
	myLogger.Info("Running server")

	api := NewAPI()
	api.Listen()
}
