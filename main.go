package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform error : %v", err)
		return
	}
	fmt.Fprintf(w, "post request successful \n")
	name := r.FormValue("name")
	email := r.FormValue("email")
	fmt.Fprintf(w, "Name : %s\n", name)
	fmt.Fprintf(w, "Email : %s \n", email)
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}
func main() {
	fileserver := http.FileServer(http.Dir("./static")) // static server
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler) // handler
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
