package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Go's HTTP package doesn't do complex request routing very well like segmenting a request URL into single parameters
// We will be using Gorilla/Mux to create routes with named parameters, GET/POST handlers and domain restrictions
func main() {
	// Creating a new Router
	// First we will create a new request router, it will later be passed as a parameter to the server
	// It will receive all HTTP connections and pass it on to the request handlers you will register on it
	muxRouter := mux.NewRouter()

	// Now we can register Request Handlers like usual
	muxRouter.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Hi there! This is the home page")
	})

	// The biggest strength of gorilla/mux Router is the ability to extract segments from the request URL
	// ex: /books/go-programming-blueprint/page/10
	// This URL has two dynamic segments
	// - Book Title Slug (go-programming-blueprint)
	// - Page (10)
	muxRouter.HandleFunc("/books/{title}/page/{page}", func(rw http.ResponseWriter, r *http.Request) {
		// Get the Book
		// Navigate to the Page
		URLVars := mux.Vars(r)
		title := URLVars["title"]
		page := URLVars["page"]
		fmt.Fprintf(rw, "Title: %s\n", title)
		fmt.Fprintf(rw, "Page: %s\n", page)

	})

	http.ListenAndServe(":80", muxRouter)
	fmt.Println("End of main")
}
