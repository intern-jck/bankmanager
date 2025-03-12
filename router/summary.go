package router

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"bankmanager/types"
)

func Summary(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// fmt.Println("form: ", r.Form)

	// default path
	// statementPath := "data/json/2018/20180104.json"
	// use fake data
	statementPath := "faker/fakedata/20180101.json"

	// get json data
	file, err := os.Open(statementPath)
	if err != nil {
		fmt.Println("json open error", err)
		http.Error(w, fmt.Sprintf("Error opening data: %v", err), http.StatusInternalServerError)
	}
	defer file.Close()

	tmpl := template.Must(template.ParseFiles("templates/account/summary.html", "templates/graphs/summary-graph.html"))

	// Create bank data struct
	bankData := types.BankJson{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&bankData)
	if err != nil {
		http.Error(w, "json data error: "+err.Error(), http.StatusInternalServerError)
	}

	beginning, _ := strconv.ParseFloat(bankData.Summary.Beginning, 64)
	ending, _ := strconv.ParseFloat(bankData.Summary.Ending, 64)
	deposits, _ := strconv.ParseFloat(bankData.Summary.Deposits, 64)
	withdrawals, _ := strconv.ParseFloat(bankData.Summary.Withdrawals, 64)

	graphData := types.GraphJson{
		ID:     "summary-graph",
		Data:   []float64{beginning, ending, deposits, withdrawals},
		Title:  "Summary",
		Labels: []string{"Beginning", "Ending", "Deposits", "Withdrawals"},
	}
	jsonData, _ := json.Marshal(graphData)

	err = tmpl.Execute(w, string(jsonData))

	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}
