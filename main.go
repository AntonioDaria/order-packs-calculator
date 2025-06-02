package main

import (
	"os"

	pc "github.com/AntonioDaria/order-packs-calculator/internal/handler/pack_calculator"
	cfg "github.com/AntonioDaria/order-packs-calculator/internal/handler/pack_config"

	"github.com/AntonioDaria/order-packs-calculator/internal/repository"
	"github.com/AntonioDaria/order-packs-calculator/internal/router"
	"github.com/AntonioDaria/order-packs-calculator/internal/server"
	pack_calculator_service "github.com/AntonioDaria/order-packs-calculator/internal/service/pack_calculator"
	pack_config_service "github.com/AntonioDaria/order-packs-calculator/internal/service/pack_config"

	"github.com/rs/zerolog"
)

func main() {
	// Set up logger
	logger := zerolog.New(os.Stderr).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	// default initial pack sizes
	initialPackSizes := []int{250, 500, 1000}

	// Initialize repository
	packCalculatorRepo := repository.NewInMemoryPackSizeRepository(initialPackSizes, logger)

	// Initialize services
	packService := pack_calculator_service.NewPackCalculatorService(packCalculatorRepo, logger)
	packConfigService := pack_config_service.NewPackConfigService(packCalculatorRepo, logger)

	// Initialize handlers
	packCalculatorHandler := pc.NewPackCalculatorHandler(packService, logger)
	packConfigHandler := cfg.NewPackConfigHandler(packConfigService, logger)

	// Group handlers
	handlers := &router.Handlers{
		PackCalculatorHandler: packCalculatorHandler,
		PackConfigHandler:     packConfigHandler,
	}
	app := router.New(handlers)

	// Set up static file serving for the frontend
	app.Static("/", "./static") // this serves index.html at http://localhost:3000/

	// Initialize and run server
	httpServer := server.New(logger, app)
	if err := httpServer.Run(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to run server")
	}
}
