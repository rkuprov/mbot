package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"

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

func (s *Store) CreateCustomer(ctx context.Context, in CustomerCreate) (string, error) {
	var id string
	err := s.pg.QueryRow(ctx, `INSERT INTO customers (name, email, contact) VALUES ($1, $2, $3) RETURNING id`,
		in.Name,
		in.Email,
		in.Contact,
	).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *Store) GetCustomer(ctx context.Context, id string) (mbotpb.Customer, error) {
	var c mbotpb.Customer
	var subs pgtype.Array[*string]
	err := s.pg.QueryRow(ctx, `
		SELECT 
	   id,
       name,
       email,
       contact
		from customers
where customers.id = $1 and customers.is_active = true
	`, id).Scan(
		&c.Id,
		&c.Name,
		&c.Email,
		&c.Contact,
	)
	if err != nil {
		return mbotpb.Customer{}, err
	}
	err = s.pg.QueryRow(ctx, `
		SELECT ARRAY_AGG(subscriptions.id) as subscription_ids
		FROM subscriptions
		WHERE subscriptions.customer_id = $1 and subscriptions.is_active = true
	`, id).Scan(&subs)

	for _, sub := range subs.Elements {
		if sub == nil {
			continue
		}
		c.SubscriptionIds = append(c.SubscriptionIds, *sub)
	}

	return c, nil
}

func (s *Store) GetCustomersAll(ctx context.Context) ([]mbotpb.Customer, error) {
	var customers []mbotpb.Customer

	rows, err := s.pg.Query(ctx, `
	SELECT
		customers.id,
		name,
		email,
		contact
	FROM customers
	WHERE customers.is_active = true 
	ORDER BY customers.id;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c mbotpb.Customer
		var subs pgtype.Array[*string]
		if err := rows.Scan(
			&c.Id,
			&c.Name,
			&c.Email,
			&c.Contact,
		); err != nil {
			return nil, err
		}
		err = s.pg.QueryRow(ctx, `
		SELECT ARRAY_AGG(subscriptions.id) as subscription_ids
		FROM subscriptions
		WHERE subscriptions.customer_id = $1 and subscriptions.is_active = true
		`, c.Id).Scan(&subs)
		if err != nil {
			return nil, err
		}

		for _, sub := range subs.Elements {
			if sub == nil {
				continue
			}
			c.SubscriptionIds = append(c.SubscriptionIds, *sub)
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (s *Store) UpdateCustomer(ctx context.Context, id string, in CustomerUpdate) error {
	_, err := s.pg.Exec(ctx, `UPDATE customers SET name = $1, email = $2, contact = $3 WHERE id = $4`,
		in.Name,
		in.Email,
		in.Contact,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteCustomer(ctx context.Context, id string) error {
	tx, err := s.pg.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	_, err = tx.Exec(ctx, `UPDATE customers SET is_active=false WHERE id = $1`, id)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, `UPDATE subscriptions SET is_active=false WHERE customer_id = $1`, id)
	if err != nil {
		return err
	}
	tx.Commit(ctx)

	return nil
}
