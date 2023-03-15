-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS users
ADD is_verified BOOLEAN DEFAULT FALSE,
ADD verified_at TIMESTAMP WITHOUT TIME ZONE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS users
DROP COLUMN is_verified,
DROP COLUMN verified_at;
-- +goose StatementEnd
