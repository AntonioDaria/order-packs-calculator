package packconfig

import (
	"context"
	"errors"

	"github.com/AntonioDaria/order-packs-calculator/internal/repository"
	"github.com/rs/zerolog"
)

var (
	ErrEmptyPackList   = errors.New("pack size list cannot be empty")
	ErrInvalidPackSize = errors.New("pack sizes must be positive integers")
)

//go:generate mockgen -source=pack_config.go -destination=mock/pack_config_mock.go -package=mocks
type PackConfig interface {
	GetPackSizes(ctx context.Context) ([]int, error)
	UpdatePackSizes(ctx context.Context, sizes []int) error
}

type PackConfigService struct {
	Logger zerolog.Logger
	Repo   repository.PackSizeRepository
}

func NewPackConfigService(repo repository.PackSizeRepository, logger zerolog.Logger) *PackConfigService {
	return &PackConfigService{
		Logger: logger,
		Repo:   repo,
	}
}

func (s *PackConfigService) GetPackSizes(ctx context.Context) ([]int, error) {
	sizes, err := s.Repo.GetAll(ctx)
	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to retrieve pack sizes")
		return nil, err
	}

	s.Logger.Info().Ints("pack_sizes", sizes).Msg("Retrieved pack sizes")
	return sizes, nil
}

func (s *PackConfigService) UpdatePackSizes(ctx context.Context, sizes []int) error {
	if len(sizes) == 0 {
		s.Logger.Warn().Msg("Attempted to update pack sizes with empty list")
		return ErrEmptyPackList
	}

	for _, size := range sizes {
		if size <= 0 {
			s.Logger.Warn().Int("invalid_pack_size", size).Msg("Invalid pack size provided")
			return ErrInvalidPackSize
		}
	}

	err := s.Repo.ReplaceAll(ctx, sizes)
	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to update pack sizes")
		return err
	}

	s.Logger.Info().Ints("updated_pack_sizes", sizes).Msg("Pack sizes successfully updated")
	return nil
}
