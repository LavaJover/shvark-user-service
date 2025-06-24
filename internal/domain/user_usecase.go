package domain

import "google.golang.org/protobuf/types/known/fieldmaskpb"

type UserUsecase interface {
	GetUserByID(userID string) (*User, error)
	GetUserByLogin(login string) (*User, error)
	CreateUser(user *User) (string, error)
	UpdateUser(userID string, user *User, mask *fieldmaskpb.FieldMask) (*User, error)
	GetUsers(page, limit int64) ([]*User, int64, error)
	SetTwoFaSecret(userID, twoFaSecret string) error
	GetTwoFaSecretByID(userID string) (string, error)
	SetTwoFaEnabled(userID string, enabled bool) error
}