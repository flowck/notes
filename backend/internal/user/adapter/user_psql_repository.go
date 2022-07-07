package adapter

import (
	"context"
	"database/sql"
	"log"
	"notes/internal/user/domain"
)

type UserPsqlRepository struct {
	client *sql.DB
}

func NewUserPsqlRepository(client *sql.DB) UserPsqlRepository {
	if client == nil {
		panic("Database client is missing")
	}

	return UserPsqlRepository{client}
}

func (r UserPsqlRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	row := r.client.QueryRowContext(ctx, `
		SELECT id, COALESCE(first_name, '') as first_name, COALESCE(last_name, '') as last_name, password
		FROM users
		WHERE email = $1
	`, email)

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Password)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Printf("Unable to query user: %s", err)
		return nil, err
	}

	return user, nil
}
func (r UserPsqlRepository) InsertUser(ctx context.Context, email string, password string) error {
	_, err := r.client.ExecContext(ctx, `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
	`, email, password)

	if err != nil {
		log.Println("Unable to insert user: %w", err)
		return err
	}

	return nil
}
