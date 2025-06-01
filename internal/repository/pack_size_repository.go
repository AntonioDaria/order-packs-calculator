package repository

import "context"

//go:generate mockgen -source=pack_size_repository.go -destination=mock/pack_size_repository_mock.go -package=mock
type PackSizeRepository interface {
	GetAll(ctx context.Context) ([]int, error)
	ReplaceAll(ctx context.Context, sizes []int) error
}
