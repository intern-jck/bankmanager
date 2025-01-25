package main

import (
	"fmt"
	"net/http"

	"bankmanager/tmplmanager"
)

func main() {

	fmt.Println("bank manager")

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.HandleFunc("/", tmplmanager.Index)

	http.HandleFunc("/account", tmplmanager.Account)
	http.HandleFunc("/settings", tmplmanager.Settings)

	http.HandleFunc("/summary", tmplmanager.Summary)
	http.HandleFunc("/deposits", tmplmanager.Deposits)
	http.HandleFunc("/withdrawals", tmplmanager.Withdrawals)
	http.HandleFunc("/graphs", tmplmanager.Graphs)
	http.HandleFunc("/budget", tmplmanager.Budget)

	fmt.Println("Server @ http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
