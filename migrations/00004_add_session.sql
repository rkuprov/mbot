-- +goose Up
-- +goose StatementBegin
CREATE TABLE session
(
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER      NOT NULL,
    token      VARCHAR(255) NOT NULL,
    is_valid   BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE session;
-- +goose StatementEnd
