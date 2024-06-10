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
	// id -> CUSTOMER
	customersBucket = []byte("customers")
	// id -> CUSTOMER
	customersDeleteBucket = []byte("customers-delete")

	// CUSTOMER -> "customer", "subscriptions", "stats"
	customerData = []byte("customer")
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

func NewWithClient(_ context.Context, db *bbolt.DB) (*Store, error) {
	c := &Store{db: db}
	if err := initBucket(c); err != nil {
		return nil, err
	}
	return c, nil
}

func initBucket(c *Store) error {
	return c.db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(customersBucket)
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists(customersDeleteBucket)
		if err != nil {
			return err
		}

		return nil
	})
}
