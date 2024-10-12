package interfaces
import "github.com/webtoons/pkg/domain"
type WebtoonUseCase interface{
	GetAllWebtoons() ([]domain.Webtoon, error)
	AddWebtoon(webtoon domain.Webtoon) error
	GetWebtoonByID(id int) (domain.Webtoon, error)
	DeleteWebtoon(id int) error
}
