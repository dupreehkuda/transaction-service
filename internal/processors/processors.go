package processors

import (
	"sync"

	"go.uber.org/zap"

	i "github.com/dupreehkuda/transaction-service/internal"
	intf "github.com/dupreehkuda/transaction-service/internal/interfaces"
)

type actions struct {
	fileWriter intf.FKeeper
	logger     *zap.Logger
	users      map[string][]i.Job
	collector  chan i.Job
	mtx        sync.Mutex
}

func (a *actions) SyncCollector() chan i.Job {
	return a.collector
}

// New creates new instance of actions
func New(fileWriter intf.FKeeper, logger *zap.Logger) *actions {
	return &actions{
		fileWriter: fileWriter,
		logger:     logger,
		users:      map[string][]i.Job{},
		collector:  make(chan i.Job, 50),
	}
}
