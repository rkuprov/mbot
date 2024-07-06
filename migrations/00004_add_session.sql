-- +goose Up
-- +goose StatementBegin
CREATE TABLE session
(
    token      VARCHAR(255) NOT NULL PRIMARY KEY ,
    is_valid   BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP + INTERVAL '30 minutes'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE session;
-- +goose StatementEnd
