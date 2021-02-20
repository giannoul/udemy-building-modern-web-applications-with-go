package handlers

import (
	"net/http"

	"github.com/giannoul/udemy-building-modern-web-applications-with-go/pkg/render"
)

// Home page handler
func Home(respWriter http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(respWriter, "home.page.tmpl")
}

// About page handler
func About(respWriter http.ResponseWriter, req *http.Request) {

	render.RenderTemplate(respWriter, "about.page.tmpl")
}
