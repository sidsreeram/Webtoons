package repository

import (
	"database/sql"
	"errors"

	"github.com/webtoons/pkg/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepositoryPostgres(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Save new user with hashed password
func (r *UserRepository) RegisterUser(user domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (username, password) VALUES ($1, $2)`
	_, err = r.DB.Exec(query, user.Username, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

// Authenticate user by checking password
func (r *UserRepository) AuthenticateUser(username, password string) (bool, error) {
	var storedPassword string
	query := `SELECT password FROM users WHERE username = $1`
	row := r.DB.QueryRow(query, username)
	err := row.Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, errors.New("user not found")
		}
		return false, err
	}

	// Compare the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		return false, errors.New("invalid credentials")
	}

	return true, nil
}
