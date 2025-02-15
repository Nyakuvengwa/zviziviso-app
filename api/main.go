package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"text/template"
)

func main() {
	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		templateIndex := template.Must(template.ParseFiles("./api/index.html"))
		templateIndex.Execute(w, nil)
	}

	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":7653", nil))
}
