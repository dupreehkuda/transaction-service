package interfaces

import (
	"net/http"

	"github.com/shopspring/decimal"

	i "github.com/dupreehkuda/transaction-service/internal"
)

type Handlers interface {
	FundsHandler(w http.ResponseWriter, r *http.Request)
}

type Stored interface {
	AddFunds(accountID string, funds decimal.Decimal) error
	WithdrawBalance(account string, withdraw decimal.Decimal) error
	CheckBalance(account string, want decimal.Decimal) bool
}

type Processors interface {
	WriteToQueue(account, operation string, funds decimal.Decimal) error
	GetQueues()
	SyncCollector() chan i.Job
}

type FKeeper interface {
	WriteNewRequest() error
	UpdateRequest() error
}
