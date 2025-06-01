package repository

import (
	"context"
	"errors"
	"sort"
	"sync"

	"github.com/rs/zerolog"
)

type InMemoryPackSizeRepository struct {
	mu     sync.RWMutex
	sizes  []int
	Logger zerolog.Logger
}

func NewInMemoryPackSizeRepository(initial []int, logger zerolog.Logger) *InMemoryPackSizeRepository {
	// Make a copy and sort for consistent behavior
	sorted := append([]int{}, initial...)
	//algorithm will perform faster if the pack sizes are sorted in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))

	return &InMemoryPackSizeRepository{
		sizes:  sorted,
		Logger: logger,
	}
}

func (r *InMemoryPackSizeRepository) GetAll(ctx context.Context) ([]int, error) {
	// use a lock to ensure thread safety
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Return a copy to avoid external mutation
	copy := append([]int{}, r.sizes...)

	r.Logger.Debug().Msgf("Retrieved pack sizes: %v", copy)
	return copy, nil
}

func (r *InMemoryPackSizeRepository) ReplaceAll(ctx context.Context, sizes []int) error {
	if len(sizes) == 0 {
		r.Logger.Warn().Msg("Attempted to replace pack sizes with empty list")
		return errors.New("at least one pack size is required")
	}

	sorted := append([]int{}, sizes...)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))

	r.mu.Lock()
	defer r.mu.Unlock()
	r.sizes = sorted

	r.Logger.Info().Ints("new_pack_sizes", sorted).Msg("Pack sizes replaced")
	return nil
}
