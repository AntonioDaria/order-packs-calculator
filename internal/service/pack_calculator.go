package service

import (
	"github.com/AntonioDaria/order-packs-calculator/internal/repository"
)

type PackCalculatorService struct {
	packRepo repository.PackSizeRepository
}

func NewPackCalculatorService(repo repository.PackSizeRepository) *PackCalculatorService {
	return &PackCalculatorService{
		packRepo: repo,
	}
}

// Compile-time assertion to ensure PackCalculatorService implements PackCalculator interface
var _ PackCalculator = (*PackCalculatorService)(nil)
