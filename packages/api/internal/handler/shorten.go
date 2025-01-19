package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener/internal/validation"
)

func (h *Handler) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	type request struct {
		URL string `json:"url"`
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	validatedURL, err := validation.ValidateAndNormalizeURL(req.URL)
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"shortUrl": fmt.Sprintf("http://localhost:8080/%s", shortCode),
	})
}
