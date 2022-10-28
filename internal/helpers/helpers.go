package helpers

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/marcelsinteur/funda-scraper-go/internal/models"
)

// ParseHtml parses the html from the Funda website and converts it to 'Property' objects
func ParseFundaHtml(r http.Response) ([]models.Property, error) {
	doc, err := goquery.NewDocumentFromReader(r.Body)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var properties []models.Property

	doc.Find(".search-result").Each(func(i int, s *goquery.Selection) {
		address := CleanString(s.Find(".search-result__header-title").Text())

		trimmedPostalCodeAndCity := CleanString(s.Find(".search-result__header-subtitle").Text())
		postalCode := strings.Split(trimmedPostalCodeAndCity, " ")[0] + strings.Split(trimmedPostalCodeAndCity, " ")[1]
		city := strings.Split(trimmedPostalCodeAndCity, " ")[2]

		price := CleanString(s.Find(".search-result-price").Text())
		floorArea := CleanString(s.Find("span[title='Gebruiksoppervlakte wonen']").Text())
		plotArea := CleanString(s.Find("span[title='Perceeloppervlakte']").Text())

		attributesNode := s.Find(".search-result-kenmerken ")
		numberOfRooms := CleanString(attributesNode.Children().Last().Text())

		properties = append(properties, models.Property{
			Address:       address,
			PostalCode:    postalCode,
			City:          city,
			Price:         price,
			FloorArea:     floorArea,
			PlotArea:      plotArea,
			NumberOfRooms: numberOfRooms,
		})
	})

	return properties, nil
}

// CleanString cleans a given string of any spaces and newline characters
func CleanString(v string) string {
	v = strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(v), "\r", ""), "\n", "")
	return v
}
