package config

import (
	"strconv"
	"strings"
)

// Config describes the search parameters and holds all the values to build a valid search
type Config struct {
	BaseUrl          string
	Location         []string
	PriceFrom        int
	PriceTo          int
	PropertyType     string
	Status           Status
	FloorArea        int
	PlotArea         int
	NumberOfRooms    int
	ConstructionType ConstructionType
	ExteriorSpace    []string
}

// GetConfig gets a Config type, populated with default values
func GetConfig() Config {
	var config = Config{
		BaseUrl:          "https://www.funda.nl/koop/",      //Base URL for the Funda website
		Location:         []string{"voorschoten", "leiden"}, //Can have multiple locations
		PriceFrom:        350000,                            //Minimum price
		PriceTo:          450000,                            //Maximum price
		PropertyType:     "woonhuis",                        //House or apartment
		Status:           Available,                         //Options are 'Available', 'UnderNegotation' or 'Sold'
		FloorArea:        100,                               //Minimum floor area
		NumberOfRooms:    4,                                 //Total number of rooms including bedrooms
		ConstructionType: Resale,                            //Options are 'Resale' or 'New'
		ExteriorSpace:    []string{"tuin"},                  //Options are 'tuin' (garden), 'balkon' (balcony) or 'dakterras' (roof terrace)
	}

	return config
}

// Status describes the status of houses that needs to be filtered on
type Status int64

const (
	Available Status = iota
	Sold
	UnderNegotation
)

// ToString converts the Status enum to a valid string which is accepted by the Funda website
func (s Status) ToString() string {
	switch s {
	case Available:
		return "beschikbaar"
	case Sold:
		return "verkocht"
	case UnderNegotation:
		return "in-onderhandeling"
	}

	return "beschikbaar"
}

// ConstructionType describes the type of construction for houses that needs to be filtered on
type ConstructionType int64

const (
	Resale ConstructionType = iota
	New
)

// ToString converts the ConstructionType enum to a valid string which is accepted by the Funda website
func (s ConstructionType) ToString() string {
	switch s {
	case Resale:
		return "bestaande-bouw"
	case New:
		return "nieuwbouw"
	}

	return "bestaande-bouw"
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
	urlString += config.Status.ToString() + "/"
	urlString += config.GetPriceString() + "/"
	urlString += config.GetFloorAreaString() + "/"
	urlString += config.PropertyType + "/"

	if config.PlotArea != 0 {
		urlString += config.GetPlotAreaString() + "/"
	}

	urlString += config.GetNumberOfRoomsString() + "/"
	urlString += config.ConstructionType.ToString() + "/"
	urlString += config.GetExteriorSpaceString()

	return urlString
}
