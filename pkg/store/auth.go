package store

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.etcd.io/bbolt"
	"mbot/pkg/datamodel"
	"time"
)

var (
	ErrTokenNotFound = errors.New("token not found")
)

func (c *Client) ConfirmToken(ctx context.Context, token string) (bool, string) {
	if token == "" {
		return false, ""
	}

	var validUntil []byte
	err := c.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("token-client"))
		validUntil = b.Get([]byte(token))
		if len(validUntil) == 0 {
			return ErrTokenNotFound
		}
		return nil
	})
	if errors.Is(err, ErrTokenNotFound) {
		return false, ""
	}
	if err != nil {
		return false, ""
	}

	date := string(validUntil)

	return true, date
}

func (c *Client) CreateToken(ctx context.Context, validUntil time.Time) string {
	token := uuid.New().String()
	err := c.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("tokens-customer"))
		return b.Put([]byte(token), []byte(validUntil.Format("2006-01-02 15:04:05")))
	})
	if err != nil {
		return ""
	}

	return token
}

func (c *Client) CreateCustomer(ctx context.Context, customer datamodel.Customer) string {
	id := uuid.New().String()
	err := c.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("customers"))
		return b.Put([]byte(id), customer.Encode())
	})
	if err != nil {
		return ""
	}

	return id
}
