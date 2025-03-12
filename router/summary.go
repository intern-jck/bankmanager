package router

import (
	"fmt"
	"html/template"
	"net/http"
)

func Summary(w http.ResponseWriter, r *http.Request) {

	fmt.Println("getting summary")

	// // parse form data
	// err := r.ParseForm()
	// if err != nil {
	// 	http.Error(w, "Error parsing form data", http.StatusBadRequest)
	// 	return
	// }

	// // get form data
	// year := r.Form.Get("account-year")
	// month := r.Form.Get("account-month")

	// get graph type
	// graphType := r.Form.Get("date-select-id")

	// fmt.Println(graphType)

	// if year == "" {
	// 	year = "2018"
	// }

	// if month == "" {
	// 	month = "01"
	// }

	// // get statement
	// statementDate := fmt.Sprintf("%s%s01", year, month)
	// statementPath := fmt.Sprintf("data/test/%s.json", statementDate)

	// fmt.Println(statementPath, graphType)

	tmpl := template.Must(template.ParseFiles("templates/summary.html"))

	// get json from path
	// file, err := os.Open(statementPath)
	// if err != nil {
	// 	http.Error(w, "Statement not found: "+err.Error(), http.StatusInternalServerError)
	// }
	// defer file.Close()

	// // parse json file into struct
	// bankData := types.BankJson{}
	// decoder := json.NewDecoder(file)
	// err = decoder.Decode(&bankData)
	// if err != nil {
	// 	http.Error(w, "Cannot read statment: "+err.Error(), http.StatusInternalServerError)
	// }

	// // create data set for graph
	// beginning, _ := strconv.ParseFloat(bankData.Summary.Beginning, 64)
	// ending, _ := strconv.ParseFloat(bankData.Summary.Ending, 64)
	// deposits, _ := strconv.ParseFloat(bankData.Summary.Deposits, 64)
	// withdrawals, _ := strconv.ParseFloat(bankData.Summary.Withdrawals, 64)

	// graphData := types.GraphJson{
	// 	ID:     "summary-graph",
	// 	Data:   []float64{beginning, ending, deposits, withdrawals},
	// 	Title:  "Summary",
	// 	Labels: []string{"Beginning", "Ending", "Deposits", "Withdrawals"},
	// }
	// jsonData, _ := json.Marshal(graphData)

	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}

}
