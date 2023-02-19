package ctxUtil

import (
	"context"
	"fmt"
)

type (
	transactionCtx     struct{}
	transactionLockCtx struct{}
)

func NewTransaction(ctx context.Context, transaction interface{}) context.Context {
	return context.WithValue(ctx, transactionCtx{}, transaction)
}

func FromTransaction(ctx context.Context) (interface{}, bool) {
	v := ctx.Value(transactionCtx{})
	return v, v != nil
}

// NewTransactionLockUpdate uses page or row locks, rows examined by the query are write-locked until the end of the current transaction.
func NewTransactionLockUpdate(ctx context.Context) (context.Context, error) {
	if err := keyMustHasNoValue(ctx, transactionLockCtx{}); err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, transactionLockCtx{}, "UPDATE"), nil
}

// NewTransactionLockShare set shared locks that permit other transactions to read the examined rows but not to update or delete them
func NewTransactionLockShare(ctx context.Context) (context.Context, error) {
	if err := keyMustHasNoValue(ctx, transactionLockCtx{}); err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, transactionLockCtx{}, "SHARE"), nil
}

func FromTransactionLock(ctx context.Context) (string, bool) {
	if v := ctx.Value(transactionLockCtx{}); v != nil {
		return v.(string), true
	}
	return "", false
}

func keyMustHasNoValue(ctx context.Context, key any) error {
	value := ctx.Value(key)
	if value != nil {
		return fmt.Errorf("the contaxt key has already been set a value %s", value)
	}
	return nil
}
