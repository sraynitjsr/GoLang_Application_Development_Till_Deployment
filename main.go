package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintln(w, "Error while Parsing the form", err)
	}
	name := r.FormValue("name")
	hobbies := r.FormValue("hobbies")
	fmt.Fprintln(w, name, "has hobbies", hobbies)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Wrong HTTP Method", http.StatusNotAcceptable)
	}
	fmt.Fprintln(w, "Hello World")
}

func main() {
	fileServer := http.FileServer(http.Dir(""))

	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting Server at 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
