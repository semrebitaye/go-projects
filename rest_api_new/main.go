package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var profiles []profile = []profile{}

type profile struct {
	Department  string `json: "department"`
	Designation string `json: "department"`
	Employee    User   `json: "employee`
}

type User struct {
	FirstName string `json: "first name"`
	LastName  string `json: "last name"`
	Email     string `json: "email"`
}

func additem(w http.ResponseWriter, r *http.Request) {
	var newProfile profile

	json.NewDecoder(r.Body).Decode(&newProfile)

	profiles = append(profiles, newProfile)

	json.NewEncoder(w).Encode(profiles)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/profiles", additem).Methods("post")

	http.ListenAndServe(":5000", router)
}
