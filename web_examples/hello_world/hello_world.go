package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Registering a Request Handler

	// Create a handler which receives all incoming HTTP connections from browsers, HTTP clients or API requests
	// A request handler alone however can't accept any HTTP connections from the outside
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	// An HTTP server has to listen on a port to pass connectiosn on to the request handler
	// Start Go's default HTTP server and listen for connections on port 80
	http.ListenAndServe(":80", nil)
	fmt.Println("End of Main!")
}
