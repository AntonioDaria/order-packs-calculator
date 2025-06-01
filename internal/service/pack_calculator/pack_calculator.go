package pack_calculator

import (
	"context"

	"github.com/AntonioDaria/order-packs-calculator/internal/repository"
	"github.com/rs/zerolog"
)

//go:generate mockgen -source=pack_calculator.go -destination=mock/pack_calculator_mock.go -package=mocks

type PackCalculator interface {
	CalculatePacks(ctx context.Context, itemCount int) (map[int]int, int, error)
}
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
