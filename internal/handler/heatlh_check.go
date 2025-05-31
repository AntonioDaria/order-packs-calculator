package healthcheck

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type Handler struct {
	logger zerolog.Logger
}

func NewHandler(logger zerolog.Logger) *Handler {
	return &Handler{logger: logger}
}

func (h *Handler) Check(c *fiber.Ctx) error {
	h.logger.Info().Msg("Health check requested")
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
