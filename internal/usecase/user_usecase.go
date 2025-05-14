package usecase

import (
	"context"

	"github.com/LavaJover/shvark-user-service/internal/domain"
	"github.com/LavaJover/shvark-user-service/internal/infrastructure/kafka"
)

type UserUsecase struct {
	Repo domain.UserRepository
	producer kafka.EventProducer
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{Repo: repo}
}

func (uc *UserUsecase) CreateUser (user *domain.User) (string, error) {
	if err := uc.producer.Produce(context.Background(), "user.created", *user); err != nil {
		return "", err
	}
	return uc.Repo.CreateUser(user)
}

func (uc *UserUsecase) GetUserByID (userID string) (*domain.User, error) {
	return uc.Repo.GetUserByID(userID)
}