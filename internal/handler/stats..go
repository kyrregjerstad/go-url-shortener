package handler

import (
	"database/sql"
	"fmt"
	"net/http"
)

func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[len("/stats/"):]

	if shortURL == "" {
		http.Error(w, "Short URL is required", http.StatusBadRequest)
		return
	}

	data, err := h.store.GetURLStats(shortURL)
	if err == sql.ErrNoRows {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Stats for %s:\n", shortURL)
	fmt.Fprintf(w, "Original URL: %s\n", data.LongURL)
	fmt.Fprintf(w, "Created: %v\n", data.CreatedAt)
	fmt.Fprintf(w, "Visits: %d\n", data.Visits)
	if data.LastVisit.Valid {
		fmt.Fprintf(w, "Last Visited: %v\n", data.LastVisit.Time)
	}
}
