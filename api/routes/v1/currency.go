package v1

import (
	"github.com/gofiber/fiber/v2"
	"strings"

	"wagetrak-api/pkg/currency"
)

func CurrencyRouterV1(app fiber.Router) {
	app.Get("/exchange", getExchangeList())
	app.Get("/exchange/:base", getExchangeRates())
	app.Get("/exchange/:base/:target", getExchangeRate())
}

func getExchangeList() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_ = currency.GetExchangeRates()
		return c.JSON(&fiber.Map{
			"base": strings.ToUpper(c.Params("base")),
		})
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