-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS customers
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL,
    contact    TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS subscriptions
(
    id              SERIAL PRIMARY KEY,
    customer_id     INT          NOT NULL references customers (id),
    token           VARCHAR(255) NOT NULL UNIQUE,
    start_date      DATE         NOT NULL,
    expiration_date DATE         NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS subscriptions;
DROP TABLE IF EXISTS customers;
-- +goose StatementEnd
