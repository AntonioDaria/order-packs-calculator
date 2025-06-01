package service

import (
	"github.com/AntonioDaria/order-packs-calculator/internal/repository"
	"github.com/rs/zerolog"
)

type PackCalculatorService struct {
	packRepo repository.PackSizeRepository
	Logger   zerolog.Logger
}

func NewPackCalculatorService(repo repository.PackSizeRepository, logger zerolog.Logger) *PackCalculatorService {
	return &PackCalculatorService{
		packRepo: repo,
		Logger:   logger,
	}
}

// Compile-time assertion to ensure PackCalculatorService implements PackCalculator interface
var _ PackCalculator = (*PackCalculatorService)(nil)
