package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() error {
	var err error
	db, err = sql.Open("sqlite3", "urls.db")
	if err != nil {
		return err
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS urls (
		short_code TEXT PRIMARY KEY,
		long_url TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		visits INTEGER DEFAULT 0,
		last_visit DATETIME
	)
	`
	_, err = db.Exec(createTable)
	return err
}

type URLData struct {
	LongURL   string
	CreatedAt time.Time
	Visits    int
	LastVisit time.Time
}

var urls = make(map[string]URLData)

func generateShortCode() (string, error) {

	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	shortCode := base64.URLEncoding.EncodeToString(b)
	return shortCode[:8], nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]

	if shortURL == "" {
		fmt.Fprintf(w, "Welcome to the URL Shortener!")
		return
	}

	var longURL string
	err := db.QueryRow(`
		UPDATE urls 
		SET visits = visits + 1,
				last_visit = ?
		WHERE short_code = ?
		RETURNING long_url`,
		time.Now(), shortURL).Scan(&longURL)

	if err == sql.ErrNoRows {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Println("Database error:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func shortenerHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	// validate URL
	_, err := url.ParseRequestURI(longURL)
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	shortURL, err := generateShortCode()
	if err != nil {
		http.Error(w, "Failed to generate short URL", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec(`
		INSERT INTO urls (short_code, long_url, created_at) 
		VALUES (?, ?, ?)
	`, shortURL, longURL, time.Now())

	if err != nil {
		http.Error(w, "Failed to save URL to database", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Shortened URL: http://localhost:8080/%s", shortURL)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[len("/stats/"):]

	if shortURL == "" {
		http.Error(w, "Short URL is required", http.StatusBadRequest)
		return
	}

	// Query the database for stats
	var longURL string
	var createdAt time.Time
	var visitCount int
	var lastVisited sql.NullTime

	err := db.QueryRow(`
			SELECT long_url, created_at, visits, last_visit 
			FROM urls 
			WHERE short_code = ?`,
		shortURL).Scan(&longURL, &createdAt, &visitCount, &lastVisited)

	if err == sql.ErrNoRows {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Format and display stats
	fmt.Fprintf(w, "Stats for %s:\n", shortURL)
	fmt.Fprintf(w, "Original URL: %s\n", longURL)
	fmt.Fprintf(w, "Created: %v\n", createdAt)
	fmt.Fprintf(w, "Visits: %d\n", visitCount)
	if lastVisited.Valid {
		fmt.Fprintf(w, "Last Visited: %v\n", lastVisited.Time)
	}
}

func main() {
	if err := initDB(); err != nil {
		fmt.Println("Failed to initialize database:", err)
		return
	}
	defer db.Close()

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/shorten", shortenerHandler)
	http.HandleFunc("/stats/", statsHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
