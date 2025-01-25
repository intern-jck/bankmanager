package tmplmanager

import (
	"net/http"
	"text/template"
)

const defaultStatementId = "20180104"

func Index(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles(
		"templates/index.html",
		"templates/components/app-dashboard.html",
		"templates/components/app-content.html"))

	err := tmpl.Execute(w, defaultStatementId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
