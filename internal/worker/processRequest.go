package worker

import (
	"go.uber.org/zap"

	i "github.com/dupreehkuda/transaction-service/internal"
)

func (w Worker) ProcessRequest(job i.Job) {
	if job.Operation == "add" {
		err := w.storage.AddFunds(job.Account, job.Amount)
		if err != nil {
			w.logger.Error("Unprocessed job", zap.Error(err), zap.Any("job", job))
			return
		}
	}

	enough := w.storage.CheckBalance(job.Account, job.Amount)
	if !enough {
		w.logger.Info("Not enough funds", zap.Any("job", job))
		return
	}

	err := w.storage.WithdrawBalance(job.Account, job.Amount)
	if err != nil {
		w.logger.Error("Unprocessed job", zap.Error(err), zap.Any("job", job))
		return
	}

	err = w.fKeeper.UpdateRequest()
	if err != nil {
		w.logger.Error("Failed to update fKeeper", zap.Error(err), zap.Any("job", job))
		return
	}

	return
}
