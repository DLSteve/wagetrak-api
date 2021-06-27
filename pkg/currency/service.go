package currency

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
	GetCurrencyList() error
	GetExchangeRates(base string) error
	GetExchangeRate(base, target string) error
}

type service struct {

}