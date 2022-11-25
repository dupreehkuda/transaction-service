package processors

import (
	"github.com/shopspring/decimal"
	"go.uber.org/zap"

	i "github.com/dupreehkuda/transaction-service/internal"
)

func (a *actions) WriteToQueue(account, operation string, funds decimal.Decimal) error {
	err := a.fileWriter.WriteNewRequest()
	if err != nil {
		a.logger.Error("Error occurred while writing in fileKeeper", zap.Error(err))
		return err
	}

	a.mtx.Lock()
	defer a.mtx.Unlock()

	a.users[account] = append(a.users[account], i.Job{
		Id:        "",
		Account:   account,
		Operation: operation,
		Amount:    funds,
	})

	return nil
}

func (a *actions) GetQueues() {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	for key := range a.users {
		for _, val := range a.users[key] {
			a.collector <- val
		}

		delete(a.users, key)
	}
}
