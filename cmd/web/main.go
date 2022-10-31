package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/CloudyKit/jet"
	"github.com/alexedwards/scs/v2"
	"github.com/marcelsinteur/funda-scraper-go/internal/config"
	"github.com/marcelsinteur/funda-scraper-go/internal/handlers"
	"github.com/marcelsinteur/funda-scraper-go/internal/helpers"
	"github.com/marcelsinteur/funda-scraper-go/internal/models"
	"github.com/marcelsinteur/funda-scraper-go/internal/render"
)

const ipAddress = "127.0.0.1:8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	err := initialize()
	if err != nil {
		log.Fatal(err)
	}

	serve := &http.Server{
		Addr:    ipAddress,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server")
	}

	// configuration, err := config.GetConfig()
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// url := configuration.BuildUrlString()

	// client := http.Client{}
	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// req.Header = http.Header{
	// 	"Content-Type": {"application/json"},
	// 	"User-Agent":   {"PostmanRuntime/7.28.4"},
	// 	"Connection":   {"keep-alive"},
	// }

	// for range time.Tick(time.Hour * 4) {
	// 	json, err := fetchData(client, req)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		return
	// 	}

	// 	helpers.GenerateJsonFiles(json)
	// }
}

func fetchData(client http.Client, req *http.Request) ([]byte, error) {
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	properties, err := helpers.ParseFundaHtml(*res)
	if err != nil {
		return nil, err
	}

	json, err := models.SerializePropertiesToJson(properties)

	if err != nil {
		return nil, err
	}
	return json, nil
}

func initialize() error {

	//Register structs that are going to be stored in the session
	gob.Register(models.Property{})

	//Change this to true when in production
	app.InProduction = false
	app.View = jet.NewHTMLSet("./views")

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewRenderer(&app)

	return nil
}
