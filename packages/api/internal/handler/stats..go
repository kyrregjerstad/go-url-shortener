package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[len("/stats/"):]

	if shortURL == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "Short URL is required",
		})
		return
	}

	data, err := h.store.GetURLStats(shortURL)
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

	var lastVisit *string
	if data.LastVisit.Valid {
		lastVisitStr := data.LastVisit.Time.Format(time.RFC3339)
		lastVisit = &lastVisitStr
	}

	response := StatsResponse{
		LongUrl:   data.LongURL,
		CreatedAt: data.CreatedAt.Format(time.RFC3339),
		Visits:    data.Visits,
		LastVisit: lastVisit,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
