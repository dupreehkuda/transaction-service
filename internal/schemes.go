package internal

import "github.com/shopspring/decimal"

// Job is a request model
type Job struct {
	Id        string
	Account   string
	Operation string
	Amount    decimal.Decimal
}
