package repository

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/webtoons/pkg/domain"
	repointerfaces "github.com/webtoons/pkg/repository/interface"
)

type UserRepository struct {
	DB *gorm.DB
}

// Change return type to UserRepository interface
func NewUserRepositoryPostgres(db *gorm.DB) repointerfaces.UserRepository {
	return &UserRepository{DB: db}
}

// Save new user with hashed password
func (r *UserRepository) RegisterUser(user domain.User) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Save the user using GORM's Create method
	if err := r.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// Authenticate user by checking password
func (r *UserRepository) AuthenticateUser(username, password string) (bool, error) {
	var user domain.User

	// Find the user by username using GORM's Where method
	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("user not found")
		}
		return false, err
	}

	// Compare the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, errors.New("invalid credentials")
	}

	return true, nil
}
// func (r *UserRepository) Delete(id int) error {
// 	// Use GORM's Delete method to remove the user
// 	if err := r.DB.Delete(&domain.User{}, id).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
