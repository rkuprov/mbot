package store_test

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rkuprov/mbot/pkg/store"
)

func TestStore_CreateGetCustomer(t *testing.T) {
	client, cleanup, err := store.NewTestStore()
	require.NoError(t, err)
	defer cleanup()

	c := store.CustomerCreate{
		Name:    gofakeit.Name(),
		Email:   gofakeit.LastName(),
		Contact: gofakeit.Phone(),
	}

	id, err := client.CreateCustomer(context.Background(), c)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
	out, err := client.GetCustomer(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, c.Name, out.Name)
	assert.Equal(t, c.Email, out.Email)
	assert.Equal(t, c.Contact, out.Contact)
}

func TestStore_CreateCustomerDuplicateEmail(t *testing.T) {
	client, cleanup, err := store.NewTestStore()
	require.NoError(t, err)
	defer cleanup()

	c := store.CustomerCreate{
		Name:    gofakeit.Name(),
		Email:   gofakeit.LastName(),
		Contact: gofakeit.Phone(),
	}

	_, err = client.CreateCustomer(context.Background(), c)
	require.NoError(t, err)

	c2 := store.CustomerCreate{
		Name:    gofakeit.Name(),
		Email:   c.Email,
		Contact: gofakeit.Phone(),
	}
	_, err = client.CreateCustomer(context.Background(), c2)
	assert.Error(t, err)
}

func TestStore_DeleteCustomer(t *testing.T) {
	client, cleanup, err := store.NewTestStore()
	require.NoError(t, err)
	defer cleanup()

	c := store.CustomerCreate{
		Name:    gofakeit.Name(),
		Email:   gofakeit.LastName(),
		Contact: gofakeit.Phone(),
	}

	id, err := client.CreateCustomer(context.Background(), c)
	require.NoError(t, err)
	sid, err := client.CreateSubscription(context.Background(), store.SubscriptionCreate{
		CustomerID:     id,
		StartDate:      time.Now(),
		ExpirationDate: time.Now().AddDate(1, 0, 0),
	})
	require.NoError(t, err)

	err = client.DeleteCustomer(context.Background(), id)
	require.NoError(t, err)

	out, err := client.GetCustomer(context.Background(), id)
	assert.Nil(t, out)
	assert.Nil(t, err)
	_, err = client.GetSubscription(context.Background(), sid)
	assert.Error(t, err)
}
