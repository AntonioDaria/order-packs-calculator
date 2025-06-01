package router

import (
	h "github.com/AntonioDaria/order-packs-calculator/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Handlers struct {
	PackCalculatorHandler *h.PackCalculatorHandler
}

func New(handlers *Handlers) *fiber.App {
	router := fiber.New()

	// Add Recover middleware to handle panics
	router.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	// pack calculator route
	router.Post("/api/calculate", handlers.PackCalculatorHandler.Calculate)

	return router
}
