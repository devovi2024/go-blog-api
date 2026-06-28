package main

import (
	"fmt"
	"net/http"
)

func StartServer() {

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)
	fmt.Println("Server Started: http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

