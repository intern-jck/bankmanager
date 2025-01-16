package tmplmanager

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"bankmanager/types"
)

func Overview(w http.ResponseWriter, r *http.Request) {
	// get json data
	file, err := os.Open("data/test.json")
	if err != nil {

		http.Error(w, fmt.Sprintf("Error opening data: %v", err), http.StatusInternalServerError)
	}
	defer file.Close()

	jsonData := types.BankJson{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jsonData)

	// get template
	tmpl := template.Must(template.ParseFiles("templates/overview.html"))

	// serve template
	err = tmpl.Execute(w, jsonData)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
