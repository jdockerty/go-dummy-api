package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"log"
)

func TestHealthHandler(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Errorf("error sending request to /health route\n%s", err.Error())
	}

	response := httptest.NewRecorder()

	HealthHandler(response, request)

	receivedResponse := response.Body.Bytes()

	var testHealthResponse HealthResponse

	json.Unmarshal(receivedResponse, &testHealthResponse)

	required := 200
	got := testHealthResponse.StatusCode
	if got != required {
		t.Errorf("Got %q, required %q", got, required)
	}

}

func ExampleHealthResponse() {
	request, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		log.Fatalf("error generate request in ExampleHealthResponse()\n%s", err.Error())
	}

	response := httptest.NewRecorder()

	HealthHandler(response, request)

	receivedResponse := response.Body.String()

	fmt.Println(receivedResponse)

	// Output:
	// {
	// 	"Message": "Success",
	// 	"StatusCode": 200
	// }

}