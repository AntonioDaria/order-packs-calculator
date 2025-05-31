package main

import (
	"os"

	hc "github.com/AntonioDaria/order-packs-calculator/internal/handler"
	"github.com/AntonioDaria/order-packs-calculator/internal/router"
	"github.com/AntonioDaria/order-packs-calculator/internal/server"

	"github.com/rs/zerolog"
)

func main() {
	// Set up logger
	logger := zerolog.New(os.Stderr).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	// Initialize handlers
	healthcheckHandler := hc.NewHandler(logger)

	// Initialize router and setup routes
	handlers := &router.Handlers{
		HealthcheckHandler: healthcheckHandler,
	}
	app := router.New(handlers)

	// Initialize and run server
	httpServer := server.New(logger, app)
	if err := httpServer.Run(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to run server")
	}
}
