package router

import (
	hc "github.com/AntonioDaria/order-packs-calculator/internal/handler"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	HealthcheckHandler *hc.Handler
}

func New(handlers *Handlers) *fiber.App {
	app := fiber.New()

	app.Get("/api/healthcheck", handlers.HealthcheckHandler.Check)

	return app
}
