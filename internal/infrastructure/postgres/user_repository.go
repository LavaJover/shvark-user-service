package postgres

import (
	"github.com/LavaJover/shvark-user-service/internal/domain"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
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
	}, nil
}

func (r *userRepository) GetUserByLogin(login string) (*domain.User, error) {
	var model UserModel
	if err := r.db.Where("login = ?", login).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return &domain.User{
		ID: model.ID,
		Username: model.Username,
		Login: model.Login,
		Password: model.PasswordHash,
	}, nil
}

func (r *userRepository) UpdateUser(userID string, user *domain.User, mask *fieldmaskpb.FieldMask) (*domain.User, error) {
	// find user to update from db
	dbUser, err := r.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// change fields from mask
	for _, path := range mask.Paths {
		switch path {
		case "username":
			dbUser.Username = user.Username
		case "login":
			dbUser.Login = user.Login
		}
	}

	res := r.db.Save(
		&UserModel{
			ID: dbUser.ID,
			Login: dbUser.Login,
			Username: dbUser.Username,
			PasswordHash: dbUser.Password,
		},
	)

	if res.Error != nil {
		return nil, err
	}

	return dbUser, nil
}

func (r *userRepository) GetUsers(page, limit int64) ([]*domain.User, int64, error) {
	var userModels []UserModel
	var total int64

	if err := r.db.Model(&UserModel{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page-1) * limit
	totalPages := (total + limit - 1) / limit
	if err := r.db.Offset(int(offset)).Limit(int(limit)).Order("created_at DESC").Find(&userModels).Error; err != nil {
		return nil, 0, err
	}

	var userRecords []*domain.User
	for _, userModel := range userModels {
		userRecords = append(userRecords, &domain.User{
			ID: userModel.ID,
			Username: userModel.Username,
			Login: userModel.Login,
			Password: userModel.PasswordHash,
			CreatedAt: userModel.CreatedAt,
			UpdatedAt: userModel.UpdatedAt,
		})
	}

	return userRecords, totalPages, nil
}