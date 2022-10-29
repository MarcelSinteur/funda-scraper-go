package helpers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

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

		price, err := GetNumberFromString(CleanString(s.Find(".search-result-price").Text()))
		if err != nil {
			log.Fatal(err)
			return
		}

		floorArea, err := GetNumberFromString(CleanString(s.Find("span[title='Gebruiksoppervlakte wonen']").Text()))
		if err != nil {
			log.Fatal(err)
			return
		}

		plotArea, err := GetNumberFromString(CleanString(s.Find("span[title='Perceeloppervlakte']").Text()))
		if err != nil {
			log.Fatal(err)
			return
		}

		attributesNode := s.Find(".search-result-kenmerken ")
		numberOfRooms, err := GetNumberFromString(CleanString(attributesNode.Children().Last().Text()))
		if err != nil {
			log.Fatal(err)
			return
		}

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

// GetNumberFromString retrieves the number inside a string
func GetNumberFromString(v string) (int, error) {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	v = re.FindString(v)

	v = strings.ReplaceAll(v, ".", "")
	v = strings.ReplaceAll(v, ",", "")

	if v == "" {
		return 0, nil
	}

	number, err := strconv.Atoi(v)

	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return number, nil
}

// GenerateJsonFiles generates a json file
func GenerateJsonFiles(json []byte) error {
	if _, err := os.Stat("./output"); os.IsNotExist(err) {
		err = os.Mkdir("./output", 0755)
		if err != nil {
			return err
		}
	}
	year, month, day := time.Now().Date()
	err := os.WriteFile(fmt.Sprintf("./output/results%d-%d-%d.json", year, month, day), json, 0644)
	if err != nil {
		return err
	}

	return nil
}
