package pack_calculator

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	pack_calculator_service "github.com/AntonioDaria/order-packs-calculator/internal/service/pack_calculator"
	mocks "github.com/AntonioDaria/order-packs-calculator/internal/service/pack_calculator/mock"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func setupTestApp(_ *testing.T, service *mocks.MockPackCalculator) *fiber.App {
	logger := zerolog.New(nil)
	handler := NewPackCalculatorHandler(service, logger)

	app := fiber.New()
	app.Post("/api/calculate", handler.Calculate)

	return app
}

func TestCalculateHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockPackCalculator(ctrl)

	mockService.EXPECT().
		CalculatePacks(gomock.Any(), 12001).
		Return(map[int]int{5000: 2, 2000: 1, 250: 1}, 12250, nil)

	app := setupTestApp(t, mockService)

	body := CalculateRequest{Items: 12001}
	payload, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/calculate", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCalculateHandler_InvalidRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockPackCalculator(ctrl)
	app := setupTestApp(t, mockService)

	req := httptest.NewRequest(http.MethodPost, "/api/calculate", bytes.NewReader([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCalculateHandler_NoCombinationFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockPackCalculator(ctrl)

	mockService.EXPECT().
		CalculatePacks(gomock.Any(), 1).
		Return(nil, 0, pack_calculator_service.ErrNoValidPackCombination)

	app := setupTestApp(t, mockService)

	body := CalculateRequest{Items: 1}
	payload, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/calculate", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
}

func TestCalculateHandler_InvalidItemCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockPackCalculator(ctrl)
	app := setupTestApp(t, mockService)

	body := CalculateRequest{Items: 0} // invalid input
	payload, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/calculate", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
