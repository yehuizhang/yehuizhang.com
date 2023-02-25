package dao

import (
	"context"
	"gorm.io/gorm"
	"yehuizhang.com/go-webapp-gin/pkg/database"
	"yehuizhang.com/go-webapp-gin/src/util/ctxUtil"
)

type ITransaction interface {
	Exec(ctx context.Context, fn func(context.Context) error) error
}

type Transaction struct {
	DB database.IPostgres
}

func InitTransaction(db database.IPostgres) ITransaction {
	return &Transaction{DB: db}
}

func (a *Transaction) Exec(ctx context.Context, fn func(context.Context) error) error {
	if _, ok := ctxUtil.FromTransaction(ctx); ok {
		return fn(ctx)
	}

	return a.DB.Client().Transaction(func(tx *gorm.DB) error {
		return fn(ctxUtil.NewTransaction(ctx, tx))
	})
}
