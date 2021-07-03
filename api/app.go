package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
	"os/signal"
	"syscall"
	routesV1 "wagetrak-api/api/routes/v1"
	"wagetrak-api/pkg/currency"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("WageTrak API"))
	})

	currencySvc := currency.NewCurrencyService()

	currencyV1 := app.Group("/currency/v1")
	routesV1.CurrencyRouterV1(currencyV1, currencySvc)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errs <- app.Listen(fmt.Sprintf(":%s", port))
	}()

	log.Println("exit", <- errs)
}
