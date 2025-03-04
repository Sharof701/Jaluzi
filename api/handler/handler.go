package handler

import (
	"jaluzi/config"
	"jaluzi/pkg/logger"
	"jaluzi/storage"
	"strconv"
)

type handler struct {
	cfg     *config.Config
	logger  logger.LoggerI
	storage storage.StorageI
}

type Response struct {
	Status      int         `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
	Error       interface{} `json:"error"`
}

type ErrorResponse struct {
	Error interface{} `json:"error"`
}

func NewHandler(cfg *config.Config, storage storage.StorageI, logger logger.LoggerI) *handler {
	return &handler{
		cfg:     cfg,
		logger:  logger,
		storage: storage,
	}
}

func (h *handler) getOffsetQuery(offset string) (int, error) {

	if len(offset) <= 0 {
		return 0, nil
	}
	return strconv.Atoi(offset)
}

func (h *handler) getLimitQuery(limit string) (int, error) {

	if len(limit) <= 0 {
		return 10, nil
	}
	return strconv.Atoi(limit)
}
