package store

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"

	"github.com/rkuprov/mbot/pkg/datamodel"
)

var (
	customerBucket     = []byte("customers")
	customerSlugBucket = []byte("customer-slug")
)

func (c *Client) CreateCustomer(_ context.Context, customer datamodel.Customer) (string, error) {
	id := uuid.New().String()
	customer.Slug = generateSlug(customer.Name, id)
	err := c.db.Update(func(tx *bbolt.Tx) error {
		err := tx.Bucket(customerBucket).Put([]byte(id), encodeCustomer(customer))
		if err != nil {
			return err
		}
		err = tx.Bucket(customerSlugBucket).Put([]byte(customer.Slug), []byte(id))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return customer.Slug, nil
}

func encodeCustomer(c datamodel.Customer) []byte {
	in := fmt.Sprintf("%s#%s#%s#%s#%s", c.Slug, c.Name, c.Email, c.Contact, c.SubscriptionID)

	return []byte(in)
}

func decodeCustomer(in string) (datamodel.Customer, error) {
	parts := strings.Split(in, "#")
	if len(parts) < 5 {
		return datamodel.Customer{}, fmt.Errorf("invalid customer data: %s", in)
	}

	return datamodel.Customer{
		Slug:           parts[0],
		Name:           parts[1],
		Email:          parts[2],
		Contact:        parts[3],
		SubscriptionID: parts[4],
	}, nil
}

func generateSlug(name, id string) string {
	idPiece := id[:4]
	namePiece := name

	if len(name) > 3 {
		namePiece = name[:3]
	}

	return fmt.Sprintf("%s-%s", namePiece, idPiece)
}

func (c *Client) GetCustomer(_ context.Context, slug string) (datamodel.Customer, error) {
	var id []byte
	err := c.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(customerSlugBucket)
		v := b.Get([]byte(slug))
		if v == nil {
			return fmt.Errorf("customer with Slug %s not found", slug)
		}
		id = v
		return nil
	})

	var encoded string
	err = c.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(customerBucket)
		v := b.Get(id)
		if v == nil {
			return fmt.Errorf("customer with ID %s not found", string(id))
		}

		encoded = string(v)
		return nil
	})

	if err != nil {
		return datamodel.Customer{}, err
	}

	return decodeCustomer(encoded)
}

func (c *Client) GetCustomersAll(_ context.Context) ([]datamodel.Customer, error) {
	var customers []datamodel.Customer
	err := c.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(customerBucket)
		return b.ForEach(func(k, v []byte) error {
			customer, err := decodeCustomer(string(v))
			if err != nil {
				return err
			}
			customers = append(customers, customer)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return customers, nil
}
