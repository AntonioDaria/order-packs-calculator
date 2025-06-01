package packconfig

import (
	"context"
	"errors"
	"testing"

	mocks "github.com/AntonioDaria/order-packs-calculator/internal/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestPackConfigService_GetPackSizes_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPackSizeRepository(ctrl)
	mockRepo.EXPECT().
		GetAll(gomock.Any()).
		Return([]int{250, 500, 1000}, nil)

	service := &PackConfigService{
		Logger: zerolog.New(nil),
		Repo:   mockRepo,
	}

	sizes, err := service.GetPackSizes(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, []int{250, 500, 1000}, sizes)
}

func TestPackConfigService_GetPackSizes_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPackSizeRepository(ctrl)
	mockRepo.EXPECT().
		GetAll(gomock.Any()).
		Return(nil, errors.New("repo failure"))

	service := &PackConfigService{
		Logger: zerolog.New(nil),
		Repo:   mockRepo,
	}

	sizes, err := service.GetPackSizes(context.Background())
	assert.Error(t, err)
	assert.Nil(t, sizes)
}

func TestPackConfigService_UpdatePackSizes_EmptyList(t *testing.T) {
	service := &PackConfigService{
		Logger: zerolog.New(nil),
		Repo:   nil, // no need for mock — will fail before repo is called
	}

	err := service.UpdatePackSizes(context.Background(), []int{})
	assert.Error(t, err)
	assert.EqualError(t, err, "pack size list cannot be empty")
}

func TestPackConfigService_UpdatePackSizes_InvalidValue(t *testing.T) {
	service := &PackConfigService{
		Logger: zerolog.New(nil),
		Repo:   nil, // no need for mock — will fail before repo is called
	}

	err := service.UpdatePackSizes(context.Background(), []int{250, -10})
	assert.Error(t, err)
	assert.EqualError(t, err, "pack sizes must be positive integers")
}

func TestPackConfigService_UpdatePackSizes_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPackSizeRepository(ctrl)
	mockRepo.EXPECT().
		ReplaceAll(gomock.Any(), []int{250, 500}).
		Return(nil)

	service := &PackConfigService{
		Logger: zerolog.New(nil),
		Repo:   mockRepo,
	}

	err := service.UpdatePackSizes(context.Background(), []int{250, 500})
	assert.NoError(t, err)
}

func TestPackConfigService_UpdatePackSizes_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPackSizeRepository(ctrl)
	mockRepo.EXPECT().
		ReplaceAll(gomock.Any(), []int{1000}).
		Return(errors.New("repo failure"))

	service := &PackConfigService{
		Logger: zerolog.New(nil),
		Repo:   mockRepo,
	}

	err := service.UpdatePackSizes(context.Background(), []int{1000})
	assert.Error(t, err)
	assert.EqualError(t, err, "repo failure")
}
