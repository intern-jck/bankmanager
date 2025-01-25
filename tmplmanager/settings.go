package tmplmanager

import (
	"net/http"
	"text/template"
)

// const defaultStatement = "data/json/2018/20180104.json"

func Settings(w http.ResponseWriter, r *http.Request) {

	// Get template
	tmpl := template.Must(template.ParseFiles("templates/settings/settings.html"))

	// Format bank data struct to pass to template
	// jsonData, _ := json.Marshal(bankData.Summary)
	// fmt.Println(string(jsonData))

	// Serve template with formatted data
	// err = tmpl.Execute(w, string(jsonData))
	err := tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}
