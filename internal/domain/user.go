package domain

import "time"

type Role string

const (
	Admin		Role = "admin"
	Lead		Role = "lead"
	TeamLead	Role = "team_lead"
	Trader		Role = "trader"
	Marketplace	Role = "marketplace"
)

type User struct {
	ID			string
	Username	string
	Login		string
	Password	string
	Role		Role
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