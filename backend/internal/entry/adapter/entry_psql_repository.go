package adapter

import (
	"context"
	"database/sql"
	"fmt"
	entry "notes/internal/entry/domain"
)

type EntryPsqlRepository struct {
	client *sql.DB
}

func NewEntryPsqlRepository(Client *sql.DB) EntryPsqlRepository {
	return EntryPsqlRepository{Client}
}

func (r EntryPsqlRepository) FindEntries(ctx context.Context, userId string) ([]entry.Entry, error) {
	entries := make([]entry.Entry, 0)

	rows, err := r.client.QueryContext(ctx, `
		SELECT id, content, user_id, created_at, updated_at
		FROM entries
		WHERE user_id = $1
	`, userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var entry entry.Entry

		rows.Scan(
			&entry.Id,
			&entry.Content,
			&entry.UserId,
			&entry.CreatedAt,
			&entry.UpdatedAt,
		)

		entries = append(entries, entry)
	}

	return entries, nil
}

func (r EntryPsqlRepository) FindEntryById(ctx context.Context, userId string, id string) (*entry.Entry, error) {
	entry := &entry.Entry{}

	row := r.client.QueryRowContext(ctx, `
		SELECT id, content, user_id, created_at, updated_at
		FROM entries
		WHERE id = $1 AND user_id = $2
	`, id, userId)
	err := row.Scan(&entry.Id, &entry.Content, &entry.UserId, &entry.CreatedAt, &entry.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (r EntryPsqlRepository) UpdateEntry(ctx context.Context, userId string, id string, content string) error {
	_, err := r.client.ExecContext(ctx, `
		UPDATE entries
		SET content = $1, updated_at = now()
		WHERE id = $2
		AND user_id = $3
	`, content, id, userId)

	if err != nil {
		return err
	}

	return nil
}

func (r EntryPsqlRepository) DeleteEntry(ctx context.Context, userId string, id string) error {
	_, err := r.client.ExecContext(ctx, `DELETE FROM entries WHERE id = $1 AND user_id = $2`, id, userId)

	if err != nil {
		return err
	}

	return nil
}

func (r EntryPsqlRepository) InsertEntry(ctx context.Context, userId string, content string) error {
	result, err := r.client.ExecContext(ctx, `INSERT INTO entries (content, user_id) VALUES($1, $2);`, content, userId)

	if err != nil {
		return err
	}

	row, err := result.RowsAffected()

	if err != nil {
		fmt.Errorf("Unable to get the rows affected: %w", err)
		return err
	}

	fmt.Println(row)

	return nil
}
