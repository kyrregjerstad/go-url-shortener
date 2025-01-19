package handler

import (
	"database/sql"
	"fmt"
	"net/http"
)

func (h *Handler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]

	if shortURL == "" {
		fmt.Fprintf(w, "Welcome to URL shortener!")
		return
	}

	var longURL string

	longURL, err := h.db.GetAndIncrementURL(shortURL)

	if err == sql.ErrNoRows {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
