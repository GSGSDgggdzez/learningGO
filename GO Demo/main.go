// Package main is the entry point of the Go program.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func formhandel(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Password = %s\n", password)
}

func main() {
	fileserver := http.FileServer(http.Dir("."))
	http.Handle("/", http.StripPrefix("/", fileserver))
	http.HandleFunc("/form", formhandel)

	http.HandleFunc("/hello", hellohandel)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func hellohandel(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello bro")
}
