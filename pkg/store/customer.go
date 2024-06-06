package store

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"go.etcd.io/bbolt"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
)

type Customer struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Contact string `json:"contact"`
}
type CustomerCreate struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Contact string `json:"contact"`
}
type CustomerUpdate = CustomerCreate

func (s *Store) CreateCustomer(_ context.Context, customer CustomerCreate) (string, error) {
	tx, err := s.db.Begin(true)
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	customerB := tx.Bucket(customerBucket)
	cB, err := customerB.CreateBucketIfNotExists([]byte("customer"))
	if err != nil {
		return "", err
	}
	id, err := cB.NextSequence()
	if err != nil {
		return "", err
	}
	if err = putIntoBucket(cB, strconv.FormatUint(id, 10), customer); err != nil {
		return "", err
	}
	err = tx.Commit()
	if err != nil {
		return "", err

	}

	return strconv.FormatUint(id, 10), nil
}

func (s *Store) GetCustomer(_ context.Context, id string) (mbotpb.Customer, error) {
	bId := []byte(id)
	var customerDb Customer
	var subIds []string
	err := s.db.View(func(tx *bbolt.Tx) error {
		root := tx.Bucket(customerBucket)
		c := root.Bucket([]byte("customer"))
		val := c.Get(bId)
		if val == nil {
			return fmt.Errorf("customer with ID %s not found", id)
		}

		subs := root.Bucket([]byte("subscriptions"))
		subIds = strings.Split(string(subs.Get(bId)), "#")

		return json.Unmarshal(val, &customerDb)
	})
	if err != nil {
		return mbotpb.Customer{}, err
	}

	return mbotpb.Customer{
		Id:              customerDb.ID,
		Name:            customerDb.Name,
		Email:           customerDb.Email,
		Contact:         customerDb.Contact,
		SubscriptionIds: subIds,
	}, nil
}

func (s *Store) GetCustomersAll(_ context.Context) ([]mbotpb.Customer, error) {
	var customers []mbotpb.Customer

	err := s.db.View(func(tx *bbolt.Tx) error {
		root := tx.Bucket(customerBucket)
		c := root.Bucket([]byte("customer"))
		return c.ForEach(func(k, v []byte) error {
			var customerDb Customer
			if err := json.Unmarshal(v, &customerDb); err != nil {
				return err
			}
			subs := root.Bucket([]byte("subscriptions"))
			subIds := strings.Split(string(subs.Get(k)), "#")
			customers = append(customers, mbotpb.Customer{
				Id:              customerDb.ID,
				Name:            customerDb.Name,
				Email:           customerDb.Email,
				Contact:         customerDb.Contact,
				SubscriptionIds: subIds,
			})
			return nil
		})
	})
	if err != nil {
		return nil, err

	}

	return customers, nil
}

func (s *Store) UpdateCustomer(_ context.Context, id string, customer CustomerUpdate) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		root := tx.Bucket(customerBucket)
		c := root.Bucket([]byte("customer"))
		val := c.Get([]byte(id))
		if len(val) == 0 {
			return fmt.Errorf("customer with ID %s not found", id)
		}

		if err := putIntoBucket(c, id, customer); err != nil {
			return err
		}

		return nil
	})
}

func (s *Store) DeleteCustomer(_ context.Context, id string) error {
	bId := []byte(id)

	tx, err := s.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var customerToDelete []byte
	var subIds []byte

	root := tx.Bucket(customerBucket)
	c := root.Bucket([]byte("customer"))
	customerToDelete = c.Get(bId)
	if customerToDelete == nil {
		return fmt.Errorf("customer with ID %s not found", id)
	}
	err = c.Delete(bId)
	if err != nil {
		return err
	}
	subs := root.Bucket(subscriptionsBucket)
	subIds = subs.Get(bId)
	err = subs.Delete(bId)
	if err != nil {
		return err
	}

	root = tx.Bucket(customerDeleteBucket)
	c, err = root.CreateBucketIfNotExists([]byte("customer"))
	if err != nil {
		return err
	}
	err = c.Put(bId, customerToDelete)
	if err != nil {
		return err
	}
	c, err = root.CreateBucketIfNotExists([]byte("subscriptions"))
	if err != nil {
		return err
	}
	err = c.Put(bId, subIds)
	if err != nil {
		return err
	}

	return nil
}

func putIntoBucket(b *bbolt.Bucket, id string, data any) error {
	if buf, err := json.Marshal(data); err != nil {
		return err
	} else if err := b.Put([]byte(id), buf); err != nil {
		return err
	}

	return nil
}
