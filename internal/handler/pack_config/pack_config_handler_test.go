package pack_config

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	service "github.com/AntonioDaria/order-packs-calculator/internal/service/pack_config"
	packconfig_service_mock "github.com/AntonioDaria/order-packs-calculator/internal/service/pack_config/mock"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func setupPackConfigTestApp(_ *testing.T, service *packconfig_service_mock.MockPackConfig) *fiber.App {
	logger := zerolog.New(nil)
	handler := NewPackConfigHandler(service, logger)

	app := fiber.New()
	app.Get("/api/packs", handler.GetPacks)
	app.Post("/api/packs", handler.UpdatePacks)
	return app
}

func TestGetPacks_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := packconfig_service_mock.NewMockPackConfig(ctrl)
	mockService.EXPECT().
		GetPackSizes(gomock.Any()).
		Return([]int{250, 500, 1000}, nil)

	app := setupPackConfigTestApp(t, mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/packs", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUpdatePacks_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := packconfig_service_mock.NewMockPackConfig(ctrl)
	mockService.EXPECT().
		UpdatePackSizes(gomock.Any(), []int{250, 500}).
		Return(nil)

	app := setupPackConfigTestApp(t, mockService)

	payload, _ := json.Marshal([]int{250, 500})
	req := httptest.NewRequest(http.MethodPost, "/api/packs", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUpdatePacks_EmptyList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := packconfig_service_mock.NewMockPackConfig(ctrl)
	mockService.EXPECT().
		UpdatePackSizes(gomock.Any(), []int{}).
		Return(service.ErrEmptyPackList)

	app := setupPackConfigTestApp(t, mockService)

	payload, _ := json.Marshal([]int{})
	req := httptest.NewRequest(http.MethodPost, "/api/packs", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
}

func TestUpdatePacks_InvalidValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := packconfig_service_mock.NewMockPackConfig(ctrl)
	mockService.EXPECT().
		UpdatePackSizes(gomock.Any(), []int{1000, -5}).
		Return(service.ErrInvalidPackSize)

	app := setupPackConfigTestApp(t, mockService)

	payload, _ := json.Marshal([]int{1000, -5})
	req := httptest.NewRequest(http.MethodPost, "/api/packs", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
}

func TestUpdatePacks_BadJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := packconfig_service_mock.NewMockPackConfig(ctrl)
	app := setupPackConfigTestApp(t, mockService)

	req := httptest.NewRequest(http.MethodPost, "/api/packs", bytes.NewReader([]byte(`{invalid`)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
