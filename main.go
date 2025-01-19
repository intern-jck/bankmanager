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
	http.HandleFunc("/overview", tmplmanager.Overview)

	http.ListenAndServe(":8080", nil)

	// Read all files in the directory
	// dirPath := "data/textpdf/2018"
	// files, err := os.ReadDir(dirPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, file := range files {

	// 	parser.CreateJson("2018/" + file.Name())

	// }
}
