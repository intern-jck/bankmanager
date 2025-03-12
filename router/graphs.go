package router

import (
	"encoding/json"
	"fmt"
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

	var data GraphRequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// get statement
	statementDate := fmt.Sprintf("%s%s01", data.Year, data.Month)
	statementPath := fmt.Sprintf("data/test/%s.json", statementDate)

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
