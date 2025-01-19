package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"url-shortener/internal/model"
)

func (h *Handler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]

	if shortURL == "" {
		log.Printf("Redirecting to home page")
		fmt.Fprintf(w, "Welcome to URL shortener!")
		return
	}

	longURL, err := h.store.GetAndIncrementURL(shortURL)
	if err == sql.ErrNoRows {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	visit := model.VisitData{
		ShortCode: shortURL,
		Timestamp: time.Now(),
		IPAddress: r.RemoteAddr,
		UserAgent: r.UserAgent(),
		Referer:   r.Referer(),
	}

	// Don't block the redirect if analytics fails
	if err := h.store.RecordVisit(shortURL, visit); err != nil {
		log.Printf("Error recording visit: %v", err)
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
