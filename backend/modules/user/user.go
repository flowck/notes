package user

import "time"

type User struct {
	Id                           string
	FirstName                    string
	LastName                     string
	Password                     string
	Email                        string
	PasswordResetToken           string
	PasswordResetTokenExpiryDate time.Time
	createdAt                    time.Time
	updatedAt                    time.Time
}
