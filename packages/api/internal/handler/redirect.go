package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"url-shortener/internal/model"
)

func (h *Handler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]

	if shortURL == "" {
		log.Printf("Redirecting to home page")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(WelcomeResponse{
			Message: "Welcome to URL shortener!",
		})
		return
	}

	longURL, err := h.store.GetAndIncrementURL(shortURL)
	if err == sql.ErrNoRows {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "URL not found",
		})
		return
	}
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "Database error",
		})
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
