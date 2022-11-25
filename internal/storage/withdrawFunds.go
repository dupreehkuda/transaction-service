package storage

import (
	"context"
	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

func (s storage) WithdrawBalance(account string, withdraw decimal.Decimal) error {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()
	pgxdecimal.Register(conn.Conn().TypeMap())

	_, err = conn.Exec(ctx, "update accounts set amount = amount - $1 where account_id = $2;", withdraw, account)
	if err != nil {
		s.logger.Error("Error when executing statement", zap.Error(err))
		return err
	}

	return nil
}
