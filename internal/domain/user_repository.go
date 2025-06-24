package domain

import "google.golang.org/protobuf/types/known/fieldmaskpb"

type UserRepository interface {
	CreateUser(*User) (string, error)
	GetUserByID(userID string) (*User, error)
	GetUserByLogin(login string) (*User, error)
	UpdateUser(userID string, user *User, mask *fieldmaskpb.FieldMask) (*User, error)
	GetUsers(page, limit int64) ([]*User, int64, error)

	SetTwoFaSecret(login, twoFaSecret string) error
}