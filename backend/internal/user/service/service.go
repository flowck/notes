package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"notes/infra"
	"notes/internal/user/domain"
	"time"
)

type UserService struct {
	repository domain.UserRepository
}

var UserNotFound error = errors.New("user was not found")
var PasswordDoesntMatch error = errors.New("password provided doesn't match")

func NewUserService(repository domain.UserRepository) UserService {
	if repository == nil {
		panic("User repository is missing.")
	}

	return UserService{repository}
}

func (s *UserService) SignUp(ctx context.Context, email string, plainTextPassword string) error {
	// Check if user exists
	existentUser, err := s.repository.FindUserByEmail(ctx, email)

	if err != nil {
		panic(err)
	}

	if existentUser != nil {
		return UserNotFound
	}

	// Hash the plainTextPassword
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), bcrypt.DefaultCost)

	err = s.repository.InsertUser(ctx, email, string(hashedPassword))

	if err != nil {
		panic(err)
	}

	return nil
}

func (s *UserService) SignIn(ctx context.Context, email string, plainTextPassword string) (string, error) {
	user, err := s.repository.FindUserByEmail(ctx, email)

	if err != nil {
		panic(err)
	}

	if user == nil {
		return "", UserNotFound
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plainTextPassword))

	if err != nil {
		return "", PasswordDoesntMatch
	}

	claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 5 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, signErr := token.SignedString([]byte(infra.Cfg.JwtSigningKey))

	if signErr != nil {
		panic(err)
	}

	return ss, nil
}
