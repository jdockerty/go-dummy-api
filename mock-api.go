package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

// HealthResponse provides a simple struct for providing a health check.
type HealthResponse struct {
	Message    string
	StatusCode int
}

// User struct is a minimal representation of simple data from the JSONPlaceholder API, only partial data is retreived.
type User struct {
	ID       int `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Users gives an intuitive way to reference a slice of User.
type Users []User

// OK is a helper function on providing the response for HealthResponse
func (r *HealthResponse) OK() []byte {

	resp := HealthResponse{
		Message:    "Success",
		StatusCode: http.StatusOK,
	}

	responseJSON, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		log.Fatalf("json populate: error when marshaling json response\n%s", err.Error())
	}

	return responseJSON

}

// HealthHandler is the function which is executed upon a request being routed to /health
func HealthHandler(w http.ResponseWriter, r *http.Request) {

	var response HealthResponse

	w.Header().Set("Content-Type", "application/json")
	w.Write(response.OK())

}

// AllUsersHandler is executed when a request is routed to /users
// This returns all 10 users within the JSONPlaceholder API, but with stripped down data that is only contained from the User struct. 
func AllUsersHandler(w http.ResponseWriter, r *http.Request) {

	var users Users

	usersResponse, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Fatalf("/users GET: error when retreiving response for all users\n%s", err.Error())
	}

	data, err := ioutil.ReadAll(usersResponse.Body)
	if err != nil {
		log.Fatalf("/users read response: error when reading data into byte array\n%s", err.Error())
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Fatalf("/users unmarshal: error when unmarshaling users into slice of structs\n%s", err.Error())
	}

	resp, _ := json.MarshalIndent(users, "", "\t")

	w.Header().Set("Content-Type", "application/json")


	w.Write(resp)
}

func main() {
	log.Println("Running server...")
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthHandler)
	r.HandleFunc("/users", AllUsersHandler)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
