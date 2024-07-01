package store_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rkuprov/mbot/pkg/cfg"
	"github.com/rkuprov/mbot/pkg/store"
)

func TestStore_CreateGetSubscription(t *testing.T) {
	ctx := context.Background()
	configs, err := cfg.Load()
	require.NoError(t, err)
	client, err := store.New(configs.Postgres)
	assert.NoError(t, err)

	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	sub1 := store.SubscriptionCreate{
		CustomerID:     "1",
		StartDate:      start,
		ExpirationDate: start.AddDate(1, 0, 0),
	}
	sub2 := store.SubscriptionCreate{
		CustomerID:     "1",
		StartDate:      start,
		ExpirationDate: start.AddDate(1, 0, 0),
	}

	id1, err := client.CreateSubscription(ctx, sub1)
	assert.NoError(t, err)
	assert.NotEmpty(t, id1)
	id2, err := client.CreateSubscription(ctx, sub2)
	assert.NoError(t, err)
	assert.NotEmpty(t, id2)

	customer, err := client.GetCustomer(ctx, "1")
	assert.NoError(t, err)
	assert.Len(t, customer.SubscriptionIds, 2)

	out1, err := client.GetSubscription(ctx, id1)
	assert.NoError(t, err)
	assert.Equal(t, sub1.CustomerID, out1.CustomerId)
	assert.Equal(t, sub1.StartDate, out1.StartDate.AsTime())
	assert.Equal(t, sub1.ExpirationDate, out1.ExpirationDate.AsTime())

	out2, err := client.GetSubscription(ctx, id2)
	assert.NoError(t, err)
	assert.Equal(t, sub2.CustomerID, out2.CustomerId)
	assert.Equal(t, sub2.StartDate, out2.StartDate.AsTime())
	assert.Equal(t, sub2.ExpirationDate, out2.ExpirationDate.AsTime())
}
