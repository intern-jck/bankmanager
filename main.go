package main

import (
	"fmt"
	"net/http"

	"bankmanager/tmplmanager"
)

func main() {

	fmt.Println("bank manager")

	// styles and javascript
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// home page
	http.HandleFunc("/", tmplmanager.Index)

	// page elements
	http.HandleFunc("/account", tmplmanager.Account)
	http.HandleFunc("/settings", tmplmanager.Settings)

	// account elements
	// http.HandleFunc("/summary", tmplmanager.Summary)
	http.HandleFunc("/summary/{id}", tmplmanager.Summary)
	http.HandleFunc("/deposits", tmplmanager.Deposits)
	http.HandleFunc("/withdrawals", tmplmanager.Withdrawals)
	http.HandleFunc("/graphs", tmplmanager.Graphs)
	http.HandleFunc("/budget", tmplmanager.Budget)

	// server
	fmt.Println("Server @ http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
