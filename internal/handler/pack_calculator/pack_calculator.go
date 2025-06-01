package pack_calculator

import (
	pack_calculator_service "github.com/AntonioDaria/order-packs-calculator/internal/service/pack_calculator"
	"github.com/rs/zerolog"
)

type PackCalculatorHandler struct {
	service pack_calculator_service.PackCalculator
	logger  zerolog.Logger
}

func NewPackCalculatorHandler(s pack_calculator_service.PackCalculator, logger zerolog.Logger) *PackCalculatorHandler {
	return &PackCalculatorHandler{
		service: s,
		logger:  logger,
	}
}
