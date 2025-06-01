package main

import (
	"os"

	hc "github.com/AntonioDaria/order-packs-calculator/internal/handler"
	"github.com/AntonioDaria/order-packs-calculator/internal/repository"
	"github.com/AntonioDaria/order-packs-calculator/internal/router"
	"github.com/AntonioDaria/order-packs-calculator/internal/server"
	"github.com/AntonioDaria/order-packs-calculator/internal/service"

	"github.com/rs/zerolog"
)

func main() {
	// Set up logger
	logger := zerolog.New(os.Stderr).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	// Initialize repository
	packCalculatorRepo := repository.NewInMemoryPackSizeRepository(nil, logger)

	// Initialize services
	packService := service.NewPackCalculatorService(packCalculatorRepo, logger)

	// Initialize handlers
	packCalculatorHandler := hc.NewPackCalculatorHandler(packService, logger)

	// Group handlers
	handlers := &router.Handlers{
		PackCalculatorHandler: packCalculatorHandler,
	}
	app := router.New(handlers)

	// Initialize and run server
	httpServer := server.New(logger, app)
	if err := httpServer.Run(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to run server")
	}
}
