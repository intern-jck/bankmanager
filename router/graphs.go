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

type GraphRequestData struct {
	Year  string `json:"year"`
	Month string `json:"month"`
}

func GetGraphData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting graph data")

	// err := r.ParseForm()
	// if err != nil {
	// 	http.Error(w, "Error parsing form data", http.StatusBadRequest)
	// 	return
	// }

	// get form data
	// year := r.Form.Get("account-year")
	// month := r.Form.Get("account-month")

	// if year == "" {
	// 	year = "2018"
	// }

	// if month == "" {
	// 	month = "01"
	// }
	var data GraphRequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("data: ", data.Year, data.Month)

	// get statement
	statementDate := fmt.Sprintf("%s%s01", data.Year, data.Month)
	statementPath := fmt.Sprintf("data/test/%s.json", statementDate)

	// fmt.Println("opening: ", year, month)

	// get json from path
	file, err := os.Open(statementPath)
	if err != nil {
		http.Error(w, "Statement not found: "+err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()

	// parse json file into struct
	bankData := types.BankJson{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&bankData)
	if err != nil {
		http.Error(w, "Cannot read statment: "+err.Error(), http.StatusInternalServerError)
	}

	// create data set for graph
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
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func Graphs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting graphs")
	// get the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// get the statement
	statementDate := fmt.Sprintf("%s%s01", r.Form.Get("account-year"), r.Form.Get("account-month"))
	statementPath := fmt.Sprintf("data/test/%s.json", statementDate)

	// get the graph type
	graphType := r.Form.Get("date-select-id")

	// get the graph
	switch graphType {
	case "summary":
		summaryGraph(statementPath, w, r)
	case "deposits":
	case "withdrawals":
	default:
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}

func summaryGraph(s string, w http.ResponseWriter, r *http.Request) {

	fmt.Println("statement: ", s)

	// get json from path
	file, err := os.Open(s)
	if err != nil {
		http.Error(w, "Statement not found: "+err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()

	// get the summary graph
	tmpl := template.Must(template.ParseFiles("templates/graphs/summary-graph.html"))

	// parse json file into struct
	bankData := types.BankJson{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&bankData)
	if err != nil {
		http.Error(w, "Cannot read statment: "+err.Error(), http.StatusInternalServerError)
	}

	// create data set for graph
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

	// pass data to graph html and return
	err = tmpl.ExecuteTemplate(w, "summary-graph", string(jsonData))

	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
	}
}
