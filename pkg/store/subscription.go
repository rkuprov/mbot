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

func (s *Store) CreateSubscription(ctx context.Context, customerID, startDate string, duration int) (string, error) {
	// ts, err := time.ParseInLocation(subscriptionFormat, startDate, time.FixedZone("MST", -7))
	// if err != nil {
	// 	return "", fmt.Errorf("failed to parse expiration date: %w", err)
	// }
	// expiry := ts.UTC().Add(time.Hour * 24 * time.Duration(duration)).Unix()
	// subID := uuid.New().String()
	// tx, err := s.db.Begin(true)
	// if err != nil {
	// 	return "", err
	// }
	// defer tx.Rollback()
	//
	// root := tx.Bucket(customersBucket)

	// return subID, nil
	return "", nil
}

func (s *Store) GetSubscription(ctx context.Context, id string) (*mbotpb.Subscription, error) {
	var exp int64
	err := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(subscriptions)
		expiryStr := b.Get([]byte(id))
		if expiryStr == nil {
			return fmt.Errorf("subscription not found")
		}
		var err error
		exp, err = strconv.ParseInt(string(expiryStr), 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse expiry: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &mbotpb.Subscription{
		SubscriptionId:     id,
		SubscriptionExpiry: timestamppb.New(time.Unix(exp, 0)),
	}, nil
}

func (s *Store) GetSubscriptionsAll(ctx context.Context) ([]*mbotpb.Subscription, error) {
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
