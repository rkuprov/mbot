package store

import (
	"context"

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

	err := s.pg.QueryRow(ctx, `
		SELECT 
			customers.id, 
			name, 
			email, 
			contact,
			subscriptions.id as subscription_id
	from customers 
	join subscriptions on subscriptions.customer_id = customers.id 
	and customers.id = $1
	`, id).Scan(
		&c.Id,
		&c.Name,
		&c.Email,
		&c.Contact,
		&c.SubscriptionIds,
	)
	if err != nil {
		return mbotpb.Customer{}, err
	}

	return c, nil
}

func (s *Store) GetCustomersAll(ctx context.Context) ([]mbotpb.Customer, error) {
	var customers []mbotpb.Customer

	rows, err := s.pg.Query(ctx, `
	SELECT
		id,
		name,
		email,
		contact
	FROM customers`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c mbotpb.Customer
		if err := rows.Scan(
			&c.Id,
			&c.Name,
			&c.Email,
			&c.Contact,
		); err != nil {
			return nil, err
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
