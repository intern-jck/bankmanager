package router

import (
	"html/template"
	"net/http"
)

func Summary(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/summary.html"))

	err := tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}
