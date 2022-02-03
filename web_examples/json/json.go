package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Only exported fields will be encoded/decoded in JSON
// Fields must start with capital letters to be exported
type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

// How to encode/decode JSOn data using the encoding/json package
func main() {

	http.HandleFunc("/decode", func(rw http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		fmt.Fprintf(rw, "%s %s is %d years old!\n", user.Firstname, user.Lastname, user.Age)
	})

	http.HandleFunc("/encode", func(rw http.ResponseWriter, r *http.Request) {
		peter := User{
			Firstname: "John",
			Lastname:  "Doe",
			Age:       25,
		}

		json.NewEncoder(rw).Encode(peter)
	})

	http.ListenAndServe(":8080", nil)

	fmt.Println("End of Main")
}
