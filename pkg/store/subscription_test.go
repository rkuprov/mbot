package store_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rkuprov/mbot/pkg/cfg"
	"github.com/rkuprov/mbot/pkg/store"
)

func TestStore_CreateSubscription(t *testing.T) {
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
	fmt.Println(customer.String())
}

func TestStore_GetSubscription(t *testing.T) {
	ctx := context.Background()
	configs, err := cfg.Load()
	require.NoError(t, err)
	client, err := store.New(configs.Postgres)
	assert.NoError(t, err)

	out, err := client.GetSubscription(ctx, "f57bfb0c-3c41-4833-b468-012f72eed020")
	assert.NoError(t, err)
	fmt.Println(out.String())

	outAll, err := client.GetSubscriptionsAll(ctx)
	assert.NoError(t, err)
	for _, sub := range outAll {
		fmt.Println(sub.String())
	}

	fmt.Println("By customer:")
	outC, err := client.GetSubscriptionByCustomer(ctx, "1")
	assert.NoError(t, err)
	for _, sub := range outC {
		fmt.Println(sub.String())
	}
}
