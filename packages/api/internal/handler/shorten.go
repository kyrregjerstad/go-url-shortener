package handler

import (
	"fmt"
	"net/http"
	"url-shortener/internal/validation"
)

func (h *Handler) ShortenUrl(w http.ResponseWriter, r *http.Request) {

	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	validatedURL, err := validation.ValidateAndNormalizeURL(longURL)
	if err != nil {
		switch err {
		case validation.ErrInvalidURL:
			http.Error(w, "Invalid URL", http.StatusBadRequest)
		case validation.ErrURLNotReachable:
			http.Error(w, "URL not reachable", http.StatusBadRequest)
		default:
			http.Error(w, "Failed to validate URL", http.StatusInternalServerError)
		}
		return
	}

	shortCode, err := h.generateShortCode()
	if err != nil {
		http.Error(w, "Failed to generate short code", http.StatusInternalServerError)
		return
	}

	if err := h.store.CreateURL(shortCode, validatedURL); err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Shortened URL: http://localhost:8080/%s", shortCode)
}
