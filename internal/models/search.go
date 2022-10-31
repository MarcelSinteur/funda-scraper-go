package models

import (
	"strings"
)

// Config describes the search parameters and holds all the values to build a valid search
type Search struct {
	BaseUrl          string
	Location         []string
	PriceFrom        string
	PriceTo          string
	PropertyType     string
	Status           string
	FloorArea        string
	PlotArea         string
	NumberOfRooms    string
	ConstructionType string
	ExteriorSpace    []string
}

// GetLocationString returns a location string formatted in the way the Funda website expects
func (search *Search) GetLocationString() string {
	locationString := ""
	for _, location := range search.Location {
		locationString += location + ","
	}

	locationString = strings.TrimRight(locationString, ",")
	return locationString
}

// GetPriceString returns a price string formatted in the way the Funda website expects
func (search *Search) GetPriceString() string {
	priceString := search.PriceFrom + "-" + search.PriceTo
	return priceString
}

// GetExteriorSpaceString returns an exterior space string formatted in the way the Funda website expects
func (search *Search) GetExteriorSpaceString() string {
	exteriorSpaceString := ""
	for _, exteriorSpace := range search.ExteriorSpace {
		exteriorSpaceString += exteriorSpace + "/"
	}

	exteriorSpaceString = strings.TrimRight(exteriorSpaceString, "/")
	return exteriorSpaceString
}

// GetNumberOfRoomsString returns a number of rooms string formatted in the way the Funda website expects
func (search *Search) GetNumberOfRoomsString() string {
	numberOfRoomsString := search.NumberOfRooms + "+kamers"
	return numberOfRoomsString
}

// GetFloorAreaString returns a floor area string formatted in the way the Funda website expects
func (search *Search) GetFloorAreaString() string {
	floorAreaString := search.FloorArea + "+woonopp"
	return floorAreaString
}

// GetPlotAreaString returns a plot area string formatted in the way the Funda website expects
func (search *Search) GetPlotAreaString() string {
	plotAreaString := search.PlotArea + "+perceelopp"
	return plotAreaString
}

// BuildUrlString builds a valid URL in the way the Funda website expects, to make the actual call
func (search *Search) BuildUrlString() string {
	urlString := search.BaseUrl
	urlString += search.GetLocationString() + "/"
	urlString += search.Status + "/"
	urlString += search.GetPriceString() + "/"

	if search.PlotArea != "0" {
		urlString += search.GetFloorAreaString() + "/"
	}

	urlString += search.PropertyType + "/"

	if search.PlotArea != "0" {
		urlString += search.GetPlotAreaString() + "/"
	}

	urlString += search.GetNumberOfRoomsString() + "/"
	urlString += search.ConstructionType + "/"
	urlString += search.GetExteriorSpaceString()

	return urlString
}
