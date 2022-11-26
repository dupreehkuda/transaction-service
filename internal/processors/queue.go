package processors

import (
	"time"

	"github.com/shopspring/decimal"
	"go.uber.org/zap"

	i "github.com/dupreehkuda/transaction-service/internal"
)

// WriteToQueue writes request to user personal queue with existing id or new id
func (p *processor) WriteToQueue(id, account, operation string, funds decimal.Decimal) error {
	if id == "" {
		id = OperationHash(account, operation, time.Now().Format(time.RFC3339))

		err := p.fileKeeper.WriteNewRequest(id, account, operation, funds)
		if err != nil {
			p.logger.Error("Error occurred while writing in fileKeeper", zap.Error(err))
			return err
		}
	}

	p.mtx.Lock()
	defer p.mtx.Unlock()

	p.users[account] = append(p.users[account], i.Job{
		Id:        id,
		Account:   account,
		Operation: operation,
		Amount:    funds,
	})

	return nil
}

// GetQueues collects requests from users queues and sends them to collector channel
func (p *processor) GetQueues() {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	for key := range p.users {
		for _, val := range p.users[key] {
			p.collector <- val
		}

		delete(p.users, key)
	}
}
