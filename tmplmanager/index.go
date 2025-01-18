package tmplmanager

import (
	"net/http"
	"text/template"
)

type Test struct {
	Name   string
	Value  int
	Floats []float32
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	data := Test{
		Name:   "test",
		Value:  123,
		Floats: []float32{123.4, 567.8, 901.2},
	}

	err := tmpl.Execute(w, template.JSEscape())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
