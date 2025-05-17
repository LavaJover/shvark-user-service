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

func (uc *UserUsecase) CheckPermission(userID string, required domain.Role) (bool, error){
	// Look for user with given ID
	user, err := uc.Repo.GetUserByID(userID)
	if err != nil{
		return false, domain.ErrUserNotFound
	}

	// Compare users role and required parameter
	if user.Role != required {
		return false, nil
	}

	return true, nil
}

// Updates user mathing given userID
func (uc *UserUsecase) UpdateUser(user *domain.User) error {
	// Look for user with given ID
	oldUser, err := uc.Repo.GetUserByID(user.ID)
	if err != nil {
		return err
	}

	// Update the fields
	oldUser.Login = user.Login
	oldUser.Username = user.Username
	oldUser.Password = user.Password

	return nil
}