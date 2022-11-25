package internal

import "github.com/shopspring/decimal"

type Job struct {
	Id        string
	Account   string
	Operation string
	Amount    decimal.Decimal
}
