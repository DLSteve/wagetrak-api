package v1

import (
	"github.com/gofiber/fiber/v2"
	"wagetrak-api/pkg/currency"
	"wagetrak-api/pkg/entities"
)

func CurrencyRouterV1(app fiber.Router, service currency.Service) {
	app.Get("/exchange", getExchangeList(service))
	app.Get("/exchange/:base", getExchangeRates(service))
	app.Get("/exchange/:base/:target", getExchangeRate(service))
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

func getExchangeRates(service currency.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		base := c.Params("base")
		rates, err := service.GetExchangeRates(base)
		if err != nil {
			if err == entities.ErrCurrencyNotFound {
				return fiber.NewError(fiber.StatusNotFound, "Provided currency code is not valid or not supported.")
			}
			return err
		}
		return c.JSON(rates)
	}
}

func getExchangeRate(service currency.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		base := c.Params("base")
		target := c.Params("target")

		rate, err := service.GetExchangeRate(base, target)
		if err != nil {
			if err == entities.ErrCurrencyNotFound {
				return fiber.NewError(fiber.StatusNotFound, "Provided currency code is not valid or not supported.")
			}
			return err
		}
		return c.JSON(rate)
	}
}