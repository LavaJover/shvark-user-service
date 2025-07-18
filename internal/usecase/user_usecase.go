package usecase

import (
	"github.com/LavaJover/shvark-user-service/internal/domain"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

type UserUsecase struct {
	Repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{repo}
}

func (uc *UserUsecase) CreateUser(user *domain.User) (string, error) {
	return uc.Repo.CreateUser(user)
}

func (uc *UserUsecase) GetUserByID(userID string) (*domain.User, error) {
	return uc.Repo.GetUserByID(userID)
}

func (uc *UserUsecase) GetUserByLogin(login string) (*domain.User, error) {
	return uc.Repo.GetUserByLogin(login)
}

func (uc *UserUsecase) UpdateUser(userID string, user *domain.User, mask *fieldmaskpb.FieldMask) (*domain.User, error) {
	return uc.Repo.UpdateUser(userID, user, mask)
}

func (uc *UserUsecase) GetUsers(page, limit int64) ([]*domain.User, int64, error) {
	return uc.Repo.GetUsers(page, limit)
}

func (uc *UserUsecase) SetTwoFaSecret(userID, twoFaSecret string) error {
	return uc.Repo.SetTwoFaSecret(userID, twoFaSecret)
}

func (uc *UserUsecase) GetTwoFaSecretByID(userID string) (string, error) {
	return uc.Repo.GetTwoFaSecretByID(userID)
}

func (uc *UserUsecase) SetTwoFaEnabled(userID string, enabled bool) error {
	return uc.Repo.SetTwoFaEnabled(userID, enabled)
}

func (uc *UserUsecase) GetTraders() ([]*domain.User, error) {
	return uc.Repo.GetTraders()
}

func (uc *UserUsecase) GetMerchants() ([]*domain.User, error) {
	return uc.Repo.GetMerchants()
}