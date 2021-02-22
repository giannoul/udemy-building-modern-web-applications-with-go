package handlers

import (
	"net/http"

	"github.com/giannoul/udemy-building-modern-web-applications-with-go/pkg/config"
	"github.com/giannoul/udemy-building-modern-web-applications-with-go/pkg/render"
)

// Repository pattern
type Repository struct {
	App *config.AppConfig
}

// Repo is the Repository used by the handlers
var Repo *Repository

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the Repository for new handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page handler
func (m *Repository) Home(respWriter http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(respWriter, "home.page.tmpl")
}

// About page handler
func (m *Repository) About(respWriter http.ResponseWriter, req *http.Request) {

	render.RenderTemplate(respWriter, "about.page.tmpl")
}
