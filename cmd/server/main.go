package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"url-shortener/internal/handler"
	"url-shortener/internal/storage"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	store, err := storage.NewStorage("urls.db")
	if err != nil {
		log.Fatalf("Failed to connect to storage: %v", err)
	}
	defer store.Close()

	handler := handler.NewHandler(store)

	http.HandleFunc("/", handler.RedirectURL)
	http.HandleFunc("/shorten", handler.ShortenUrl)
	http.HandleFunc("/stats/", handler.GetStats)
	http.HandleFunc("/analytics/", handler.GetAnalytics)

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
