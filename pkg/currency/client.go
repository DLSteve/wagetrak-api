package currency

import (
	"encoding/xml"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"sync"
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

type Client struct {
	mu          sync.Mutex
	cachedRates *entities.EuroExchangeRates
	Logger      *zap.SugaredLogger
	HttpClient  *fiber.Client
}

func (c *Client) fetchEuroExchangeRates() (entities.EuroExchangeRates, error) {
	if c.HttpClient == nil {
		c.HttpClient = &fiber.Client{}
	}

	resp := c.HttpClient.Get(euroFxUrl)
	stat, body, errs := resp.String()
	if errs != nil {
		return entities.EuroExchangeRates{}, errs[0]
	}

	if stat == 200 {
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
			Updated:       t,
			LastRefreshed: time.Now(),
			Currencies:    envelope.Cube.Cube.Currencies,
		}, nil

	}
	return entities.EuroExchangeRates{}, fmt.Errorf("error while fetching exchange rates")
}

func needsRefresh(updated, lastRefresh, now time.Time) bool {
	if updated.Format("2006-01-02") == now.Format("2006-01-02") {
		return false
	}

	diff := now.Sub(lastRefresh)

	if diff.Hours() < 2 {
		return false
	}

	return true
}

func (c *Client) GetEuroExchangeRates() (entities.EuroExchangeRates, error) {
	// If cache is empty
	if c.cachedRates == nil || needsRefresh(c.cachedRates.Updated, c.cachedRates.LastRefreshed, time.Now()) {
		c.Logger.Info("msg", "Refreshing exchange rates cache")
		rates, err := c.fetchEuroExchangeRates()
		if err != nil {
			return rates, err
		}

		c.mu.Lock()
		c.cachedRates = &rates
		c.mu.Unlock()
	} else {
		c.Logger.Debug("msg", "Exchange rates cache hit")
	}

	return *c.cachedRates, nil
}
