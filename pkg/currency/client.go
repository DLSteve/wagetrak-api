package currency

import (
	"encoding/xml"
	"github.com/gofiber/fiber/v2"
	"log"
)

const euroFxUrl = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

type envelope struct {
	Cube outerCube `xml:"Cube"`
}

type outerCube struct {
	Cube innerCube `xml:"Cube"`
}

type innerCube struct {
	Time       string     `xml:"time,attr"`
	Currencies []currency `xml:"Cube"`
}

type currency struct {
	Currency string  `xml:"currency,attr"`
	Rate     float32 `xml:"rate,attr"`
}

func GetExchangeRates() error {
	client := fiber.Client{}
	resp := client.Get(euroFxUrl)
	stat, body, errs := resp.String()
	if errs != nil {
		return errs[0]
	}

	if stat == 200 {
		//log.Println(body)
		var envelope envelope
		err := xml.Unmarshal([]byte(body), &envelope)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println(envelope.Cube.Cube.Currencies[0])
	}
	return nil
}
