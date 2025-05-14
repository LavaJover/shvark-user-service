package domain

type UserRepository interface {
	CreateUser(*User) (string, error)
	GetUserByID(userID string) (*User, error)
}