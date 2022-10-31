package handlers

import (
	"net/http"

	"github.com/marcelsinteur/funda-scraper-go/internal/config"
	"github.com/marcelsinteur/funda-scraper-go/internal/models"
	"github.com/marcelsinteur/funda-scraper-go/internal/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.jet", &models.TemplateData{})
}
