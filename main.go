package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

// Home page handler
func Home(respWriter http.ResponseWriter, req *http.Request) {
	renderTemplate(respWriter, "home.page.tmpl")
}

// About page handler
func About(respWriter http.ResponseWriter, req *http.Request) {

	renderTemplate(respWriter, "about.page.tmpl")
}

func renderTemplate(respWriter http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parseTemplate.Execute(respWriter, nil)
	if err != nil {
		fmt.Println("erro parsing template: ", err)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
