package domain

type UserUsecase interface {
	GetUserByID(userID string) (*User, error)
	GetUserByLogin(login string) (*User, error)
	CreateUser(user *User) (string, error)
	UpdateUser(user *User) error
	CheckPermission(userID string, required Role) (bool, error)
}