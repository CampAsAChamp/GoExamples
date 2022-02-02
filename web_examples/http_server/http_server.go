package main

import (
	"fmt"
	"net/http"
)

/*
	A basic HTTP server has a few key jobs to take care of
	- Process Dynamic Requests: Process incoming requests from users who browse the website, log into their accounts or post images
	- Serve Static Assets: Server JS, CSS, and images to browsers to create a dynamic experience for the user
	- Accept Connections: Listen on a specific port to be able to accept connections from the internet

*/
func main() {
	// Takes a path to match and a function to execute
	// When somebody browses your website, they will be greated with a nice message
	// For the dynamic aspect, the Request contains all information about the request and it's parameters
	// You can read GET parameters with r.URL.Query().Get("token")
	// You can read POST parameters with r.FormValue("email")
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Welcome to my website!")
	})

	// To serve static assets like JS, CSS, and images, we can use the built-in http.FileServer and tell it where to serve the files from
	fs := http.FileServer(http.Dir("static/"))

	// Need to point a URL path at it
	// To serve files correctly, we need to strip away a part of the URL path.
	// Usually this is the name of the directory our files live in
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Listen on a port and accept connections from the internet
	http.ListenAndServe(":80", nil)

	fmt.Println("End of Main")
}
