package repository

import "context"

type PackSizeRepository interface {
	GetAll(ctx context.Context) ([]int, error)
	ReplaceAll(ctx context.Context, sizes []int) error
}
