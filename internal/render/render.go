package render

import (
	"html/template"
	"log"
	"net/http"

	"github.com/CloudyKit/jet"
	"github.com/justinas/nosurf"
	"github.com/marcelsinteur/funda-scraper-go/internal/config"
	"github.com/marcelsinteur/funda-scraper-go/internal/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig
var templatePath = "./templates"

// NewRenderer sets the config for the template package
func NewRenderer(a *config.AppConfig) {
	app = a
}

// AddDefaultData adds data for all templates
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)

	return td
}

// RenderTemplate renders a template
func RenderTemplate(writer http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
	// var tc map[string]*template.Template

	// if app.UseCache {
	// 	// get the template cache from the app config
	// 	tc = app.TemplateCache
	// } else {
	// 	tc, _ = CreateTemplateCache()
	// }

	// t, ok := tc[tmpl]
	// if !ok {
	// 	return errors.New("Can't get template from cache")
	// }

	// buf := new(bytes.Buffer)

	// td = AddDefaultData(td, r)

	// _ = t.Execute(buf, td)

	// _, err := buf.WriteTo(w)
	// if err != nil {
	// 	fmt.Println("error writing template to browser", err)
	// 	return err
	// }

	// return nil

	t, err := app.View.GetTemplate(tmpl)
	if err != nil {
		log.Println(t.Name)
		log.Println(err)
		return err
	}

	vars := make(jet.VarMap)
	if err = t.Execute(writer, vars, nil); err != nil {
		return err
	}

	return nil
}
