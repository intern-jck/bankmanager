package bankapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"bankmanager/types"
)

type TestJson struct {
	ID     string
	Data   []int
	Title  string
	Labels []string
}

func GetData(w http.ResponseWriter, r *http.Request) {

	// id := r.PathValue("id")
	// fmt.Printf("ID: %v\n", id)

	// year := id[:4]

	// statementPath := "data/json/" + year + "/" + id + ".json"
	statementPath := "data/json/2018/20180104.json"
	fmt.Println("GETTING: ", statementPath)

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	val1 := r.FormValue("value-1")
	val2 := r.FormValue("value-2")
	fmt.Println(val1, val2)

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

	// Get template.Marshal(testData)

	// // Serve template with formatted data
	// err = tmpl.Execute(w, string(jsonData))
	// // err := tmpl.Execute(w, nil)

	// if err != nil {
	// 	http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	// }

	// tmpl := template.Must(template.ParseFiles("templates/account/graphs.html"))

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

	// jsonData, _ := json.Marshal(testData)

	// // Serve template with formatted data
	// err = tmpl.Execute(w, string(jsonData))
	// // err := tmpl.Execute(w, nil)

	// if err != nil {
	// 	http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	// }

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the data to JSON and write it to the response
	json.NewEncoder(w).Encode(testData)
}
