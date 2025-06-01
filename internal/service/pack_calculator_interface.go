package service

import (
	"context"
)

type PackCalculator interface {
	CalculatePacks(ctx context.Context, itemCount int) (map[int]int, int, error)
}
