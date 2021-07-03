package v1

import (
	"github.com/gofiber/fiber/v2"
	"strings"

	"wagetrak-api/pkg/currency"
)

func CurrencyRouterV1(app fiber.Router, service currency.Service) {
	app.Get("/exchange", getExchangeList(service))
	app.Get("/exchange/:base", getExchangeRates())
	app.Get("/exchange/:base/:target", getExchangeRate())
}

func getExchangeList(service currency.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		curList, err := service.GetCurrencyList()
		if err != nil {
			return err
		}

		return c.JSON(curList)
	}
}

func getExchangeRates() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"base": strings.ToUpper(c.Params("base")),
		})
	}
}

func getExchangeRate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"base": strings.ToUpper(c.Params("base")),
			"target": strings.ToUpper(c.Params("target")),
		})
	}
}