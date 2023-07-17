package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", LoggerMiddleware(helloHandler))
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %s, Method: %s\n", r.URL.Path, r.Method)
		next(w, r)
		fmt.Printf("Response: %s, Status: %d\n", r.URL.Path, 200)
	}
}
