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

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// During development, allow all origins
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	store, err := storage.NewStorage("urls.db")
	if err != nil {
		log.Fatalf("Failed to connect to storage: %v", err)
	}
	defer store.Close()

	handler := handler.NewHandler(store)
	router := http.NewServeMux()

	router.HandleFunc("GET /", handler.RedirectURL)
	router.HandleFunc("POST /shorten", handler.ShortenUrl)
	router.HandleFunc("GET /stats/", handler.GetStats)
	router.HandleFunc("GET /analytics/", handler.GetAnalytics)

	server := &http.Server{
		Addr:    ":8080",
		Handler: corsMiddleware(router),
	}

	fmt.Println("Server is running on port 8080...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
