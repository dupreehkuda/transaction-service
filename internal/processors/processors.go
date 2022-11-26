package processors

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sync"

	"go.uber.org/zap"

	i "github.com/dupreehkuda/transaction-service/internal"
	intf "github.com/dupreehkuda/transaction-service/internal/interfaces"
)

type processor struct {
	fileKeeper intf.FKeeper
	logger     *zap.Logger
	users      map[string][]i.Job
	collector  chan i.Job
	mtx        sync.Mutex
}

func (p *processor) SyncCollector() chan i.Job {
	return p.collector
}

// New creates new instance of actions
func New(fileWriter intf.FKeeper, logger *zap.Logger) *processor {
	return &processor{
		fileKeeper: fileWriter,
		logger:     logger,
		users:      map[string][]i.Job{},
		collector:  make(chan i.Job, 50),
	}
}

// ReadUnprocessedOnLaunch reads all unprocessed request and writes them to queue after startup
func (p *processor) ReadUnprocessedOnLaunch() {
	jobs := p.fileKeeper.GetUnprocessed()

	for _, val := range jobs {
		err := p.WriteToQueue(val.Id, val.Account, val.Operation, val.Amount)
		if err != nil {
			p.logger.Error("Unprocessed", zap.Any("job", val))
			continue
		}
	}
}

// OperationHash returns short hash string
func OperationHash(account, operation, date string) string {
	coupled := fmt.Sprintf("%s%s%s", account, operation, date)
	hsha := sha1.Sum([]byte(coupled))

	return hex.EncodeToString(hsha[:5])
}
