package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) ShortenUrl(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortCode, err := h.generateShortCode()
	if err != nil {
		http.Error(w, "Failed to generate short code", http.StatusInternalServerError)
		return
	}

	if err := h.db.CreateURL(shortCode, longURL); err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Shortened URL: http://localhost:8080/%s", shortCode)
}
