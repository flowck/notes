package repositories

import (
	"context"
	"database/sql"
	"notes/internal/infra"
	"time"
)

type Entry struct {
	Id        string
	Content   string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FindEntries(ctx context.Context) ([]Entry, error) {
	entries := make([]Entry, 0)

	rows, err := infra.DbConn.QueryContext(ctx, `SELECT id, content, user_id, created_at, updated_at FROM entries`)

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

func FindEntryById(ctx context.Context, id string) (*Entry, error) {
	entry := &Entry{}

	row := infra.DbConn.QueryRowContext(ctx, `SELECT id, content, user_id, created_at, updated_at FROM entries WHERE id = $1`, id)
	err := row.Scan(&entry.Id, &entry.Content, &entry.UserId, &entry.CreatedAt, &entry.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return entry, nil
}

func UpdateEntry(ctx context.Context, id string, content string) error {

	_, err := infra.DbConn.ExecContext(ctx, `UPDATE entries SET content = $1 WHERE id = $2`, content, id)

	if err != nil {
		return err
	}

	return nil
}

func DeleteEntry(ctx context.Context, id string) error {
	_, err := infra.DbConn.ExecContext(ctx, `DELETE FROM entries WHERE id = $1`, id)

	if err != nil {
		return err
	}

	return nil
}

func InsertEntry(ctx context.Context, content string, userId string) error {

	_, err := infra.DbConn.ExecContext(ctx, `INSERT INTO entries (content, user_id) VALUES($1, $2);`, content, userId)

	if err != nil {
		return err
	}

	return nil
}
