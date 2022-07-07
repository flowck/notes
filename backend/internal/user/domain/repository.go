package domain

import "context"

type UserRepository interface {
	InsertUser(ctx context.Context, email string, password string) error
	FindUserByEmail(ctx context.Context, email string) (*User, error)
}
