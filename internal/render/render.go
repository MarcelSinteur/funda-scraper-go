package render

import (
	"net/http"

	"github.com/CloudyKit/jet"
	"github.com/justinas/nosurf"
	"github.com/marcelsinteur/funda-scraper-go/internal/config"
	"github.com/marcelsinteur/funda-scraper-go/internal/models"
)

var app *config.AppConfig

// NewRenderer sets the config for the template package
func NewRenderer(a *config.AppConfig) {
	app = a
}

// AddDefaultData adds data for all templates
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)

	return td
}

// RenderTemplate renders a template
func RenderTemplate(writer http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
	t, err := app.View.GetTemplate(tmpl)
	if err != nil {
		return err
	}

	vars := make(jet.VarMap)
	if err = t.Execute(writer, vars, nil); err != nil {
		return err
	}

	return nil
}
