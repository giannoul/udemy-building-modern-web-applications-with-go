package handlers

import (
	"net/http"

	"github.com/giannoul/udemy-building-modern-web-applications-with-go/pkg/config"
	"github.com/giannoul/udemy-building-modern-web-applications-with-go/pkg/models"
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
	remoteIP := req.RemoteAddr
	m.App.Session.Put(req.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(respWriter, "home.page.tmpl", &models.TemplateData{})
}

// About page handler
func (m *Repository) About(respWriter http.ResponseWriter, req *http.Request) {

	// business logic
	remoteIP := m.App.Session.GetString(req.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["test"] = "hello again"
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(respWriter, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
