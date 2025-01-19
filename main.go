package main

import (
	"bankmanager/parser"
	"fmt"
	// "net/http"
	// "bankmanager/tmplmanager"
)

func main() {
	fmt.Println("bank manager")
	// http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// http.HandleFunc("/", tmplmanager.Index)
	// http.HandleFunc("/overview", tmplmanager.Overview)

	// http.ListenAndServe(":8080", nil)

	parser.CreateJson("2018/20180104.txt")
}
