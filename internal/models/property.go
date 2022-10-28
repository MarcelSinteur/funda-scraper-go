package models

import "encoding/json"

type Property struct {
	Address       string
	PostalCode    string
	City          string
	Price         string
	FloorArea     string
	PlotArea      string
	NumberOfRooms string
}

func ParsePropertiesToJson(props []Property) ([]byte, error) {
	json, err := json.Marshal(props)
	if err != nil {
		return nil, err
	}

	return json, nil
}
