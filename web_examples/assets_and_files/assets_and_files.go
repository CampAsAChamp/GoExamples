package main

import "net/http"

// Serve static files like CSS, JS, or images from a specific directory
func main() {

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
