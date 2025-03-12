package main

import (
	"bankmanager/router"
	"fmt"
	"net/http"
)

func main() {

	// styles and javascript
	fs := http.FileServer(http.Dir("public"))
	handler := http.StripPrefix("/public/", fs)
	http.Handle("/public/", handler)

	// home page
	http.HandleFunc("/", router.Index)

	// account details
	http.HandleFunc("/summary", router.Summary)

	// page elements
	http.HandleFunc("/graphs", router.GetGraphData)

	// server
	fmt.Println("Server @ http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
