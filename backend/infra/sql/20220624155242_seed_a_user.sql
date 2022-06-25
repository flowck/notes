-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id, first_name, last_name, email, password)
    VALUES ('fe1433f8-8576-4e04-87df-031778028bd5', 'Firmino', 'Changani', 'firmino.changani', 'adminadmin');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users;
-- +goose StatementEnd
