package entities

import (
	"time"
)

type Currency struct {
	Code    string     `json:"code,omitempty"`
	Name    string     `json:"name"`
	Rate    float32    `json:"rate,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
}

type Rates struct {
	Updated time.Time           `json:"updated"`
	Rates   map[string]Currency `json:"rates"`
}

type EuroExchangeRate struct {
	Currency string  `xml:"currency,attr"`
	Rate     float32 `xml:"rate,attr"`
}

type EuroExchangeRates struct {
	Updated       time.Time
	LastRefreshed time.Time
	Currencies    []EuroExchangeRate
}
