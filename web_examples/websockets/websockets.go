package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//panic if s is not a slice
func reverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

// We will build a simple server which echoes back everything we send to it
func main() {

	http.HandleFunc("/echo", func(rw http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(rw, r, nil) // Error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			reverseSlice(msg)
			// Write the message back to the browser reversed
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}

		}

	})

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)

	fmt.Println("End of Main")
}
