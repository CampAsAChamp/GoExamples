package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging - Logs all requests with its path and the time it took to process
func Logging() Middleware {

	// Create a new Middleware
	return func(hf http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(rw http.ResponseWriter, r *http.Request) {

			// Do Middleware things
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()

			// Call the next Middleware/Handler in the chain
			hf(rw, r)
		}
	}
}

// Method - Ensure that a URL can only be requested with a specific method
// Else returns a 400 Bad Request
func Method(m string) Middleware {

	// Create a new Middleware
	return func(hf http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(rw http.ResponseWriter, r *http.Request) {

			// Do Middleware things
			if r.Method != m {
				http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// Call the next Middleware/Handler in the chain
			hf(rw, r)
		}
	}
}

// Chain - Applies Middlewares to a http.HandlerFunc
func Chain(hf http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		hf = m(hf)
	}

	return hf
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {

	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":8080", nil)
	fmt.Println("End of Main")
}
