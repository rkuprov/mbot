package store_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rkuprov/mbot/pkg/store"
)

func TestCreateSubscription(t *testing.T) {
	ctx := context.Background()
	c, cleanUp := NewTestClient(ctx)
	defer cleanUp()
	cid, err := c.CreateCustomer(ctx, store.CustomerCreate{
		Name:    "john doe",
		Email:   "jd@gmail.com",
		Contact: "it's a me! mario",
	})

	id, err := c.CreateSubscription(ctx, cid, "2021-01-01", 30)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)

	sub, err := c.GetSubscription(ctx, id)
	assert.NoError(t, err)
	assert.NotEmpty(t, sub)
	assert.Equal(t, "2021-01-31", sub.GetSubscriptionExpiry().AsTime().Format("2006-01-02"))
	assert.Equal(t, id, sub.GetSubscriptionId())
}

func TestGetSubscriptionAll(t *testing.T) {
	ctx := context.Background()
	c, cleanUp := NewTestClient(ctx)
	defer cleanUp()
	cid, err := c.CreateCustomer(ctx, store.CustomerCreate{
		Name:    "john doe",
		Email:   "jd@gmail.com",
		Contact: "it's a me! mario",
	})

	id, err := c.CreateSubscription(ctx, cid, "2021-01-01", 30)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)

	id2, err := c.CreateSubscription(ctx, cid, "2021-01-01", 120)
	assert.NoError(t, err)
	assert.NotEmpty(t, id2)

	sub, err := c.GetSubscriptionsAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, sub, 2)
	assert.Equal(t, "2021-01-31", sub[0].GetSubscriptionExpiry().AsTime().Format("2006-01-02"))
	assert.Equal(t, id, sub[0].GetSubscriptionId())
	assert.Equal(t, "2021-05-01", sub[1].GetSubscriptionExpiry().AsTime().Format("2006-01-02"))
	assert.Equal(t, id2, sub[1].GetSubscriptionId())
}
