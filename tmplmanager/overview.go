package tmplmanager

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"bankmanager/types"
)

type test struct {
	Name   string
	Value  int
	Floats []float32
}

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
	if err != nil {
		fmt.Println(err)
		http.Error(w, "json data error: "+err.Error(), http.StatusInternalServerError)
	}

	// get template
	tmpl := template.Must(template.ParseFiles("templates/overview.html"))

	data := test{
		Name:   "test",
		Value:  123,
		Floats: []float32{123.4, 567.8, 901.2},
	}

	// serve template
	err = tmpl.ExecuteTemplate(w, "overview", data)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}
