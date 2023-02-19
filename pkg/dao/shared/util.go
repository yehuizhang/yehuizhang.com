package shared

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"yehuizhang.com/go-webapp-gin/src/util/ctxUtil"
)

func getDB(ctx context.Context, db *gorm.DB) *gorm.DB {
	if tx, ok := ctxUtil.FromTransaction(ctx); ok {
		if tx, ok := tx.(*gorm.DB); ok {
			if lock, ok := ctxUtil.FromTransactionLock(ctx); ok {
				return tx.Clauses(clause.Locking{Strength: lock})
			}
			return tx
		}
	}

	return db
}

func GetDB(ctx context.Context, db *gorm.DB, model interface{}) *gorm.DB {
	return getDB(ctx, db).Model(model)
}
