package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener/internal/validation"
)

func (h *Handler) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "Invalid request body",
		})
		return
	}

	if req.URL == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "URL is required",
		})
		return
	}

	validatedURL, err := validation.ValidateAndNormalizeURL(req.URL)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		var message string
		switch err {
		case validation.ErrInvalidURL:
			message = "Invalid URL"
		case validation.ErrURLNotReachable:
			message = "URL not reachable"
		default:
			message = "Failed to validate URL"
		}
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: message,
		})
		return
	}

	shortCode, err := h.generateShortCode()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "Failed to generate short code",
		})
		return
	}

	if err := h.store.CreateURL(shortCode, validatedURL); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "Failed to save URL",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ShortenResponse{
		ShortURL: fmt.Sprintf("http://localhost:8080/%s", shortCode),
	})
}
