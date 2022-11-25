package processors

import (
	i "github.com/dupreehkuda/transaction-service/internal/interfaces"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type actions struct {
	fileWriter i.FKeeper
	logger     *zap.Logger
}

func (a actions) WriteToQueue(account, operation string, funds decimal.Decimal) error {
	//TODO implement me
	panic("implement me")
}

// New creates new instance of actions
func New(fileWriter i.FKeeper, logger *zap.Logger) *actions {
	return &actions{
		fileWriter: fileWriter,
		logger:     logger,
	}
}
