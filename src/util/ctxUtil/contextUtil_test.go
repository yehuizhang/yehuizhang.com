package ctxUtil

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromTransaction(t *testing.T) {
	ctx := context.Background()
	result, ok := FromTransaction(ctx)

	assert.False(t, ok)
	assert.Nil(t, result)

	ctx = NewTransaction(ctx, "value")
	result, ok = FromTransaction(ctx)

	assert.True(t, ok)
	assert.Equal(t, "value", result)
}

func TestFromTransactionLockUpdate(t *testing.T) {
	ctx := context.Background()
	result, ok := FromTransactionLock(ctx)

	assert.False(t, ok)

	ctx, _ = NewTransactionLockUpdate(ctx)
	result, ok = FromTransactionLock(ctx)

	assert.True(t, ok)
	assert.Equal(t, "UPDATE", result)
}

func TestFromTransactionLockShareWithValue(t *testing.T) {
	ctx := context.Background()

	ctx, _ = NewTransactionLockShare(ctx)
	result, ok := FromTransactionLock(ctx)

	assert.True(t, ok)
	assert.Equal(t, "SHARE", result)
}

func TestSetMultipleTransaction(t *testing.T) {
	ctx := NewTransaction(context.Background(), "value")
	{
		result, ok := FromTransaction(ctx)
		assert.True(t, ok)
		assert.Equal(t, "value", result)
	}

	ctx, _ = NewTransactionLockUpdate(ctx)
	{
		result, ok := FromTransactionLock(ctx)
		assert.True(t, ok)
		assert.Equal(t, "UPDATE", result)
	}
	ctx, err := NewTransactionLockShare(ctx)
	{
		assert.NotNil(t, err)
		result, ok := FromTransactionLock(ctx)
		assert.True(t, ok)
		assert.Equal(t, "UPDATE", result)
	}

	ctx, err = NewTransactionLockUpdate(ctx)
	{
		assert.NotNil(t, err)
		result, ok := FromTransactionLock(ctx)
		assert.True(t, ok)
		assert.Equal(t, "UPDATE", result)
	}
}
