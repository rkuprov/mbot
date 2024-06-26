package store

import (
	"context"
	"errors"
	"time"
)

var (
	ErrTokenNotFound = errors.New("token not found")
)

func (s *Store) ConfirmToken(ctx context.Context, token string) (bool, string) {
	return false, ""
}

func (s *Store) CreateToken(ctx context.Context, validUntil time.Time) string {
	return ""
}
