package tmplmanager

import (
	"bankmanager/types"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func Account(w http.ResponseWriter, r *http.Request) {

	// Get template
	tmpl := template.Must(template.ParseFiles("templates/account/account.html"))

	err := tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}

func Summary(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	if id != "" {
		fmt.Println("id: ", id)
		http.Error(w, "id error: ", http.StatusInternalServerError)

		// fmt.Printf("ID: %v\n", id)
		// year := id[:4]
		// fmt.Println("ID:", year, id)
	} else {
		fmt.Println("No id param")
	}
	// Get template

	tmpl := template.Must(template.ParseFiles("templates/account/summary.html"))

	err := tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}

func Deposits(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/account/deposits.html"))

	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}

func Withdrawals(w http.ResponseWriter, r *http.Request) {

	// Get template
	tmpl := template.Must(template.ParseFiles("templates/account/withdrawals.html"))

	err := tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}

func Graphs(w http.ResponseWriter, r *http.Request) {

	// id := r.PathValue("id")
	// fmt.Printf("ID: %v\n", id)

	// year := id[:4]

	// statementPath := "data/json/" + year + "/" + id + ".json"
	statementPath := "data/json/2018/20180104.json"
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
	tmpl := template.Must(template.ParseFiles("templates/account/graphs.html"))

	// Format bank data struct to pass to template
	jsonData, _ := json.Marshal(bankData)
	// fmt.Println(string(jsonData))

	// Serve template with formatted data
	err = tmpl.Execute(w, string(jsonData))
	// err := tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}

func Budget(w http.ResponseWriter, r *http.Request) {

	// Get template
	tmpl := template.Must(template.ParseFiles("templates/account/budget.html"))

	err := tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}
