package store_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rkuprov/mbot/pkg/cfg"
	"github.com/rkuprov/mbot/pkg/store"
)

func TestStore_CreateGetCustomer(t *testing.T) {
	configs, err := cfg.Load()
	require.NoError(t, err)
	client, err := store.New(configs.Postgres)
	assert.NoError(t, err)

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
