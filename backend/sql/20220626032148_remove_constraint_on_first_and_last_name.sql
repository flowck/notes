-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ALTER COLUMN first_name DROP NOT NULL;
ALTER TABLE users ALTER COLUMN last_name DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users ALTER COLUMN first_name VARCHAR(100) SET NOT NULL;
ALTER TABLE users ALTER COLUMN last_name VARCHAR(100) SET NOT NULL;
-- +goose StatementEnd
