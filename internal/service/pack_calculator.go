package service

import (
	"order-packs-calculator/internal/repository"
)

type PackCalculatorService struct {
	packRepo repository.PackSizeRepository
}

func NewPackCalculatorService(repo repository.PackSizeRepository) *PackCalculatorService {
	return &PackCalculatorService{
		packRepo: repo,
	}
}
