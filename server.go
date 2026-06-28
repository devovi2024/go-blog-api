package main

import (
	"fmt"
	"net/http"
)

func StartServer() {

	http.HandleFunc("/", HomeHandler)

	fmt.Println("Server Started: http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

