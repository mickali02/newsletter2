// Filename: cmd/web/main.go
package main

import (
	"log"
	"net/http"
	"html/template"
)

// A handler function named 'home'
func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("ui/html/home.tmpl")
	if err != nil { 
     log.Print(err.Error())
     http.Error(w, "Internal Server Error",http.StatusInternalServerError)   
     return 
	}
	data := map[string]string {
    "Title": "Tuesday | Thursday",
    "Greeting": "Good morning!",
    "Message": "I am having a fantastic day today",
	}
	err = tmpl.Execute(w, data)
	if err != nil { 
		log.Print(err.Error())
		http.Error(w, "Internal Server Error",http.StatusInternalServerError)   
		return 
   }
   
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
