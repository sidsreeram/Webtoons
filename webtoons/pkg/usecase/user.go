package usecase

import (
	"github.com/webtoons/pkg/api/security"
	"github.com/webtoons/pkg/domain"
	repointerfaces "github.com/webtoons/pkg/repository/interface"
	services "github.com/webtoons/pkg/usecase/interfaces"
)

type AuthUseCase struct {
	UserRepo repointerfaces.UserRepository
}

func NewUserUsecase(userrepo repointerfaces.UserRepository) services.UserUsecase {
	return &AuthUseCase{UserRepo: userrepo}
}

func (uc *AuthUseCase) RegisterUser(user domain.User) error {
	return uc.UserRepo.RegisterUser(user)
}

func (uc *AuthUseCase) LoginUser(username, password string) (string, error) {
	valid, err := uc.UserRepo.AuthenticateUser(username, password)
	if err != nil || !valid {
		return "", err
	}

	return security.GenerateJWT(username)
}
