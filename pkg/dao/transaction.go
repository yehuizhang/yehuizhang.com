package dao

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
	"yehuizhang.com/go-webapp-gin/pkg/database"
	"yehuizhang.com/go-webapp-gin/src/util/ctxUtil"
)

var transactionSet = wire.NewSet(wire.Struct(new(Transaction), "*"))

type Transaction struct {
	DB database.IPostgres
}

func (a *Transaction) Exec(ctx context.Context, fn func(context.Context) error) error {
	if _, ok := ctxUtil.FromTransaction(ctx); ok {
		return fn(ctx)
	}

	return a.DB.Client().Transaction(func(tx *gorm.DB) error {
		return fn(ctxUtil.NewTransaction(ctx, tx))
	})
}
