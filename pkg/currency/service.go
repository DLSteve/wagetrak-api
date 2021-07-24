package currency

import (
	"fmt"
	"strings"
	"wagetrak-api/pkg/entities"
)

var currencyDesc = map[string]string{
	"USD": "United States Dollar",
	"EUR": "European Union Euro",
	"GBP": "Pound Sterling",
	"JPY": "Japanese Yen",
	"BGN": "Bulgaria Lev",
	"CZK": "Czech Republic Koruna",
	"DKK": "Danish Krone",
	"HUF": "Hungary Forint",
	"PLN": "Polish Zloty",
	"RON": "Romanian Leu",
	"SEK": "Sweden Krona",
	"CHF": "Swiss Franc",
	"ISK": "Icelandic Krona",
	"NOK": "Norwegian Krone",
	"HRK": "Croatian Kuna",
	"RUB": "Russian Ruble",
	"TRY": "Turkish Lira",
	"AUD": "Australian Dollar",
	"BRL": "Brazilian Real",
	"CAD": "Canadian Dollar",
	"CNY": "Chinese Yuan Renminbi",
	"HKD": "Hong Kong Dollar",
	"IDR": "Indonesian Rupiah",
	"ILS": "Israeli Shekel",
	"INR": "India Rupee",
	"KRW": "South Korean Won",
	"MXN": "Mexican Peso",
	"MYR": "Malaysian Ringgit",
	"NZD": "New Zealand Dollar",
	"PHP": "Philippine Peso",
	"SGD": "Singapore Dollar",
	"THB": "Thai Baht",
	"ZAR": "South African Rand",
}

type Service interface {
	GetCurrencyList() ([]entities.Currency, error)
	GetExchangeRates(base string) (entities.Rates, error)
	GetExchangeRate(base, target string) error
}

type service struct {
	client *Client
}

func NewCurrencyService(client *Client) Service {
	return &service{
		client,
	}
}

func (s *service) GetCurrencyList() ([]entities.Currency, error) {
	var curList []entities.Currency

	rates, err := s.client.GetEuroExchangeRates()
	if err != nil {
		return curList, err
	}

	for idx, c := range rates.Currencies {
		// Insert Euro into the list in the second slot.
		if idx == 1 {
			eurDesc := currencyDesc["EUR"]
			curList = append(curList, entities.Currency{
				Code: "EUR",
				Name: eurDesc,
			})
		}
		if desc, ok := currencyDesc[c.Currency]; ok {
			curList = append(curList, entities.Currency{
				Code: c.Currency,
				Name: desc,
			})
		}
	}

	return curList, nil
}

func (s *service) GetExchangeRates(base string) (entities.Rates, error) {
	base = strings.ToUpper(base)
	euroRates, err := s.client.GetEuroExchangeRates()
	if err != nil {
		return entities.Rates{}, err
	}

	rates := entities.Rates{
		Updated: euroRates.Updated,
	}

	curList := make(map[string]entities.Currency)
	var euroExchMap = make(map[string]float32)

	for _, c := range euroRates.Currencies {
		euroExchMap[strings.ToUpper(c.Currency)] = c.Rate
	}

	baseEx, ok := euroExchMap[base]
	if !ok {
		return rates, fmt.Errorf("no exchange rate found for provided base currency")
	}

	for _, c := range euroRates.Currencies {
		if desc, ok := currencyDesc[c.Currency]; ok {
			if "EUR" == base {
				curList[c.Currency] = entities.Currency{
					Name: desc,
					Rate: c.Rate,
				}
			} else if c.Currency == base {
				eurDesc := currencyDesc["EUR"]
				curList["EUR"] = entities.Currency{
					Name: eurDesc,
					Rate: 1 / c.Rate,
				}
			} else {
				curList[c.Currency] = entities.Currency{
					Name: desc,
					Rate: CalculateExchangeFromBase(baseEx, c.Rate),
				}
			}
		}
	}

	rates.Rates = curList

	return rates, nil
}

func (s *service) GetExchangeRate(base, target string) error {
	return nil
}