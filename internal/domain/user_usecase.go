package domain

import "google.golang.org/protobuf/types/known/fieldmaskpb"

type UserUsecase interface {
	GetUserByID(userID string) (*User, error)
	GetUserByLogin(login string) (*User, error)
	CreateUser(user *User) (string, error)
	UpdateUser(userID string, user *User, mask *fieldmaskpb.FieldMask) (*User, error)
}