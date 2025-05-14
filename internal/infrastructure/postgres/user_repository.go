package postgres

import (
	"time"

	"github.com/LavaJover/shvark-user-service/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) (domain.UserRepository, error) {
	return &userRepository{
		db: db,
	}, nil
}

func (r *userRepository) CreateUser(user *domain.User) (string, error) {
	model := &UserModel{
		ID: uuid.New().String(),
		Login: user.Login,
		Username: user.Username,
		PasswordHash: user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := r.db.Create(model).Error
	if err == nil {
		user.ID = model.ID
	}
	return user.ID, err
}

func (r *userRepository) GetUserByID(userID string) (*domain.User, error) {
	var model UserModel
	if err := r.db.Where("id = ?", userID).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound{
			return nil, domain.ErrUserNotFound
		}
	}

	return &domain.User{
		ID: model.ID,
		Login: model.Login,
		Username: model.Username,
		Password: model.PasswordHash,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}