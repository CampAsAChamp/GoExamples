package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

// How to store data in session cookies using the gorilla/sessions package
// Cookies are small pieces of data stored in teh browser of a user and are sent to our server on each request
// In them we can store whether or not a user is logged into our website and figure out who they actually are (in our system)
// In this example we will only allow authenticated users to view our secret message on the /secret page
// To get access to it, they will first have to visit /login to get a valid session cookie, which logs them in
// Additionally they can visit /logout to revoke access to our secret message

var (
	// Key must be 16, 24, or 32 bytes long (AES-128, AES-192, or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, inMap := session.Values["authenticated"].(bool); !auth || inMap {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print Secret Message
	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ..

	// Set user as authetnicated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ..

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func main() {

	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)

	fmt.Println("End of Main")
}
