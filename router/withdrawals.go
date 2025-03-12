package router

import (
	"html/template"
	"net/http"
)

func Withdrawals(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/account/withdrawals.html"))

	err := tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}
