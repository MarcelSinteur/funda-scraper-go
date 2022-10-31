package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/marcelsinteur/funda-scraper-go/internal/config"
	"github.com/marcelsinteur/funda-scraper-go/internal/forms"
	"github.com/marcelsinteur/funda-scraper-go/internal/helpers"
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

// Search starts a Goroutine that periodically queries the Funda website
func (m *Repository) Search(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	form := forms.New(r.PostForm)
	form.IsNumber("price_from", "price_to", "floor_area", "plot_area", "number_of_rooms")

	search := models.Search{
		BaseUrl:          m.App.BaseUrl,
		Location:         strings.Split(r.Form.Get("location"), ","),
		PriceFrom:        r.Form.Get("price_from"),
		PriceTo:          r.Form.Get("price_to"),
		PropertyType:     r.Form.Get("property_type"),
		Status:           r.Form.Get("status"),
		FloorArea:        r.Form.Get("floor_area"),
		PlotArea:         r.Form.Get("plot_area"),
		NumberOfRooms:    r.Form.Get("number_of_rooms"),
		ConstructionType: r.Form.Get("construction_type"),
		ExteriorSpace:    strings.Split(r.Form.Get("exterior_space"), ","),
	}

	if !form.IsValid() {
		data := make(map[string]interface{})
		data["search"] = search

		render.RenderTemplate(w, r, "home.jet", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	url := search.BuildUrlString()

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"User-Agent":   {"PostmanRuntime/7.28.4"},
		"Connection":   {"keep-alive"},
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	properties, err := helpers.ParseFundaHtml(*res)
	if err != nil {
		log.Fatal(err)
	}

	json, err := models.SerializePropertiesToJson(properties)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
