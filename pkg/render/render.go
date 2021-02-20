package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// RenderTemplate renders templates using text/templates
func RenderTemplate(respWriter http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parseTemplate.Execute(respWriter, nil)
	if err != nil {
		fmt.Println("erro parsing template: ", err)
		return
	}
}
