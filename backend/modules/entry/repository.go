package entry

import (
	"context"
	"database/sql"
	"fmt"
	"notes/infra"
)

func findEntries(ctx context.Context, userId string) ([]Entry, error) {
	entries := make([]Entry, 0)

	rows, err := infra.DbConn.QueryContext(ctx, `
		SELECT id, content, user_id, created_at, updated_at
		FROM entries
		WHERE user_id = $1
	`, userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var entry Entry

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

func findEntryById(ctx context.Context, id string, userId string) (*Entry, error) {
	entry := &Entry{}

	row := infra.DbConn.QueryRowContext(ctx, `
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

func updateEntry(ctx context.Context, id string, userId string, content string) error {
	_, err := infra.DbConn.ExecContext(ctx, `
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

func deleteEntry(ctx context.Context, id string) error {
	_, err := infra.DbConn.ExecContext(ctx, `DELETE FROM entries WHERE id = $1`, id)

	if err != nil {
		return err
	}

	return nil
}

func insertEntry(ctx context.Context, content string, userId string) error {

	fmt.Println("User id", userId)

	result, err := infra.DbConn.ExecContext(ctx, `INSERT INTO entries (content, user_id) VALUES($1, $2);`, content, userId)

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
