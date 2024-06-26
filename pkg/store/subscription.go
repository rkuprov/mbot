package store

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
)

const (
	subscriptionFormat = "2006-01-02"
)

type SubscriptionCreate struct {
	CustomerID     string
	StartDate      time.Time
	ExpirationDate time.Time
}

type SubscriptionUpdate struct {
	SubscriptionID string
	StartDate      *timestamppb.Timestamp
	ExpirationDate *timestamppb.Timestamp
}

// CreateSubscription creates a new subscription for a customer in the Customers bucket as well as in the SubscriptionsLUT one
// for a quicker lookup from the endpoint.
func (s *Store) CreateSubscription(ctx context.Context, in SubscriptionCreate) (string, error) {
	var id string
	err := s.pg.QueryRow(ctx, `INSERT INTO subscriptions (
	customer_id,
	token,
	start_date,
	end_date
	) VALUES ($1, $2, $3, $4) RETURNING id`,
		in.CustomerID,
		uuid.New().String(),
		in.StartDate,
		in.StartDate,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *Store) GetSubscription(ctx context.Context, id string) (*mbotpb.Subscription, error) {
	var sub mbotpb.Subscription
	err := s.pg.QueryRow(ctx, `
   		SELECT
   			id,
   			customer_id,
   			start_date,	
   			end_date	
   		FROM subscriptions
   		WHERE id = $1
   `, id).Scan(
		&sub.SubscriptionId,
		&sub.StartDate,
		&sub.ExpirationDate,
	)
	if err != nil {
		return nil, err
	}

	return &sub, nil
}

func (s *Store) GetSubscriptionsAll(ctx context.Context) ([]*mbotpb.Subscription, error) {
	var out []*mbotpb.Subscription
	rows, err := s.pg.Query(ctx, `SELECT id, customer_id, start_date, end_date FROM subscriptions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sub mbotpb.Subscription
		if err = rows.Scan(&sub.SubscriptionId, &sub.StartDate, &sub.ExpirationDate); err != nil {
			return nil, err
		}
		out = append(out, &sub)
	}

	return out, nil
}

func (s *Store) GetSubscriptionByCustomer(ctx context.Context, customerID string) ([]*mbotpb.Subscription, error) {
	var out []*mbotpb.Subscription
	rows, err := s.pg.Query(ctx, `SELECT id, start_date, end_date FROM subscriptions WHERE customer_id = $1`, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sub mbotpb.Subscription
		if err = rows.Scan(&sub.SubscriptionId, &sub.StartDate, &sub.ExpirationDate); err != nil {
			return nil, err
		}
		out = append(out, &sub)
	}

	return out, nil
}

// not implemented
func (s *Store) DeleteSubscription(_ context.Context, id, subId string) error {
	// return s.db.Update(func(tx *bbolt.Tx) error {
	// 	err := tx.Bucket(subscriptions).Delete([]byte(subId))
	// 	if err != nil {
	// 		return fmt.Errorf("failed to delete from subscriptions: %w", err)
	// 	}
	// 	err = tx.Bucket(customersBucket).Bucket([]byte(id)).Bucket(subscriptions).Delete([]byte(subId))
	// 	if err != nil {
	// 		return fmt.Errorf("failed to delete from customer: %w", err)
	// 	}
	return nil
}

// UpdateSubscription will perform up update.
func (s *Store) UpdateSubscription(ctx context.Context, in SubscriptionUpdate) error {
	sub, err := s.GetSubscription(ctx, in.SubscriptionID)
	if err != nil {
		return fmt.Errorf("failed to get subscription: %w", err)
	}

	if err = validUpdate(sub, in); err != nil {
		return fmt.Errorf("invalid update: %w", err)
	}

	_, err = s.pg.Exec(ctx, `
		UPDATE subscriptions
		SET
			start_date = $1,
			end_date = $2
		WHERE id = $3
	`, in.StartDate, in.ExpirationDate, in.SubscriptionID)
	if err != nil {
		return fmt.Errorf("failed to update subscription: %w", err)
	}

	return nil
}

func validUpdate(sub *mbotpb.Subscription, in SubscriptionUpdate) error {
	if in.StartDate != nil {
		if in.StartDate.AsTime().Before(time.Now()) {
			return fmt.Errorf("start date cannot be in the past")
		}
		if sub.StartDate.AsTime().Before(time.Now()) {
			return fmt.Errorf("subscirption is already active")
		}
	}

	if in.ExpirationDate != nil {
		if in.ExpirationDate.AsTime().Before(sub.StartDate.AsTime()) {
			return fmt.Errorf("expiration date cannot be before start date")
		}
	}

	if in.StartDate != nil && in.ExpirationDate != nil {
		if in.StartDate.AsTime().After(in.ExpirationDate.AsTime()) {
			return fmt.Errorf("start date cannot be after expiration date")
		}
	}

	return nil
}
