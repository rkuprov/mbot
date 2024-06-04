package store

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
)

var (
	customerBucket                   = []byte("customers")
	customerSlugBucket               = []byte("customer-slug")
	customerDeleteBucket             = []byte("customer-delete")
	customerSlugDeleteBucket         = []byte("customer-slug-delete")
	customerSubscriptionDeleteBucket = []byte("customer-subscription-delete")
)

func (c *Client) CreateCustomer(_ context.Context, name, email, contact string) (string, error) {
	id := uuid.New().String()
	slug := generateSlug(name, id)
	err := c.db.Update(func(tx *bbolt.Tx) error {
		err := tx.Bucket(customerBucket).Put([]byte(id), encodeCustomer(mbotpb.Customer{
			Slug:    slug,
			Name:    name,
			Email:   email,
			Contact: contact,
		}))
		if err != nil {
			return err
		}
		err = tx.Bucket(customerSlugBucket).Put([]byte(slug), []byte(id))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return slug, nil
}

func encodeCustomer(c mbotpb.Customer) []byte {
	in := fmt.Sprintf("%s#%s#%s#%s", c.Slug, c.Name, c.Email, c.Contact)

	return []byte(in)
}

func decodeCustomer(in string) (mbotpb.Customer, error) {
	parts := strings.Split(in, "#")
	if len(parts) < 4 {
		return mbotpb.Customer{}, fmt.Errorf("invalid customer data: %s", in)
	}

	return mbotpb.Customer{
		Slug:    parts[0],
		Name:    parts[1],
		Email:   parts[2],
		Contact: parts[3],
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

func (c *Client) GetCustomer(_ context.Context, slug string) (mbotpb.Customer, error) {
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
	subscriptions := []string{}
	err = c.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(customerBucket)
		v := b.Get(id)
		if v == nil {
			return fmt.Errorf("customer with ID %s not found", string(id))
		}

		encoded = string(v)
		subs := tx.Bucket(customerSubscriptionBucket)
		subIds := subs.Get(id)
		if len(subIds) > 0 {
			subscriptions = strings.Split(string(subIds), "#")
		}
		return nil
	})

	if err != nil {
		return mbotpb.Customer{}, err
	}

	out, err := decodeCustomer(encoded)
	if err != nil {
		return mbotpb.Customer{}, err
	}
	out.SubscriptionIds = subscriptions
	return out, nil
}

func (c *Client) GetCustomersAll(_ context.Context) ([]mbotpb.Customer, error) {
	var customers []mbotpb.Customer
	err := c.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(customerBucket)
		ids := []string{}

		err := b.ForEach(func(k, v []byte) error {
			customer, err := decodeCustomer(string(v))
			if err != nil {
				return err
			}
			customers = append(customers, customer)
			ids = append(ids, string(k))
			return nil
		})
		if err != nil {
			return err
		}

		subs := tx.Bucket(customerSubscriptionBucket)
		for i, id := range ids {
			subIds := subs.Get([]byte(id))
			if len(subIds) > 0 {
				customers[i].SubscriptionIds = strings.Split(string(subIds), "#")
			}

		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (c *Client) UpdateCustomer(_ context.Context, slug, name, email, contact string) error {
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
	if err != nil {
		return err
	}

	err = c.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(customerBucket)
		customer, err := decodeCustomer(string(b.Get(id)))
		if err != nil {
			return err
		}
		customer.Name = name
		customer.Email = email
		customer.Contact = contact
		return b.Put(id, encodeCustomer(customer))
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteCustomer(_ context.Context, slug string) error {
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
	if err != nil {
		return err
	}

	err = c.db.Update(func(tx *bbolt.Tx) error {
		// delete customer
		b := tx.Bucket(customerBucket)
		delC := b.Get(id)
		db := tx.Bucket(customerDeleteBucket)
		err = db.Put(id, delC)
		if err != nil {
			return err
		}
		err = b.Delete(id)
		if err != nil {
			return err
		}

		// delete subscription
		subs := tx.Bucket(customerSubscriptionBucket)
		subIds := subs.Get(id)
		if len(subIds) > 0 {
			dSubs := tx.Bucket(customerSubscriptionDeleteBucket)
			err = dSubs.Put(id, subIds)
			if err != nil {
				return err
			}
			err = subs.Delete(id)
			if err != nil {
				return err
			}
		}
		err = subs.Delete(id)
		if err != nil {
			return err
		}

		// delete slug
		slugB := tx.Bucket(customerSlugBucket)
		delID := slugB.Get([]byte(slug))
		err = tx.Bucket(customerSlugDeleteBucket).Put([]byte(slug), delID)
		if err != nil {
			return err
		}
		err = slugB.Delete([]byte(slug))
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
