package storage

import (
	"database/sql"
	"log"
	"time"
	"url-shortener/internal/model"

	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	CreateURL(shortCode, longURL string) error
	GetAndIncrementURL(shortCode string) (string, error)
	GetURLStats(shortCode string) (model.URLData, error)
	RecordVisit(shortCode string, visit model.VisitData) error
	GetURLAnalytics(shortCode string) ([]model.VisitData, error)
	Close() error
}

// SQLiteStore implements the Storage interface using SQLite
type SQLiteStore struct {
	db *sql.DB
}

func NewStorage(dbPath string) (Storage, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Create tables if they don't exist
	queries := []string{
		`CREATE TABLE IF NOT EXISTS urls (
					short_code TEXT PRIMARY KEY,
					long_url TEXT NOT NULL,
					created_at DATETIME NOT NULL,
					visits INTEGER DEFAULT 0,
					last_visit DATETIME
			)`,
		`CREATE TABLE IF NOT EXISTS visits (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					short_code TEXT NOT NULL,
					timestamp DATETIME NOT NULL,
					ip_address TEXT,
					user_agent TEXT,
					referer TEXT,
					FOREIGN KEY(short_code) REFERENCES urls(short_code)
			)`,
	}

	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			log.Printf("Error creating table: %v", err)
			db.Close()
			return nil, err
		}
	}

	return &SQLiteStore{db}, nil
}

func (s *SQLiteStore) CreateURL(shortCode, longURL string) error {
	_, err := s.db.Exec(`
		INSERT INTO urls (short_code, long_url, created_at) 
		VALUES (?, ?, ?)
	`, shortCode, longURL, time.Now())
	return err
}

func (s *SQLiteStore) GetAndIncrementURL(shortCode string) (string, error) {
	var longURL string
	err := s.db.QueryRow(`
		UPDATE urls 
		SET visits = visits + 1, 
		last_visit = ? 
		WHERE short_code = ? 
		RETURNING long_url
	`, time.Now(), shortCode).Scan(&longURL)
	return longURL, err
}

func (s *SQLiteStore) GetURLStats(shortCode string) (model.URLData, error) {
	var url model.URLData
	err := s.db.QueryRow(`
		SELECT long_url, created_at, visits, last_visit 
		FROM urls 
		WHERE short_code = ?
	`, shortCode).Scan(&url.LongURL, &url.CreatedAt, &url.Visits, &url.LastVisit)
	return url, err
}

func (s *SQLiteStore) RecordVisit(shortCode string, visit model.VisitData) error {
	log.Printf("Recording visit for shortCode: %s", shortCode)
	_, err := s.db.Exec(`
					INSERT INTO visits (
									short_code, 
									timestamp, 
									ip_address, 
									user_agent, 
									referer
					) VALUES (?, ?, ?, ?, ?)
	`, shortCode, visit.Timestamp, visit.IPAddress, visit.UserAgent, visit.Referer)
	if err != nil {
		log.Printf("Error recording visit: %v", err)
	}
	return err
}
func (s *SQLiteStore) GetURLAnalytics(shortCode string) ([]model.VisitData, error) {
	log.Printf("Getting analytics for shortCode: %s", shortCode)
	rows, err := s.db.Query(`
					SELECT id, timestamp, ip_address, user_agent, referer
					FROM visits
					WHERE short_code = ?
					ORDER BY timestamp DESC
	`, shortCode)
	if err != nil {
		log.Printf("Error querying analytics: %v", err)
		return nil, err
	}
	defer rows.Close()

	var visits []model.VisitData
	for rows.Next() {
		var visit model.VisitData
		visit.ShortCode = shortCode
		err := rows.Scan(
			&visit.ID,
			&visit.Timestamp,
			&visit.IPAddress,
			&visit.UserAgent,
			&visit.Referer,
		)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		visits = append(visits, visit)
	}
	log.Printf("Found %d visits", len(visits))
	return visits, nil
}

func (s *SQLiteStore) Close() error {
	return s.db.Close()
}
