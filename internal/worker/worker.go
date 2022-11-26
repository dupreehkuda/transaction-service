package worker

import (
	"time"

	"go.uber.org/zap"

	i "github.com/dupreehkuda/transaction-service/internal"
	intf "github.com/dupreehkuda/transaction-service/internal/interfaces"
)

type Worker struct {
	fKeeper    intf.FKeeper
	processor  intf.Processors
	storage    intf.Stored
	logger     *zap.Logger
	aggregator chan i.Job
}

// Run runs the request worker
func (w Worker) Run() {
	for {
		select {
		case msg := <-w.aggregator:
			w.logger.Debug("Reading agg: ", zap.Any("msg", msg))
			w.ProcessRequest(msg)
		default:
			time.Sleep(500 * time.Millisecond)
			go w.processor.GetQueues()
			//w.logger.Debug("waiting...")
		}
	}
}

// New creates new instance of Worker
func New(fKeeper intf.FKeeper, processor intf.Processors, storage intf.Stored, logger *zap.Logger) *Worker {
	return &Worker{
		fKeeper:    fKeeper,
		processor:  processor,
		storage:    storage,
		logger:     logger,
		aggregator: processor.SyncCollector(),
	}
}
