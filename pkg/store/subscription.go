package store

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
)

const (
	subscriptionFormat = "2006-01-02"
)

var (
	subscriptionsBucket        = []byte("subscriptions")
	customerSubscriptionBucket = []byte("customer-subscription")
)

func (c *Client) CreateSubscription(ctx context.Context, customerSlug, startDate string, duration int) (string, error) {
	ts, err := time.ParseInLocation(subscriptionFormat, startDate, time.FixedZone("MST", -7))
	if err != nil {
		return "", fmt.Errorf("failed to parse expiration date: %w", err)
	}
	expiry := ts.UTC().Add(time.Hour * 24 * time.Duration(duration)).Unix()

	subID := uuid.New().String()
	err = c.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(subscriptionsBucket)
		err = b.Put([]byte(subID), []byte(strconv.Itoa(int(expiry))))
		if err != nil {
			return err
		}
		cs := tx.Bucket(customerSubscriptionBucket)
		current := cs.Get([]byte(customerSlug))
		updated := []byte(subID)
		if len(current) > 0 {
			updated = bytes.Join([][]byte{current, []byte(subID)}, []byte("#"))
		}

		return cs.Put([]byte(customerSlug), updated)
	})
	if err != nil {
		return "", err
	}

	fmt.Println("Created subscription", subID, "for customer", customerSlug, "with expiry", expiry)
	return subID, nil
}

func (c *Client) GetSubscription(ctx context.Context, id string) (*mbotpb.Subscription, error) {
	var exp int64
	err := c.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(subscriptionsBucket)
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
