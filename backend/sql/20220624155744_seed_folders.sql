-- +goose Up
-- +goose StatementBegin
INSERT INTO folders (name) VALUES ('Ideas');
INSERT INTO folders (name) VALUES ('Projects');
INSERT INTO folders (name) VALUES ('Goals');
INSERT INTO folders (name) VALUES ('Groceries');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FROM folders;
-- +goose StatementEnd
