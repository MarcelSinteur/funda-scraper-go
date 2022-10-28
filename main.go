package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/marcelsinteur/funda-scraper-go/internal/config"
	"github.com/marcelsinteur/funda-scraper-go/internal/helpers"
	"github.com/marcelsinteur/funda-scraper-go/internal/models"
)

var configuration = config.GetConfig()

func main() {
	url := configuration.BuildUrlString()

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

	for range time.Tick(time.Second * 5) {
		json, err := fetchData(client, req)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(json)
	}
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

	json, err := parseToJson(properties)

	if err != nil {
		return nil, err
	}
	return json, nil
}

func parseToJson(properties []models.Property) ([]byte, error) {
	json, err := json.Marshal(properties)
	if err != nil {
		return nil, err
	}

	return json, nil
}
