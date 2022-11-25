package handlers

import (
	"github.com/shopspring/decimal"
	"go.uber.org/zap"

	i "github.com/dupreehkuda/transaction-service/internal/interfaces"
)

type handlers struct {
	processor i.Processors
	logger    *zap.Logger
}

// New creates new instance of handlers
func New(processor i.Processors, logger *zap.Logger) *handlers {
	return &handlers{
		processor: processor,
		logger:    logger,
	}
}

type request struct {
	Account   string          `json:"account"`
	Operation string          `json:"operation"`
	Amount    decimal.Decimal `json:"amount"`
}
