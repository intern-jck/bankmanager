package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// var greetings = []string{"Hello, World!", "Hola, Mundo!", "Bonjour, Monde!", "Hallo, Welt!"}
// var index = 0

func main() {
	fmt.Println("bank manager")

	// http.HandleFunc("/", tmplmanager.IndexTemplate)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte(greetings[index]))
	// 	index = (index + 1) % len(greetings)
	// })

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/test", testHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	// testMsg := "<h2>TEST MESSAGE WORKS!</h2>"
	// fmt.Fprint(w, testMsg)
	tmpl := template.Must(template.ParseFiles("templates/test.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
