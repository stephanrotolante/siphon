package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.Handle("/proxy-route", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Server", r.Method)

		// Set the Content-Type header to indicate JSON
		w.Header().Set("Content-Type", "application/text")

		// Write the JSON data to the response writer
		w.Write([]byte("Proxy res"))
	}))

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
