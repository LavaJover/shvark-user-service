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
		TwoFaSecret: user.TwoFaSecret,
		TwoFaEnabled: user.TwoFaEnabled,
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
		TwoFaSecret: user.TwoFaSecret,
		TwoFaEnabled: user.TwoFaEnabled,
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
		TwoFaSecret: r.User.TwoFaSecret,
		TwoFaEnabled: r.User.TwoFaEnabled,
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
			TwoFaSecret: respUser.TwoFaSecret,
			TwoFaEnabled: respUser.TwoFaEnabled,
		},
	}, nil
}

func (h *UserHandler) GetUsers(ctx context.Context, r *userpb.GetUsersRequest) (*userpb.GetUsersResponse, error) {
	page, limit := r.Page, r.Limit
	userRecords, totalPages, err := h.UserUsecase.GetUsers(page, limit)
	if err != nil {
		return nil, err
	}

	var users []*userpb.User
	for _, userRecord := range userRecords {
		users = append(users, &userpb.User{
			UserId: userRecord.ID,
			Login: userRecord.Login,
			TwoFaSecret: userRecord.TwoFaSecret,
			Username: userRecord.Username,
			Password: userRecord.Password,
			TwoFaEnabled: userRecord.TwoFaEnabled,
		})
	}

	return &userpb.GetUsersResponse{
		TotalPages: int32(totalPages),
		Users: users,
	}, nil
}

func (h *UserHandler) SetTwoFaSecret(ctx context.Context, r *userpb.SetTwoFaSecretRequest) (*userpb.SetTwoFaSecretResponse, error) {
	userID, twoFaSecret := r.UserId, r.TwoFaSecret
	err := h.UserUsecase.SetTwoFaSecret(userID, twoFaSecret)
	if err != nil {
		return nil, err
	}

	return &userpb.SetTwoFaSecretResponse{}, nil
}

func (h *UserHandler) GetTwoFaSecretByID(ctx context.Context, r *userpb.GetTwoFaSecretByIDRequest) (*userpb.GetTwoFaSecretByIDResponse, error) {
	twoFaSecret, err := h.UserUsecase.GetTwoFaSecretByID(r.UserId)
	if err != nil {
		return nil, err
	}

	return &userpb.GetTwoFaSecretByIDResponse{
		TwoFaSecret: twoFaSecret,
	}, nil
}