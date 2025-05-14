package domain

type UserRepository interface {
	CreateUser(login, username, password string) (string, error)
	GetUserByID(userID string) (*User, error)
}