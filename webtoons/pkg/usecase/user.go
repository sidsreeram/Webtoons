package usecase

import (
    "github.com/webtoons/pkg/domain"
    "github.com/webtoons/pkg/repository"
    "github.com/webtoons/pkg/api/security"
    services "github.com/webtoons/pkg/usecase/interfaces"
)

type AuthUseCase struct {
    UserRepo repository.UserRepository
}
func NewUserUsecase(userrepo repository.UserRepository) services.UserUsecase{
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
