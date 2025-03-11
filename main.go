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
	// http.HandleFunc("/account", router.Account)
	// http.HandleFunc("/settings", router.Settings)
	// http.HandleFunc("/graphs", router.Graphs)
	// http.HandleFunc("/budget", router.Budget)
	// data
	// http.HandleFunc("/data", bankapi.GetData)

	// server
	fmt.Println("Server @ http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

	// faker.CreateStatement(2018, 01, 01)
}
