package router

import (
	pc "github.com/AntonioDaria/order-packs-calculator/internal/handler/pack_calculator"
	cfg "github.com/AntonioDaria/order-packs-calculator/internal/handler/pack_config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Handlers struct {
	PackCalculatorHandler *pc.PackCalculatorHandler
	PackConfigHandler     *cfg.PackConfigHandler
}

func New(handlers *Handlers) *fiber.App {
	router := fiber.New()

	// Add Recover middleware to handle panics
	router.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	// Pack calculator route
	router.Post("/api/calculate", handlers.PackCalculatorHandler.Calculate)

	// Pack config routes
	router.Get("/api/packs", handlers.PackConfigHandler.GetPacks)
	router.Post("/api/packs", handlers.PackConfigHandler.UpdatePacks)

	return router
}
