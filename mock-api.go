package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type HealthResponse struct {
	Message    string
	StatusCode int
}

type User struct {
	ID       int `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Users []User

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

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	var response HealthResponse

	w.Header().Set("Content-Type", "application/json")
	w.Write(response.OK())

}

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
