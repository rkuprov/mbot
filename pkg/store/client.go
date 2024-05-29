package store

import (
	"context"
	"log"
	"time"

	"go.etcd.io/bbolt"
)

type Client struct {
	db *bbolt.DB
}

func NewClient(ctx context.Context) (*Client, func() error) {
	c := &Client{}
	db, err := bbolt.Open("my.db", 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal("failed to open database: ", err)
	}
	c.db = db
	if err = initBucket(c); err != nil {
		log.Fatal("failed to initialize bucket: ", err)
	}
	return c, db.Close
}

func initBucket(c *Client) error {
	return c.db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(customerBucket)
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists(customerSlugBucket)
		if err != nil {
			return err
		}

		return nil
	})
}
