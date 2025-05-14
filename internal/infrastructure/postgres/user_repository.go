package postgres

import (
	"gorm.io/gorm"
)

type userRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) (*domain.UserRepository, error) {
	return &userRepository{
		DB: db,
	}, nil
}

func (r *userRepository) CreateUser(login, username, password string) (string, error) {
	model := &UserModel{
		
	}
}

func (r *userRepository) GetUserByID(userID string) (*domain.User, error) {

}