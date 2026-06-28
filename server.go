package main

import (
	"fmt"
	"net/http"
)

func StartServer() {

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/json", JSONHandler)
	http.HandleFunc("/user", UserHandler)
	http.HandleFunc("/method", MethodHandler)
	http.HandleFunc("/users", CreateUserHandler)

	fmt.Println("Server Started: http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server Error:", err)
	}
}