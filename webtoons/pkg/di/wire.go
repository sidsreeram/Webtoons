package di

import (
	"github.com/google/wire"
	http "github.com/webtoons/pkg/api"
	"github.com/webtoons/pkg/api/delivery"
	"github.com/webtoons/pkg/config"
	"github.com/webtoons/pkg/db"
	"github.com/webtoons/pkg/repository"
	"github.com/webtoons/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabase,
		// Repositories
		repository.NewUserRepositoryPostgres,
		repository.NewWebtoonRepository,
		// Use cases
		usecase.NewWebtoonUseCase,
		usecase.NewUserUsecase,
		// Handlers (AuthHandler and WebtoonHandler)
		delivery.NewAuthHandler,    // AuthHandler Provider
		delivery.NewWebtoonHandler, // WebtoonHandler Provider
		// Server HTTP
		http.NewServerHTTP,
	)
	return &http.ServerHTTP{}, nil
}
