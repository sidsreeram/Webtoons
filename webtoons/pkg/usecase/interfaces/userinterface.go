package interfaces

import "github.com/webtoons/pkg/domain"

type UserUsecase interface {
	RegisterUser(user domain.User) error
	LoginUser(username, password string) (string, error)
}
