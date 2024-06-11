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
	subscriptions = []byte("subscriptions")
)

// CreateSubscription creates a new subscription for a customer in the Customers bucket as well as in the SubscriptionsLUT one
// for a quicker lookup from the endpoint.
func (s *Store) CreateSubscription(_ context.Context, customerID, startDate string, duration int) (string, error) {
	cId := []byte(customerID)
	var id string

	t, err := time.Parse(subscriptionFormat, startDate)
	if err != nil {
		return "", fmt.Errorf("failed to parse start date: %w", err)
	}
	expiry := t.Add(time.Duration(duration) * 24 * time.Hour)

	err = s.db.Update(func(tx *bbolt.Tx) error {
		var uid uint64
		uid, err = tx.Bucket(customersBucket).Bucket(cId).Bucket(subscriptions).NextSequence()
		id = strconv.FormatUint(uid, 10)

		err = tx.Bucket(customersBucket).Bucket(cId).Bucket(subscriptions).Put(
			[]byte(id),
			[]byte(strconv.FormatInt(expiry.Unix(), 10)),
		)
		if err != nil {
			return fmt.Errorf("failed to create subscription: %w", err)
		}

		return tx.Bucket(subscriptionsLUT).Put([]byte(id), []byte((strconv.FormatInt(expiry.Unix(), 10))))
	})

	return id, nil
}

func (s *Store) GetSubscription(_ context.Context, id string) (*mbotpb.Subscription, error) {
	var out *mbotpb.Subscription
	err := s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(subscriptionsLUT).Get([]byte(id))
		if v == nil {
			return fmt.Errorf("subscription not found")
		}
		expiry, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse expiry: %w", err)
		}
		out = &mbotpb.Subscription{
			SubscriptionId:     id,
			SubscriptionExpiry: timestamppb.New(time.Unix(expiry, 0)),
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
		cs := tx.Bucket(subscriptionsLUT)
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
