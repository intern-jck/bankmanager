package main

import (
	"bankmanager/router"
	"fmt"
	"net/http"
)

func main() {

	// styles and javascript
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// home page
	http.HandleFunc("/", router.Index)

	// account details
	http.HandleFunc("/summary", router.Summary)
	http.HandleFunc("/deposits", router.Deposits)
	http.HandleFunc("/withdrawals", router.Withdrawals)

	// page elements
	http.HandleFunc("/graphs", router.Graphs)

	// server
	fmt.Println("Server @ http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

	// Make fake data tot test
	// faker.CreateStatement(2018, 01)
	// faker.CreateStatement(2018, 02)
	// faker.CreateStatement(2018, 03)
	// faker.CreateStatement(2018, 04)
}
