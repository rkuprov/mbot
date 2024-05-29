package store

import (
	"context"

	"go.etcd.io/bbolt"
)

type Client struct {
	db bbolt.DB
}

func NewClient(ctx context.Context) *Client {
	return &Client{}
}
