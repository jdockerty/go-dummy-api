package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// API is a simple struct to provide access to an internal ID for the running server.
type API struct {
	ID string
}

// Listen will start the server, awaiting requests on port 8080.
func (api *API) Listen() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// HealthHandler is the function which is executed upon a request being routed to /health
func (api *API) HealthHandler(w http.ResponseWriter, r *http.Request) {

	myLogger.InfoAPIMessage(api.ID, "request received on /health")
	resp := HealthResponse{
		ID:         api.ID,
		Message:    "Success",
		StatusCode: http.StatusOK,
	}

	respJSON, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		log.Fatalf("json populate: error when marshaling json response\n%s", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)

	myLogger.InfoAPIMessage(api.ID, "returned a response to /health")
}

// AllUsersHandler is executed when a request is routed to /users
// This returns all 10 users within the JSONPlaceholder API, but with stripped down data that is only contained from the User struct.
func (api *API) AllUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users Users

	myLogger.InfoAPIMessage(api.ID, "starting all users request")

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

	myLogger.InfoAPIMessage(api.ID, "response sent for all users")
}

// SingleUserHandler functions in a similar way to AllUsersHandler, except a single user is returned.
// The specific user ID is retreived from the URL path.
func (api *API) SingleUserHandler(w http.ResponseWriter, r *http.Request) {

	var user User
	id := mux.Vars(r)["id"]

	logMsg := fmt.Sprintf("request for /users/%s", id)
	myLogger.InfoAPIMessage(api.ID, logMsg)

	endpointWithID := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%s", id)
	userResp, err := http.Get(endpointWithID)
	if err != nil {
		log.Fatalf("/users/%s GET: error when retreiving response for %s\n%s", id, id, err.Error())
	}

	data, err := ioutil.ReadAll(userResp.Body)
	if err != nil {
		log.Fatalf("/users/%s read response: error when reading data into byte array\n%s", id, err.Error())
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		log.Fatalf("/users/%s unmarshal: error when unmarshaling user\n%s", id, err.Error())
	}

	resp, _ := json.MarshalIndent(user, "", "\t")

	w.Header().Set("Content-Type", "application/json")

	w.Write(resp)

	logMsg = fmt.Sprintf("response sent for /users/%s", id)
	myLogger.InfoAPIMessage(api.ID, logMsg)
}

// Helper function for identifying different APIs when running multiple instances
// e.g. through a load balancer to verify routing or in different containers.
func makeID() string {
	myLogger.Debug("Generating API ID")

	source := rand.NewSource(time.Now().UnixNano())
	num := rand.New(source).Intn(5000)
	apiID := fmt.Sprintf("api-%d", num)

	myLogger.InfoAPIMessage(apiID, "Generated API ID")
	return apiID
}

// NewAPI does stuff
func NewAPI() *API {
	api := new(API)
	newID := makeID()

	api.ID = newID

	r := mux.NewRouter()
	r.HandleFunc("/health", api.HealthHandler)
	r.HandleFunc("/users", api.AllUsersHandler)
	r.HandleFunc("/users/{id:[0-9]+}", api.SingleUserHandler)
	http.Handle("/", r)

	return api
}
