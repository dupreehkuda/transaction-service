package storage

import (
	"context"

	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

// CheckBalance checks if there is enough funds on account for withdrawal
func (s storage) CheckBalance(account string, want decimal.Decimal) bool {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return false
	}
	defer conn.Release()
	pgxdecimal.Register(conn.Conn().TypeMap())

	var current decimal.Decimal

	err = conn.QueryRow(ctx, "select amount from accounts where account_id = $1;", account).Scan(&current)
	if err != nil {
		s.logger.Error("Error when executing statement", zap.Error(err))
		return false
	}

	if current.LessThan(want) {
		return false
	}

	return true
}
