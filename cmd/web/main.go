// Filename: cmd/web/main.go
package main

import (
	"log"
	"net/http"
)

// A handler function named 'home'
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from newsletter"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this the about page"))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this the sign up page"))
}

func main() {

	// mux is our router (multiplexer)
	mux := http.NewServeMux()

	// the route pattern/endpoint/URL path
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/signup", signup)

	// start a local web server
	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
