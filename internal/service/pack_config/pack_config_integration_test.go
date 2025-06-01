package packconfig

import (
	"context"
	"testing"

	"github.com/AntonioDaria/order-packs-calculator/internal/repository"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestPackConfigService_UpdatePackSizes_Integration(t *testing.T) {
	logger := zerolog.New(nil)
	repo := repository.NewInMemoryPackSizeRepository(nil, logger)

	service := &PackConfigService{
		Logger: logger,
		Repo:   repo,
	}

	err := service.UpdatePackSizes(context.Background(), []int{123, 456})
	assert.NoError(t, err)

	sizes, _ := repo.GetAll(context.Background())
	assert.Equal(t, []int{456, 123}, sizes) // because it's sorted descending
}
