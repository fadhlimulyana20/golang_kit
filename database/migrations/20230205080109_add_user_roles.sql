-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_roles(
    user_id INT4,
    role_id INT4,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_roles FOREIGN KEY (role_id) REFERENCES roles(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_roles
-- +goose StatementEnd
