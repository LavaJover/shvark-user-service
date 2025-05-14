package repository

import "github.com/LavaJover/shvark-user-service/internal/domain/model"

type UserRepository interface {
	GetUserByID(id string) (*model.User, error)
}