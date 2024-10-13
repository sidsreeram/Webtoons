package repointerfaces
import "github.com/webtoons/pkg/domain"

type WebtoonRepository interface {
	GetAll() ([]domain.Webtoon, error);
	GetByID(id int) (domain.Webtoon, error)
	Save(webtoon domain.Webtoon) error
	Delete(id int) error
}
type UserRepository interface {
	RegisterUser(user domain.User) error
	AuthenticateUser(username, password string) (bool, error) 
	// Delete(id int) error
}
