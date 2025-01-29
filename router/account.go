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

type TestJson struct {
	ID     string
	Data   []int
	Title  string
	Labels []string
}

type GraphJson struct {
	ID     string
	Data   []float64
	Title  string
	Labels []string
}

func Account(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/account/account.html"))

	err := tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}

func Summary(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	year := r.FormValue("account-year")
	month := r.FormValue("account-month")
	fmt.Println("FORM:", year, month)

	if year != "" && month != "" {
		// get the file
		pattern := "data/json/" + year + "/" + year + month + "*.json"
		fmt.Println(pattern)

		// matches, err := filepath.Glob(pattern)
		// if err != nil {
		// 	fmt.Println("Error:", err)
		// 	// return
		// }

		// if len(matches) == 0 {
		// 	fmt.Println("No matching files found.")
		// 	// return
		// }

		// for _, match := range matches {
		// 	file, err := os.Open(match)
		// 	if err != nil {
		// 		fmt.Println("Error opening file:", err)
		// 		continue
		// 	}
		// 	defer file.Close()

		// 	fmt.Println("FOUND:", match)
		// }
	}

	// statementPath := "data/json/" + year + "/" + id + ".json"
	statementPath := "data/json/2018/20180104.json"

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

	// // Create test data
	// testData := TestJson{
	// 	ID:     "summary-graph",
	// 	Data:   []int{1, 2, 3, 4, 5},
	// 	Title:  "Summary",
	// 	Labels: []string{"one", "two", "three", "four", "five"},
	// }

	// jsonData, _ := json.Marshal(bankData.Summary)
	// fmt.Println(string(jsonData))

	beginning, _ := strconv.ParseFloat(bankData.Summary.Beginning, 64)
	ending, _ := strconv.ParseFloat(bankData.Summary.Ending, 64)
	deposits, _ := strconv.ParseFloat(bankData.Summary.Deposits, 64)
	checks, _ := strconv.ParseFloat(bankData.Summary.Checks, 64)
	debit, _ := strconv.ParseFloat(bankData.Summary.Debit, 64)
	electronic, _ := strconv.ParseFloat(bankData.Summary.Electronic, 64)
	fees, _ := strconv.ParseFloat(bankData.Summary.Fees, 64)

	graphData := GraphJson{
		ID:     "summary-graph",
		Data:   []float64{beginning, ending, deposits, checks, debit, electronic, fees},
		Title:  "Summary",
		Labels: []string{"Beginning", "Ending", "Deposits", "Checks", "Debit", "Electronic", "Fees"},
	}
	jsonData, _ := json.Marshal(graphData)

	err = tmpl.Execute(w, string(jsonData))

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
	// fmt.Println("GETTING: ", statementPath)

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

	tmpl := template.Must(template.ParseFiles("templates/account/graphs.html"))

	// Format bank data struct to pass to template
	// jsonData, _ := json.Marshal(bankData)
	// fmt.Println(string(jsonData))

	// Create test data
	testData := TestJson{
		ID:     "bar-graph",
		Data:   []int{1, 2, 3, 4, 5},
		Title:  "Test Bar Graph",
		Labels: []string{"one", "two", "three", "four", "five"},
	}

	jsonData, _ := json.Marshal(testData)

	// Serve template with formatted data
	err = tmpl.Execute(w, string(jsonData))
	// err := tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}

func Budget(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/account/budget.html"))

	err := tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}
