package currency

import "wagetrak-api/pkg/entities"

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
	GetExchangeRates(base string) error
	GetExchangeRate(base, target string) error
}

type service struct {

}

func NewCurrencyService() Service {
	return &service{}
}

func (s *service) GetCurrencyList() ([]entities.Currency, error) {
	var curList []entities.Currency

	rates, err := GetEuroExchangeRates()
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

func (s *service) GetExchangeRates(base string) error {
	return nil
}

func (s *service) GetExchangeRate(base, target string) error {
	return nil
}