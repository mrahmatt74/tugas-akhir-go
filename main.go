package main

import (
	"html/template"
	"net/http"
)

type Person struct {
	Name string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("main.html"))
		p := Person{Name: "John"}
		t.Execute(w, p)
	})

	map[string] int {
		"a": 12345679,
		"b": -1234567,
		"c": +1234567
	}

	http.ListenAndServe(":8080", nil)
}
