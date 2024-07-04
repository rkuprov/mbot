package store

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/rkuprov/mbot/pkg/gen/mbotpb"
)

const (
	subscriptionFormat = "2006-01-02"
)

var (
	// ErrSubscriptionNotFound is returned when a subscription is not found.
	ErrSubscriptionNotFound  = fmt.Errorf("subscription not found")
	ErrSubscriptionExpired   = fmt.Errorf("subscription expired")
	ErrSubscriptionNotActive = fmt.Errorf("subscription not active")
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
	id,
	start_date,
	expiration_date
	) VALUES ($1, $2, $3, $4) RETURNING id`,
		in.CustomerID,
		uuid.New().String(),
		in.StartDate,
		in.ExpirationDate,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *Store) GetSubscription(ctx context.Context, id string) (*mbotpb.Subscription, error) {
	var sub mbotpb.Subscription
	var startDate time.Time
	var expirationDate time.Time
	err := s.pg.QueryRow(ctx, `
   		SELECT
   			id,
   			customer_id,
   			start_date,	
   			expiration_date	
   		FROM subscriptions
   		WHERE id = $1 and is_active = true
   `, id).Scan(
		&sub.SubscriptionId,
		&sub.CustomerId,
		&startDate,
		&expirationDate,
	)
	if err != nil {
		return nil, err
	}
	sub.StartDate = timestamppb.New(startDate)
	sub.ExpirationDate = timestamppb.New(expirationDate)

	return &sub, nil
}

func (s *Store) GetSubscriptionsAll(ctx context.Context) ([]*mbotpb.Subscription, error) {
	var out []*mbotpb.Subscription
	rows, err := s.pg.Query(ctx, `SELECT id, customer_id, start_date, expiration_date FROM subscriptions ORDER BY customer_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var customerID int
		var startDate time.Time
		var expirationDate time.Time
		if err = rows.Scan(&id, &customerID, &startDate, &expirationDate); err != nil {
			return nil, err
		}
		out = append(out, &mbotpb.Subscription{
			SubscriptionId: id,
			CustomerId:     strconv.Itoa(customerID),
			StartDate:      timestamppb.New(startDate),
			ExpirationDate: timestamppb.New(expirationDate),
		})
	}

	return out, nil
}

func (s *Store) GetSubscriptionByCustomer(ctx context.Context, customerID string) ([]*mbotpb.Subscription, error) {
	var out []*mbotpb.Subscription
	rows, err := s.pg.Query(ctx, `SELECT id, start_date, expiration_date FROM subscriptions 
	WHERE customer_id = $1 ORDER BY start_date`, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sub mbotpb.Subscription
		var startDate time.Time
		var expirationDate time.Time
		if err = rows.Scan(&sub.SubscriptionId, &startDate, &expirationDate); err != nil {
			return nil, err
		}
		sub.StartDate = timestamppb.New(startDate)
		sub.ExpirationDate = timestamppb.New(expirationDate)
		out = append(out, &sub)
	}

	return out, nil
}

func (s *Store) DeleteSubscription(ctx context.Context, id string) error {
	_, err := s.pg.Exec(ctx, `
		UPDATE subscriptions SET is_active = false WHERE id = $1
	`, id)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSubscription will perform up update.
func (s *Store) UpdateSubscription(ctx context.Context, in SubscriptionUpdate) (*mbotpb.UpdateSubscriptionResponse, error) {
	var resp mbotpb.UpdateSubscriptionResponse
	sub, err := s.GetSubscription(ctx, in.SubscriptionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscription: %w", err)
	}
	resp.StartDate = sub.StartDate
	resp.ExpirationDate = sub.ExpirationDate

	if err = validUpdate(sub, in); err != nil {
		return nil, fmt.Errorf("invalid update: %w", err)
	}

	_, err = s.pg.Exec(ctx, `
		UPDATE subscriptions
		SET
			start_date = $1,
			expiration_date = $2
		WHERE id = $3
	`, in.StartDate.AsTime(), in.ExpirationDate.AsTime(), in.SubscriptionID)
	if err != nil {
		return nil, fmt.Errorf("failed to update subscription: %w", err)
	}
	resp.UpdatedExpirationDate = in.ExpirationDate
	resp.UpdatedStartDate = in.StartDate

	return &resp, nil
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

func (s *Store) ConfirmSubscription(ctx context.Context, token string) error {
	var startDate time.Time
	var expirationDate time.Time
	var id string
	err := s.pg.QueryRow(ctx, `
	SELECT 
	id,
	start_date,
	expiration_date
	FROM subscriptions 
	WHERE id = $1 
	`,
		token,
	).Scan(&id, &startDate, &expirationDate)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrSubscriptionNotFound
	}
	if err != nil {
		return err
	}
	if time.Now().Before(startDate) {
		return ErrSubscriptionNotActive
	}
	if time.Now().After(expirationDate) {
		return ErrSubscriptionExpired
	}

	return nil
}
