package main

import (
	"fmt"
	"net/http"
)

func StartServer() {

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/json", JSONHandler)
	fmt.Println("Server Started: http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

