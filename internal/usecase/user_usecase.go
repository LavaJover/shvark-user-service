package usecase

import "github.com/LavaJover/shvark-user-service/internal/domain"

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