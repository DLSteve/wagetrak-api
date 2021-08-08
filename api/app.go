package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
	routesV1 "wagetrak-api/api/routes/v1"
	_ "wagetrak-api/pkg/config"
	"wagetrak-api/pkg/currency"
	"wagetrak-api/pkg/utils"
)

func main() {
	// setup fiber
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("WageTrak API"))
	})

	// setup logger
	var logger *zap.Logger
	isDebug, _ := utils.GetEnvBool("DEBUG")
	if isDebug {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	// setup currency client
	currencyClient := currency.Client{
		Logger: sugar,
	}
	currencySvc := currency.NewCurrencyService(&currencyClient)

	// setup API routes
	currencyV1 := app.Group("/currency/v1")
	routesV1.CurrencyRouterV1(currencyV1, currencySvc)

	// setup HTTP server
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
