package grpcapi

import (
	"context"

	"github.com/LavaJover/shvark-user-service/internal/domain"
	userpb "github.com/LavaJover/shvark-user-service/proto/gen"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	domain.UserUsecase
}

func (h *UserHandler) CreateUser(ctx context.Context, r *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user := &domain.User{
		Username: r.Username,
		Login: r.Login,
		Password: r.Password,
	}
	userID, err := h.UserUsecase.CreateUser(user)
	return &userpb.CreateUserResponse{
		UserId: userID,
	}, err
}

func (h *UserHandler) GetUserByID(ctx context.Context, r *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, err := h.UserUsecase.GetUserByID(r.UserId)
	return &userpb.GetUserResponse{
		UserId: user.ID,
		Username: user.Username,
	}, err
}