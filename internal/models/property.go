package models

import "encoding/json"

type Property struct {
	Address       string
	PostalCode    string
	City          string
	Price         int
	FloorArea     int
	PlotArea      int
	NumberOfRooms int
}

//SerializePropertiesToJson serializes a slice of 'Property' to JSON
func SerializePropertiesToJson(props []Property) ([]byte, error) {
	json, err := json.MarshalIndent(props, "", "   ")
	if err != nil {
		return nil, err
	}

	return json, nil
}
