package worker

import (
	i "github.com/dupreehkuda/transaction-service/internal/interfaces"
	"go.uber.org/zap"
)

type Worker struct {
	fKeeper i.FKeeper
	logic   i.Processors
	storage i.Stored
	logger  *zap.Logger
}

func (w Worker) Run() {

}

// New creates new instance of Worker
func New(fKeeper i.FKeeper, logic i.Processors, storage i.Stored, logger *zap.Logger) *Worker {
	return &Worker{
		fKeeper: fKeeper,
		logic:   logic,
		storage: storage,
		logger:  logger,
	}
}
