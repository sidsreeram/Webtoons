package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/webtoons/pkg/domain"
)

type WebtoonRepository struct {
    DB *sql.DB
}

func NewWebtoonRepository(db *sql.DB) *WebtoonRepository {
    return &WebtoonRepository{DB: db}
}

func (r *WebtoonRepository) GetAll() ([]domain.Webtoon, error) {
    var allWebtoons []domain.Webtoon
    query := `SELECT id, title, description, characters FROM webtoons`

    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var webtoon domain.Webtoon
        var characters []byte // to scan JSONB characters
        if err := rows.Scan(&webtoon.ID, &webtoon.Title, &webtoon.Description, &characters); err != nil {
            return nil, err
        }

        // Convert characters JSONB to []string
        if err := json.Unmarshal(characters, &webtoon.Characters); err != nil {
            return nil, err
        }

        allWebtoons = append(allWebtoons, webtoon)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return allWebtoons, nil
}


func (r *WebtoonRepository) GetByID(id int) (domain.Webtoon, error) {
    var webtoon domain.Webtoon
    var characters []byte // to scan JSONB characters

    query := `SELECT id, title, description, characters FROM webtoons WHERE id = $1`
    row := r.DB.QueryRow(query, id)

    err := row.Scan(&webtoon.ID, &webtoon.Title, &webtoon.Description, &characters)
    if err != nil {
        if err == sql.ErrNoRows {
            return webtoon, nil // return empty result if no row found
        }
        return webtoon, err
    }

    // Convert characters JSONB to []string
    if err := json.Unmarshal(characters, &webtoon.Characters); err != nil {
        return webtoon, err
    }

    return webtoon, nil
}

func (r *WebtoonRepository) Save(webtoon domain.Webtoon) error {
    query := `INSERT INTO webtoons (title, description, characters) VALUES ($1, $2, $3)`

    charactersJSON, err := json.Marshal(webtoon.Characters) // Convert []string to JSONB
    if err != nil {
        return err
    }

    _, err = r.DB.Exec(query, webtoon.Title, webtoon.Description, charactersJSON)
    if err != nil {
        return err
    }

    return nil
}


func (r *WebtoonRepository) Delete(id int) error {
    query := `DELETE FROM webtoons WHERE id = $1`

    _, err := r.DB.Exec(query, id)
    if err != nil {
        return err
    }

    return nil
}
