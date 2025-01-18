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

	bankData := types.BankJson{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&bankData)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "json data error: "+err.Error(), http.StatusInternalServerError)
	}

	// get template
	tmpl := template.Must(template.ParseFiles("templates/overview.html"))

	jsonData, _ := json.Marshal(bankData)

	// serve template
	err = tmpl.Execute(w, string(jsonData))

	if err != nil {
		fmt.Println(err)
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}
