package interfaces
import "github.com/webtoons/pkg/domain"

type WebtoonRepository interface {
	GetAll() ([]domain.Webtoon, error);
	GetByID(id int) (domain.Webtoon, error)
	Save(webtoon domain.Webtoon) error
	Delete(id int) error
}
