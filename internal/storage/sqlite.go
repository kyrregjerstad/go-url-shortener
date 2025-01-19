package storage

import (
	"database/sql"
	"time"
	"url-shortener/internal/model"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sql.DB
}

func NewDatabase(dbPath string) (*Store, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Create table if it doesn't exist
	createTable := `
	CREATE TABLE IF NOT EXISTS urls (
			short_code TEXT PRIMARY KEY,
			long_url TEXT NOT NULL,
			created_at DATETIME NOT NULL,
			visits INTEGER DEFAULT 0,
			last_visit DATETIME
	)`

	if _, err := db.Exec(createTable); err != nil {
		db.Close()
		return nil, err
	}

	return &Store{db}, nil
}

func (db *Store) CreateURL(shortCode, longURL string) error {
	_, err := db.db.Exec(`
		INSERT INTO urls (short_code, long_url, created_at) 
		VALUES (?, ?, ?)
	`, shortCode, longURL, time.Now())
	return err
}

func (db *Store) GetAndIncrementURL(shortCode string) (string, error) {
	var longURL string
	err := db.db.QueryRow(`
		UPDATE urls 
		SET visits = visits + 1, 
		last_visit = ? 
		WHERE short_code = ? 
		RETURNING long_url
	`, time.Now(), shortCode).Scan(&longURL)
	return longURL, err
}

func (db *Store) GetURLStats(shortCode string) (model.URLData, error) {
	var url model.URLData
	err := db.db.QueryRow(`
		SELECT long_url, created_at, visits, last_visit 
		FROM urls 
		WHERE short_code = ?
	`, shortCode).Scan(&url.LongURL, &url.CreatedAt, &url.Visits, &url.LastVisit)
	return url, err
}

func (s *Store) Close() error {
	return s.db.Close()
}
