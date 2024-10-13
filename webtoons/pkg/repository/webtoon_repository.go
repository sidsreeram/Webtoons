package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/webtoons/pkg/domain"
	repointerfaces "github.com/webtoons/pkg/repository/interface"
)

type WebtoonRepository struct {
	DB *gorm.DB
}

func NewWebtoonRepository(db *gorm.DB) repointerfaces.WebtoonRepository {
	return &WebtoonRepository{db}
}

// GetAll retrieves all webtoons from the database
func (r *WebtoonRepository) GetAll() ([]domain.Webtoon, error) {
	var allWebtoons []domain.Webtoon

	// Use GORM's Find method to fetch all records
	if err := r.DB.Find(&allWebtoons).Error; err != nil {
		return nil, err
	}

	// No need to manually unmarshal JSON, GORM handles JSONB fields
	return allWebtoons, nil
}

// GetByID retrieves a specific webtoon by its ID
func (r *WebtoonRepository) GetByID(id int) (domain.Webtoon, error) {
	var webtoon domain.Webtoon

	// Use GORM's First method to find the webtoon by ID
	if err := r.DB.First(&webtoon, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return webtoon, nil // Return empty result if no row found
		}
		return webtoon, err
	}

	return webtoon, nil
}

// Save adds a new webtoon to the database
func (r *WebtoonRepository) Save(webtoon domain.Webtoon) error {
	// GORM automatically handles JSONB when saving structs
	if err := r.DB.Create(&webtoon).Error; err != nil {
		return err
	}

	return nil
}

// Delete removes a webtoon by ID from the database
func (r *WebtoonRepository) Delete(id int) error {
	// Use GORM's Delete method to remove the record
	if err := r.DB.Delete(&domain.Webtoon{}, id).Error; err != nil {
		return err
	}

	return nil
}
