package store

import (
	"context"
	"log"
	"time"

	"go.etcd.io/bbolt"
)

/*
CUSTOMERS (id -> CUSTOMER)
	|_CUSTOMER --
	   |_META
			|_id -> string
			|_name -> string
			|_email -> string
			|_contact -> string
        |_SUBSCRIPTIONS
			|_SUBSCRIPTION
					|_id -> string
                        JSON
						|_start_date -> time
						|_duration -> int
			            |_info -> string
        |_Stats
			|_TBD

DELETED (slug -> CUSTOMER)
*/

var (
	// slug -> customer
	customerBucket                   = []byte("customers")
	customerSlugBucket               = []byte("customer-slug")
	customerDeleteBucket             = []byte("customer-delete")
	customerSlugDeleteBucket         = []byte("customer-slug-delete")
	customerSubscriptionDeleteBucket = []byte("customer-subscription-delete")
)

type Store struct {
	db *bbolt.DB
}

func NewClient(ctx context.Context) (*Store, func() error) {
	c := &Store{}
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

func initBucket(c *Store) error {
	return c.db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(customerBucket)
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists(customerDeleteBucket)
		if err != nil {
			return err
		}

		return nil
	})
}
