package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		for key, values := range r.Header {
			fmt.Println("Header:", key)
			for _, v := range values {
				fmt.Println(fmt.Sprintf("\t%s\n", v))
			}
		}
		// Set the Content-Type header to indicate JSON
		w.Header().Set("fake-fake", "random-header")

		// Write the JSON data to the response writer
		w.Write([]byte("This is some dumb shit"))
	}))

	fmt.Println("Starting server...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
