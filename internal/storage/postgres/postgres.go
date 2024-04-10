package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Prrromanssss/URLShortener/internal/storage"
	"github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(storageURL string) (*Storage, error) {
	const op = "storage.postgres.New"
	db, err := sql.Open("postgres", storageURL)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Storage{db: db}, nil
}

func (s *Storage) SaveURL(urlToSave, alias string) (int64, error) {
	const op = "storage.postgres.SaveURL"

	stmt, err := s.db.Prepare("INSERT INTO url (url, alias) VALUES ($1, $2) RETURNING id;")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var id int64
	err = stmt.QueryRow(urlToSave, alias).Scan(&id)
	if err != nil {
		pgErr, ok := err.(*pq.Error)
		if !ok {
			return 0, fmt.Errorf("%s: %w", op, err)
		}

		if pgErr.Code == "23505" {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrURLExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
