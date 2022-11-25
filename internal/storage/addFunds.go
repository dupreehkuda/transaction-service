package storage

import (
	"context"

	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

func (s storage) AddFunds(accountID string, funds decimal.Decimal) error {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()
	pgxdecimal.Register(conn.Conn().TypeMap())

	_, err = conn.Exec(ctx, `insert into accounts (account_id, amount) values ($1, $2) on conflict (account_id)
	do update set amount = accounts.amount + $2 where accounts.account_id = $1;`, accountID, funds)
	if err != nil {
		s.logger.Error("Error when executing addition statement", zap.Error(err))
		return err
	}

	return nil
}
