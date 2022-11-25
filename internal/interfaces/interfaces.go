package interfaces

import (
	"github.com/shopspring/decimal"
	"net/http"
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
}

type FKeeper interface {
	WriteNewRequest()
	UpdateRequest()
}
