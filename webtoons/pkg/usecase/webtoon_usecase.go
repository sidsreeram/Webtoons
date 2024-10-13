package usecase

import (
	"errors"

	"github.com/webtoons/pkg/domain"
	repointerfaces "github.com/webtoons/pkg/repository/interface"
)

// WebtoonUseCaseImpl implements WebtoonUseCase interface.
type WebtoonUseCase struct {
	WebtoonRepo repointerfaces.WebtoonRepository
}

// NewWebtoonUseCase creates a new WebtoonUseCase.
func NewWebtoonUseCase(repo repointerfaces.WebtoonRepository) *WebtoonUseCase {
	return &WebtoonUseCase{WebtoonRepo: repo}
}

// GetAllWebtoons fetches all webtoons from the repository.
func (uc *WebtoonUseCase) GetAllWebtoons() ([]domain.Webtoon, error) {
	webtoons, err := uc.WebtoonRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return webtoons, nil
}

// AddWebtoon adds a new webtoon to the repository.
func (uc *WebtoonUseCase) AddWebtoon(webtoon domain.Webtoon) error {
	if webtoon.Title == "" || webtoon.Description == "" || len(webtoon.Characters) == 0 {
		return errors.New("webtoon data is incomplete")
	}
	return uc.WebtoonRepo.Save(webtoon)
}

// GetWebtoonByID fetches a webtoon by its ID.
func (uc *WebtoonUseCase) GetWebtoonByID(id int) (domain.Webtoon, error) {
	webtoon, err := uc.WebtoonRepo.GetByID(id)
	if err != nil {
		return domain.Webtoon{}, err
	}
	return webtoon, nil
}

// DeleteWebtoon removes a webtoon by its ID.
func (uc *WebtoonUseCase) DeleteWebtoon(id int) error {
	return uc.WebtoonRepo.Delete(id)
}
