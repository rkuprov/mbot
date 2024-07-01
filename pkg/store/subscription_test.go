package store_test

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"

	"github.com/rkuprov/mbot/pkg/store"
)

func TestStore_CreateGetSubscription(t *testing.T) {

	ctx := context.Background()
	client, cleanup, err := store.NewTestStore()
	assert.NoError(t, err)
	defer cleanup()

	cid, err := client.CreateCustomer(ctx, store.CustomerCreate{
		Name:    gofakeit.Name(),
		Email:   gofakeit.Email(),
		Contact: gofakeit.Phone(),
	})

	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	sub1 := store.SubscriptionCreate{
		CustomerID:     cid,
		StartDate:      start,
		ExpirationDate: start.AddDate(1, 0, 0),
	}
	sub2 := store.SubscriptionCreate{
		CustomerID:     cid,
		StartDate:      start,
		ExpirationDate: start.AddDate(1, 0, 0),
	}

	id1, err := client.CreateSubscription(ctx, sub1)
	assert.NoError(t, err)
	assert.NotEmpty(t, id1)
	id2, err := client.CreateSubscription(ctx, sub2)
	assert.NoError(t, err)
	assert.NotEmpty(t, id2)

	customer, err := client.GetCustomer(ctx, cid)
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
