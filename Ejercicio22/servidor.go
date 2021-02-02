package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template, _ := template.ParseFiles("../html/index.html")
		template.Execute(w, nil)
	})

	http.ListenAndServe(":8000", nil)
}

type Persona struct {
}
