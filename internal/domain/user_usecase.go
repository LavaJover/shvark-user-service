package domain

type UserUsecase interface {
	GetUserByID(userID string) (*User, error)
	CreateUser(user *User) (string, error)
}