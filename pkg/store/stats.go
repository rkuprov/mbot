package store

import "context"

func (s *Store) RecordUsage(ctx context.Context, subID string) error {
	_, err := s.pg.Exec(ctx, `
	INSERT INTO stats (customer_id, subscription_id)
	SELECT 
	customers.id as customer_id,
	subscriptions.id as subscription_id
	FROM customers join public.subscriptions on customers.id = subscriptions.customer_id
	WHERE subscriptions.id = $1
	`, subID)
	if err != nil {
		return err
	}

	return nil
}
