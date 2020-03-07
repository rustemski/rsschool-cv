package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/styles/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", handler1)

	fmt.Printf("Starting server for testing HTTP POST...\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handler1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("singolo1.html"))
	tmpl.Execute(w, "singolo1.html")
}
