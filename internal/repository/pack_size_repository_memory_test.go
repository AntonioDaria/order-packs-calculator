package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryPackSizeRepository_GetAll(t *testing.T) {
	initial := []int{100, 50, 200}
	repo := NewInMemoryPackSizeRepository(initial)

	packs, err := repo.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, []int{200, 100, 50}, packs) // sorted descending

	// ensure it's a copy, not a reference
	packs[0] = 999
	newPacks, _ := repo.GetAll(context.Background())
	assert.Equal(t, 200, newPacks[0])
}

func TestInMemoryPackSizeRepository_ReplaceAll(t *testing.T) {
	repo := NewInMemoryPackSizeRepository([]int{100})

	err := repo.ReplaceAll(context.Background(), []int{30, 70, 10})
	assert.NoError(t, err)

	packs, _ := repo.GetAll(context.Background())
	assert.Equal(t, []int{70, 30, 10}, packs)
}

func TestInMemoryPackSizeRepository_ReplaceAll_EmptyInput(t *testing.T) {
	repo := NewInMemoryPackSizeRepository([]int{100})

	err := repo.ReplaceAll(context.Background(), []int{})
	assert.Error(t, err)
	assert.EqualError(t, err, "at least one pack size is required")
}
