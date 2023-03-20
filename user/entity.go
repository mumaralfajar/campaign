package user

import "time"

type User struct {
	ID                                                 int
	Name, Occupation, Email, PasswordHash, Role, Token string
	CreatedAt, UpdatedAt                               time.Time
}
