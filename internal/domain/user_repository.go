package domain

type UserRepository interface {
	CreateUser(*User) (string, error)
	GetUserByID(userID string) (*User, error)
	GetUserByLogin(login string) (*User, error)
	UpdateUser(user *User) error
}