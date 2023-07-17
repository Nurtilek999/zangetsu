package handler

import (
	"zangetsu/internal/domain/service"
	"zangetsu/pkg/logging"
)

type Handler struct {
	services *service.Service
	logger   logging.Logger
}

func NewHandler(s *service.Service, logger logging.Logger) *Handler {
	return &Handler{
		services: s,
		logger:   logger,
	}
}
