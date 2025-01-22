package tmplmanager

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"bankmanager/types"
)

// const defaultStatement = "data/json/2018/20180104.json"

func Overview(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	fmt.Printf("ID: %v\n", id)

	year := id[:4]

	statementPath := "data/json/" + year + "/" + id + ".json"
	fmt.Println("GETTING: ", statementPath)

	// get json data
	file, err := os.Open(statementPath)
	if err != nil {
		fmt.Println("json open error", err)
		http.Error(w, fmt.Sprintf("Error opening data: %v", err), http.StatusInternalServerError)
	}
	defer file.Close()

	// Create bank data struct
	bankData := types.BankJson{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&bankData)
	if err != nil {
		http.Error(w, "json data error: "+err.Error(), http.StatusInternalServerError)
	}

	// Get template
	tmpl := template.Must(template.ParseFiles("templates/overview/overview.html"))

	// Format bank data struct to pass to template
	// jsonData, _ := json.Marshal(bankData.Summary)
	// fmt.Println(string(jsonData))

	// Serve template with formatted data
	// err = tmpl.Execute(w, string(jsonData))
	err = tmpl.Execute(w, bankData.Summary)

	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}
