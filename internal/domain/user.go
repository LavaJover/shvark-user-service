package domain

import "time"

type User struct {
	ID			string
	Username	string
	Login		string
	Password	string
	CreatedAt	time.Time
	UpdatedAt 	time.Time
}

func NewUser(username, login, password string) (*User, error) {
	return &User{
		Login: login,
		Username: username,
		Password: password,
	}, nil
}