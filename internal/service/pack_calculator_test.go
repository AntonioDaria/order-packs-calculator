package service

import (
	"context"
	"errors"
	"testing"

	"github.com/AntonioDaria/order-packs-calculator/internal/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCalculatePacks_ExactMatchHappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockPackSizeRepository(ctrl)

	// mock the repository call to return specific pack sizes
	mockRepo.EXPECT().
		GetAll(gomock.Any()).
		Return([]int{250, 500, 1000}, nil).
		Times(1)

	service := NewPackCalculatorService(mockRepo)

	result, total, err := service.CalculatePacks(context.Background(), 1750)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 1750, total)
	expected := map[int]int{1000: 1, 500: 1, 250: 1}
	assert.Equal(t, expected, result)
}

func TestCalculatePacks_EdgeCase_WithMockRepo(t *testing.T) {
	// Setup gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock repository
	mockRepo := mock.NewMockPackSizeRepository(ctrl)

	// Mock the GetAll method to return specific pack sizes
	mockRepo.EXPECT().
		GetAll(gomock.Any()).
		Return([]int{23, 31, 53}, nil).
		Times(1)

	// Initialize service with mock repo
	service := NewPackCalculatorService(mockRepo)

	// Call the method with a large number that is an edge case
	result, total, err := service.CalculatePacks(context.Background(), 500000)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 500000, total)
	expected := map[int]int{23: 2, 31: 7, 53: 9429}
	assert.Equal(t, expected, result)
}

func TestCalculatePacks_OverageRequired(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockPackSizeRepository(ctrl)

	// mock the repository call to return specific pack sizes
	mockRepo.EXPECT().
		GetAll(gomock.Any()).
		Return([]int{250, 500, 1000, 2000, 5000}, nil).
		Times(1)

	service := NewPackCalculatorService(mockRepo)

	result, total, err := service.CalculatePacks(context.Background(), 12001)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 12250, total)
	assert.Equal(t, map[int]int{5000: 2, 2000: 1, 250: 1}, result)
}

func TestCalculatePacks_InvalidItemCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockPackSizeRepository(ctrl)
	service := NewPackCalculatorService(mockRepo)

	_, _, err := service.CalculatePacks(context.Background(), 0)

	assert.Error(t, err)
	assert.EqualError(t, err, "itemCount must be greater than 0")
}

func TestCalculatePacks_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockPackSizeRepository(ctrl)
	mockRepo.EXPECT().
		GetAll(gomock.Any()).
		Return(nil, errors.New("db error"))

	service := NewPackCalculatorService(mockRepo)

	_, _, err := service.CalculatePacks(context.Background(), 10)

	assert.Error(t, err)
	assert.EqualError(t, err, "db error")
}
