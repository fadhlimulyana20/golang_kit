-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id SERIAL,
    name VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
