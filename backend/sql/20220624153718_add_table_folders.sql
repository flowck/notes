-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS folders (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name varchar(100),
    created_at TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT now()
);

ALTER TABLE entries ADD COLUMN folder_id UUID;
ALTER TABLE entries ADD CONSTRAINT fk_folder_id
    FOREIGN KEY (folder_id)
    REFERENCES folders(id)
    ON DELETE SET NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE entries DROP CONSTRAINT fk_folder_id;
ALTER TABLE entries DROP COLUMN folder_id;
DROP TABLE folders;
-- +goose StatementEnd
