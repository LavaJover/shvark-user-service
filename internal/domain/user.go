package domain

import "time"

type UserRole string

const (
	Trader 	 UserRole = "TRADER"
	Merchant UserRole = "MERCHANT"
	Admin	 UserRole = "ADMIN"
)

type User struct {
	ID				string
	Username		string
	Login			string
	Password		string
	Role 			UserRole
	TwoFaSecret		string
	TwoFaEnabled 	bool
	CreatedAt		time.Time
	UpdatedAt 		time.Time
}

func NewUser(username, login, password string) (*User, error) {
	return &User{
		Login: login,
		Username: username,
		Password: password,
	}, nil
}