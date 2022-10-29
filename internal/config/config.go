package config

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

// Config describes the search parameters and holds all the values to build a valid search
type Config struct {
	BaseUrl          string
	Location         []string `json:"location"`
	PriceFrom        int      `json:"price_from"`
	PriceTo          int      `json:"price_to"`
	PropertyType     string   `json:"property_type"`
	Status           string   `json:"status"`
	FloorArea        int      `json:"floor_area"`
	PlotArea         int      `json:"plot_area"`
	NumberOfRooms    int      `json:"number_of_rooms"`
	ConstructionType string   `json:"construction_type"`
	ExteriorSpace    []string `json:"exterior_space"`
}

// GetConfig gets a Config type, populated with default values
func GetConfig() (Config, error) {
	configJson, err := os.ReadFile("./config.json")

	if err != nil {
		return Config{}, err
	}

	config, err := parseToConfig(configJson)
	if err != nil {
		return Config{}, err
	}

	baseUrl := "https://www.funda.nl/koop/"
	config.BaseUrl = baseUrl

	return config, nil
}

// parseToConfig parses the config.json file to a Config type
func parseToConfig(j []byte) (Config, error) {
	var config Config
	err := json.Unmarshal(j, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

// GetLocationString returns a location string formatted in the way the Funda website expects
func (config *Config) GetLocationString() string {
	locationString := ""
	for _, location := range config.Location {
		locationString += location + ","
	}

	locationString = strings.TrimRight(locationString, ",")
	return locationString
}

// GetPriceString returns a price string formatted in the way the Funda website expects
func (config *Config) GetPriceString() string {
	priceString := strconv.Itoa(config.PriceFrom) + "-" + strconv.Itoa(config.PriceTo)
	return priceString
}

// GetExteriorSpaceString returns an exterior space string formatted in the way the Funda website expects
func (config *Config) GetExteriorSpaceString() string {
	exteriorSpaceString := ""
	for _, exteriorSpace := range config.ExteriorSpace {
		exteriorSpaceString += exteriorSpace + "/"
	}

	exteriorSpaceString = strings.TrimRight(exteriorSpaceString, "/")
	return exteriorSpaceString
}

// GetNumberOfRoomsString returns a number of rooms string formatted in the way the Funda website expects
func (config *Config) GetNumberOfRoomsString() string {
	numberOfRoomsString := strconv.Itoa(config.NumberOfRooms) + "+kamers"
	return numberOfRoomsString
}

// GetFloorAreaString returns a floor area string formatted in the way the Funda website expects
func (config *Config) GetFloorAreaString() string {
	floorAreaString := strconv.Itoa(config.FloorArea) + "+woonopp"
	return floorAreaString
}

// GetPlotAreaString returns a plot area string formatted in the way the Funda website expects
func (config *Config) GetPlotAreaString() string {
	plotAreaString := strconv.Itoa(config.PlotArea) + "+perceelopp"
	return plotAreaString
}

// BuildUrlString builds a valid URL in the way the Funda website expects, to make the actual call
func (config *Config) BuildUrlString() string {
	urlString := config.BaseUrl
	urlString += config.GetLocationString() + "/"
	urlString += config.Status + "/"
	urlString += config.GetPriceString() + "/"

	if config.PlotArea != 0 {
		urlString += config.GetFloorAreaString() + "/"
	}

	urlString += config.PropertyType + "/"

	if config.PlotArea != 0 {
		urlString += config.GetPlotAreaString() + "/"
	}

	urlString += config.GetNumberOfRoomsString() + "/"
	urlString += config.ConstructionType + "/"
	urlString += config.GetExteriorSpaceString()

	return urlString
}
