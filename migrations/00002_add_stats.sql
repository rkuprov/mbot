-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS stats
(
    customer_id INT NOT NULL references customers (id),
    subscription_id VARCHAR(36) NOT NULL references subscriptions (id),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX stats_customer_id_idx ON stats (customer_id);
CREATE INDEX stats_subscription_id_idx ON stats (subscription_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS stats_customer_id_idx;
DROP INDEX IF EXISTS stats_subscription_id_idx;
DROP TABLE IF EXISTS stats;
-- +goose StatementEnd
