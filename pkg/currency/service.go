package currency

import (
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
	GetExchangeRate(base, target string) (entities.Currency, error)
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
	euroExMap := getEuroExchangeMap(euroRates)

	baseEx, ok := euroExMap[base]
	if !ok && base == "EUR" {
		baseEx = 1.0
	} else if !ok {
		return rates, entities.ErrCurrencyNotFound
	}

	for _, c := range euroRates.Currencies {
		if desc, ok := currencyDesc[c.Currency]; ok {
			if base != c.Currency {
				curList[c.Currency] = entities.Currency{
					Name: desc,
					Rate: CalculateExchangeFromBase(baseEx, c.Rate),
				}
			}
		}
	}

	// Adds Euro to the list
	if base != "EUR" {
		if desc, ok := currencyDesc["EUR"]; ok {
			curList["EUR"] = entities.Currency{
				Name: desc,
				Rate: CalculateExchangeFromBase(baseEx, 1.0),
			}
		}
	}

	rates.Rates = curList

	return rates, nil
}

func (s *service) GetExchangeRate(base, target string) (entities.Currency, error) {
	base = strings.ToUpper(base)
	target = strings.ToUpper(target)

	currency := entities.Currency{}

	euroRates, err := s.client.GetEuroExchangeRates()
	if err != nil {
		return currency, err
	}

	currency.Updated = &euroRates.Updated

	euroExMap := getEuroExchangeMap(euroRates)

	baseEx, ok := euroExMap[base]
	if !ok && base == "EUR" {
		baseEx = 1.0
	} else if !ok {
		return currency, entities.ErrCurrencyNotFound
	}

	targetEx, ok := euroExMap[target]
	if !ok && target == "EUR" {
		targetEx = 1.0
	} else if !ok {
		return currency, entities.ErrCurrencyNotFound
	}

	if desc, ok := currencyDesc[target]; ok {
		if base == target {
			currency.Code = target
			currency.Name = desc
			currency.Rate = 1.0
		} else {
			currency.Code = target
			currency.Name = desc
			currency.Rate = CalculateExchangeFromBase(baseEx, targetEx)
		}
	}

	return currency, err
}

func getEuroExchangeMap(rates entities.EuroExchangeRates) map[string]float32 {
	var euroMap = make(map[string]float32)
	for _, c := range rates.Currencies {
		euroMap[strings.ToUpper(c.Currency)] = c.Rate
	}

	return euroMap
}
