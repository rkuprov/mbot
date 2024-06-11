package store

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go.etcd.io/bbolt"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
)

const (
	subscriptionFormat = "2006-01-02"
)

var (
	customer_subscriptions = []byte("customer-subscriptions")
	subscription_customer  = []byte("subscription-customer")
)

type SubscriptionCreate struct {
	CustomerID string
	StartDate  string
	Duration   int
}

type SubscriptionUpdate struct {
	SubscriptionID string
	StartDate      string
	Duration       int
}

// CreateSubscription creates a new subscription for a customer in the Customers bucket as well as in the SubscriptionsLUT one
// for a quicker lookup from the endpoint.
func (s *Store) CreateSubscription(_ context.Context, in SubscriptionCreate) (string, error) {
	cId := []byte(in.CustomerID)
	var id string

	t, err := time.Parse(subscriptionFormat, in.StartDate)
	if err != nil {
		return "", fmt.Errorf("failed to parse start date: %w", err)
	}
	expiration := t.Add(time.Duration(in.Duration) * 24 * time.Hour)

	err = s.db.Update(func(tx *bbolt.Tx) error {
		var uid uint64
		uid, err = tx.Bucket(subscriptions).NextSequence()
		id = strconv.FormatUint(uid, 10)

		err = tx.Bucket(customersBucket).Bucket(cId).Bucket(subscriptions).Put(
			[]byte(id),
			[]byte(strconv.FormatInt(expiration.Unix(), 10)),
		)
		if err != nil {
			return fmt.Errorf("failed to create subscription: %w", err)
		}

		err = tx.Bucket(subscription_customer).Put([]byte(id), cId)
		if err != nil {
			return fmt.Errorf("failed to create subscription-customer association: %w", err)
		}

		return tx.Bucket(subscriptions).Put([]byte(id), []byte((strconv.FormatInt(expiration.Unix(), 10))))
	})

	return id, nil
}

func (s *Store) GetSubscription(_ context.Context, id string) (*mbotpb.Subscription, error) {
	var out *mbotpb.Subscription
	err := s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(subscriptions).Get([]byte(id))
		if v == nil {
			return fmt.Errorf("subscription not found")
		}
		expiration, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse expiration: %w", err)
		}
		out = &mbotpb.Subscription{
			SubscriptionId:     id,
			SubscriptionExpiry: timestamppb.New(time.Unix(expiration, 0)),
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *Store) GetSubscriptionsAll(_ context.Context) ([]*mbotpb.Subscription, error) {
	var out []*mbotpb.Subscription
	err := s.db.View(func(tx *bbolt.Tx) error {
		cs := tx.Bucket(subscriptions)
		err := cs.ForEach(func(k, v []byte) error {
			expiry, err := strconv.ParseInt(string(v), 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse expiry: %w", err)
			}
			out = append(out, &mbotpb.Subscription{
				SubscriptionId:     string(k),
				SubscriptionExpiry: timestamppb.New(time.Unix(expiry, 0)),
			})
			return nil
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *Store) GetSubscriptionByCustomer(_ context.Context, customerID string) ([]*mbotpb.Subscription, error) {
	var out []*mbotpb.Subscription
	var err error

	err = s.db.View(func(tx *bbolt.Tx) error {
		c := tx.Bucket(customersBucket).Bucket([]byte(customerID)).Bucket(subscriptions)
		err = c.ForEach(func(k, v []byte) error {
			expiry, err := strconv.ParseInt(string(v), 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse expiry: %w", err)
			}
			out = append(out, &mbotpb.Subscription{
				SubscriptionId:     string(k),
				SubscriptionExpiry: timestamppb.New(time.Unix(expiry, 0)),
			})
			return nil
		})

		return err
	})
	if err != nil {
		return nil, err
	}

	return out, nil
}

// not implemented
func (s *Store) DeleteSubscription(_ context.Context, id, subId string) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		err := tx.Bucket(subscriptions).Delete([]byte(subId))
		if err != nil {
			return fmt.Errorf("failed to delete from subscriptions: %w", err)
		}
		err = tx.Bucket(customersBucket).Bucket([]byte(id)).Bucket(subscriptions).Delete([]byte(subId))
		if err != nil {
			return fmt.Errorf("failed to delete from customer: %w", err)
		}
		return nil
	})
}

func (s *Store) UpdateSubscription(_ context.Context, in SubscriptionUpdate) error {
	t, err := time.Parse(subscriptionFormat, in.StartDate)
	if err != nil {
		return fmt.Errorf("failed to parse start date: %w", err)
	}
	expiry := t.Add(time.Duration(in.Duration) * 24 * time.Hour)
	var cID string
	err = s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(subscription_customer).Get([]byte(in.SubscriptionID))
		if v == nil {
			return fmt.Errorf("subscription not found")
		}
		cID = string(v)

		return nil
	})

	return s.db.Update(func(tx *bbolt.Tx) error {
		err = tx.Bucket(subscriptions).Put([]byte(in.SubscriptionID), []byte(strconv.FormatInt(expiry.Unix(), 10)))
		if err != nil {
			return fmt.Errorf("failed to update subscription: %w", err)
		}
		err = tx.Bucket(customersBucket).Bucket([]byte(cID)).Bucket(subscriptions).Put([]byte(in.SubscriptionID),
			[]byte(strconv.FormatInt(expiry.Unix(), 10)))
		if err != nil {
			return fmt.Errorf("failed to update customer: %w", err)
		}

		return nil
	})
}
