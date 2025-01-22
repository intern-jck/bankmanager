package main

import (
	"fmt"
	"net/http"

	// "os"

	// "bankmanager/parser"
	"bankmanager/tmplmanager"
)

func main() {
	fmt.Println("bank manager")

	// Read all files in the directory
	// dirPath := "data/textpdf/2018"
	// files, err := os.ReadDir(dirPath)
	// if err != nil {
	// 	fmt.Println("read dir error: ", err)
	//  return
	// }

	// for _, file := range files {
	// 	parser.CreateJson("2018/" + file.Name())
	// }

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.HandleFunc("/", tmplmanager.Index)
	http.HandleFunc("/overview/{id}", tmplmanager.Overview)

	fmt.Println("Server @ http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
