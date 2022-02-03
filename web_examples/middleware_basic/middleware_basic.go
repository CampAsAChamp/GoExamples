package main

import (
	"fmt"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(rw, r)
	}
}

func foo(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "foo there")
}

func bar(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "bar there")
}

// Create basic logging Middleware
// Middleware simply takes a http.HandlerFunc as one of its parameters, wraps it and returns a new http.Handlerfunc for the server to call
func main() {

	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))

	http.ListenAndServe(":8080", nil)
	fmt.Println("End of Main")
}
