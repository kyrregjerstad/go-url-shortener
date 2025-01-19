package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/internal/handler"
	"url-shortener/internal/storage"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := storage.NewDatabase("urls.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	handler := handler.NewHandler(db)

	http.HandleFunc("/", handler.RedirectURL)
	http.HandleFunc("/shorten", handler.ShortenUrl)
	http.HandleFunc("/stats/", handler.GetStats)

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
