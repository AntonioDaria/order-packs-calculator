package pack_config

import (
	packconfig "github.com/AntonioDaria/order-packs-calculator/internal/service/pack_config"
	"github.com/rs/zerolog"
)

type PackConfigHandler struct {
	Service packconfig.PackConfig
	Logger  zerolog.Logger
}

func NewPackConfigHandler(service packconfig.PackConfig, logger zerolog.Logger) *PackConfigHandler {
	return &PackConfigHandler{
		Service: service,
		Logger:  logger,
	}
}
