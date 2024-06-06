package store

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

var (
	ErrTokenNotFound = errors.New("token not found")
)

func (s *Store) ConfirmToken(ctx context.Context, token string) (bool, string) {
	if token == "" {
		return false, ""
	}

	var validUntil []byte
	err := s.db.View(func(tx *bbolt.Tx) error {
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

func (s *Store) CreateToken(ctx context.Context, validUntil time.Time) string {
	token := uuid.New().String()
	err := s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("tokens-customer"))
		return b.Put([]byte(token), []byte(validUntil.Format("2006-01-02 15:04:05")))
	})
	if err != nil {
		return ""
	}

	return token
}
