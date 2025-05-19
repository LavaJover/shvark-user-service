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

func (h *UserHandler) GetUserByID(ctx context.Context, r *userpb.GetUserByIDRequest) (*userpb.GetUserByIDResponse, error) {
	user, err := h.UserUsecase.GetUserByID(r.UserId)
	return &userpb.GetUserByIDResponse{
		Login: user.Login,
		UserId: user.ID,
		Username: user.Username,
		Password: user.Password,
	}, err
}

func (h *UserHandler) GetUserByLogin(ctx context.Context, r *userpb.GetUserByLoginRequest) (*userpb.GetUserByLoginResponse, error) {
	user, err := h.UserUsecase.GetUserByLogin(r.Login)
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserByLoginResponse{
		UserId: user.ID,
		Login: user.Login,
		Username: user.Username,
		Password: user.Password,
	}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, r *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	// TO DO:
	// - Check permissions

	userID := r.UserId
	newUser := &domain.User{
		ID: r.User.UserId,
		Username: r.User.Username,
		Login: r.User.Login,
		Password: r.User.Password,
	}
	respUser, err := h.UserUsecase.UpdateUser(userID, newUser, r.UpdateMask)
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			UserId: respUser.ID,
			Login: respUser.Login,
			Username: respUser.Username,
			Password: respUser.Password,
		},
	}, nil
}