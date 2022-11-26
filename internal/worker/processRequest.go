package worker

import (
	"go.uber.org/zap"

	i "github.com/dupreehkuda/transaction-service/internal"
)

// ProcessRequest is a request processor that handles all requests
func (w Worker) ProcessRequest(job i.Job) {
	if job.Operation == "add" {
		err := w.storage.AddFunds(job.Account, job.Amount)
		if err != nil {
			w.logger.Error("Unprocessed job", zap.Error(err), zap.Any("job", job))
			return
		}

		go w.fKeeper.UpdateRequest(job.Id)
		return
	}

	enough := w.storage.CheckBalance(job.Account, job.Amount)
	if !enough {
		w.logger.Info("Not enough funds", zap.Any("job", job))
		go w.fKeeper.UpdateRequest(job.Id)
		return
	}

	err := w.storage.WithdrawBalance(job.Account, job.Amount)
	if err != nil {
		w.logger.Error("Unprocessed job", zap.Error(err), zap.Any("job", job))
		return
	}

	go w.fKeeper.UpdateRequest(job.Id)

	return
}
