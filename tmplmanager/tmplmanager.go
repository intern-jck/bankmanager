package tmplmanager

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func IndexTemplate(w http.ResponseWriter, r *http.Request) {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("wd error\n")
	}

	tmpl, err := template.ParseFiles(
		wd+"/public/index.html",
		wd+"/public/styles.css")

	if err != nil {
		fmt.Println("parse file: ", err)
	}

	err = tmpl.Execute(w, "ws://"+r.Host+"/esp")
	if err != nil {
		fmt.Println("execute template:", err.Error())
	}

}
