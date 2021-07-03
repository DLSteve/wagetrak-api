package entities

import (
	"time"
)

type Currency struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type EuroExchangeRate struct {
	Currency string  `xml:"currency,attr"`
	Rate     float32 `xml:"rate,attr"`
}

type EuroExchangeRates struct {
	Updated     time.Time
	LastChecked time.Time
	Currencies  []EuroExchangeRate
}
