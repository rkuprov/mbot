package store_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rkuprov/mbot/pkg/store"
)

func TestStore_RecordUsage(t *testing.T) {
	ctx := context.Background()
	c, cleanup, err := store.NewTestStore()
	require.NoError(t, err)
	defer cleanup()

	cid, err := c.CreateCustomer(ctx, store.CustomerCreate{
		Name:    "John Doe",
		Email:   "doe@gmail.com",
		Contact: "call John's wife",
	})
	assert.NoError(t, err)

	sid, err := c.CreateSubscription(ctx, store.SubscriptionCreate{
		CustomerID:     cid,
		StartDate:      time.Now(),
		ExpirationDate: time.Now().AddDate(0, 1, 0),
	})
	assert.NoError(t, err)

	err = c.RecordUsage(ctx, sid)
	assert.NoError(t, err)
	err = c.RecordUsage(ctx, sid)
	assert.NoError(t, err)

	rows, err := c.Pool.Query(ctx, "SELECT count(*) FROM stats WHERE subscription_id = $1", sid)
	assert.NoError(t, err)
	defer rows.Close()
	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		assert.NoError(t, err)
	}
	assert.Equal(t, 2, count)
}
