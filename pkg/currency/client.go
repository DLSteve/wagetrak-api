package currency

import (
	"encoding/xml"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
	"wagetrak-api/pkg/entities"
)

const euroFxUrl = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

type envelope struct {
	Cube outerCube `xml:"Cube"`
}

type outerCube struct {
	Cube innerCube `xml:"Cube"`
}

type innerCube struct {
	Time       string                      `xml:"time,attr"`
	Currencies []entities.EuroExchangeRate `xml:"Cube"`
}

func GetEuroExchangeRates() (entities.EuroExchangeRates, error) {
	client := fiber.Client{}
	resp := client.Get(euroFxUrl)
	stat, body, errs := resp.String()
	if errs != nil {
		return entities.EuroExchangeRates{}, errs[0]
	}

	if stat == 200 {
		//log.Println(body)
		var envelope envelope
		err := xml.Unmarshal([]byte(body), &envelope)
		if err != nil {
			return entities.EuroExchangeRates{}, err
		}

		layout := "2006-07-02"
		t, err := time.Parse(layout, envelope.Cube.Cube.Time)
		if err != nil {
			return entities.EuroExchangeRates{}, err
		}

		return entities.EuroExchangeRates{
			Updated: t,
			LastChecked: time.Now(),
			Currencies: envelope.Cube.Cube.Currencies,
		}, nil

	}
	return entities.EuroExchangeRates{}, fmt.Errorf("error while fetching exchange rates")
}
