package main

import (
	"fmt"
	"net/http"

	"bankmanager/bankapi"
	"bankmanager/router"
)

func main() {

	fmt.Println("bank manager")

	// styles and javascript
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// home page
	http.HandleFunc("/", router.Index)

	// page elements
	http.HandleFunc("/account", router.Account)
	http.HandleFunc("/settings", router.Settings)

	// account elements
	http.HandleFunc("/summary", router.Summary)
	http.HandleFunc("/deposits", router.Deposits)
	http.HandleFunc("/withdrawals", router.Withdrawals)
	http.HandleFunc("/graphs", router.Graphs)
	http.HandleFunc("/budget", router.Budget)

	// data
	http.HandleFunc("/data", bankapi.GetData)

	// server
	fmt.Println("Server @ http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
