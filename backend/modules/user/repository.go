package user

import (
	"context"
	"database/sql"
	"log"
	"notes/infra"
)

func findUserByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{}
	row := infra.DbConn.QueryRowContext(ctx, `
		SELECT id, COALESCE(first_name, '') as first_name, COALESCE(last_name, '') as last_name
		FROM users
		WHERE email = $1
	`, email)

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Printf("Unable to query user: %s", err)
		return nil, err
	}

	return user, nil
}

func insertUser(ctx context.Context, email string, password string) error {
	_, err := infra.DbConn.ExecContext(ctx, `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
	`, email, password)

	if err != nil {
		log.Println("Unable to insert user: %w", err)
		return err
	}

	return nil
}
