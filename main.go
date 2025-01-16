package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("bank manager")

	// http.HandleFunc("/", tmplmanager.IndexTemplate)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":8080", nil)
}
