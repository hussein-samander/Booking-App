package handlers

import (
	"net/http"

	"github.com/hussein-samander/Booking-App/config"
	"github.com/hussein-samander/Booking-App/models"
	"github.com/hussein-samander/Booking-App/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type (wow)
type Repository struct {
	App *config.AppConfig
}

// NewRepository creates a new repository (wow)
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (repos *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "home.page.html", &models.TemplateData{})
	remoteIP := r.RemoteAddr
	repos.App.Session.Put(r.Context(), "remote_ip", remoteIP)
}

func (repos *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["Test"] = "Dan"
	remoteIP := repos.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplates(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (repos *Repository) Bob(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "bob.page.html", &models.TemplateData{})
}
