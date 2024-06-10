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

func (s *Store) CreateCustomer(_ context.Context, in CustomerCreate) (string, error) {
	tx, err := s.db.Begin(true)
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	customerB := tx.Bucket(customersBucket)
	id, err := customerB.NextSequence()
	if err != nil {
		return "", err
	}

	cB, err := customerB.CreateBucketIfNotExists([]byte(strconv.FormatUint(id, 10)))
	if err != nil {
		return "", err
	}
	_, err = cB.CreateBucketIfNotExists(subscriptions)
	if err != nil {
		return "", err
	}
	cd, err := cB.CreateBucketIfNotExists(customerData)
	if err != nil {
		return "", err
	}

	if err = putIntoBucket(cd, strconv.FormatUint(id, 10), in); err != nil {
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
		root := tx.Bucket(customersBucket)
		cB := root.Bucket(bId)
		val := cB.Bucket(customerData).Get(bId)
		if val == nil {
			return fmt.Errorf("customer with ID %s not found", id)
		}

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
		root := tx.Bucket(customersBucket)
		c := root.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {

			var customerDb Customer
			cd := root.Bucket(k).Bucket(customerData).Get(k)
			if cd == nil {
				return fmt.Errorf("customer with ID %s not found", k)
			}
			if err := json.Unmarshal(cd, &customerDb); err != nil {
				return err
			}

			var subIds []string
			subs := root.Bucket(k).Bucket(subscriptions).Get(k)
			if subs != nil {
				subIds = strings.Split(string(subs), "#")
			}

			customers = append(customers, mbotpb.Customer{
				Id:              customerDb.ID,
				Name:            customerDb.Name,
				Email:           customerDb.Email,
				Contact:         customerDb.Contact,
				SubscriptionIds: subIds,
			})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (s *Store) UpdateCustomer(_ context.Context, id string, in CustomerUpdate) error {
	bId := []byte(id)
	return s.db.Update(func(tx *bbolt.Tx) error {
		root := tx.Bucket(customersBucket)
		cB := root.Bucket(bId)
		val := cB.Bucket(customerData).Get(bId)
		if val == nil {
			return fmt.Errorf("customer with ID %s not found", id)
		}

		return putIntoBucket(cB.Bucket(customerData), id, in)
	})
}

func putIntoBucket(b *bbolt.Bucket, id string, data any) error {
	if buf, err := json.Marshal(data); err != nil {
		return err
	} else if err := b.Put([]byte(id), buf); err != nil {
		return err
	}

	return nil
}
