package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Go Blog API")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About Page")
}

func JSONHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome to Go Blog API",
		"status":  "success",
	})
}

func UserHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	fmt.Fprintf(w, "User ID: %s", id)
}

func MethodHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Request Method: %s", r.Method)
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		fmt.Fprintln(w, "Only POST method is allowed")
		return
	}

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	fmt.Fprintf(w, "Name: %s, Age: %d", user.Name, user.Age)
}