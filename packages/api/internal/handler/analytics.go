package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

func (h *Handler) GetAnalytics(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[len("/analytics/"):]

	if shortURL == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "Short URL is required",
		})
		return
	}

	visits, err := h.store.GetURLAnalytics(shortURL)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "Error fetching analytics",
		})
		return
	}

	// Convert visits to response format
	responseVisits := make([]Visit, len(visits))
	for i, v := range visits {
		var referer *string
		if v.Referer != "" {
			referer = &v.Referer
		}

		responseVisits[i] = Visit{
			Timestamp: v.Timestamp.Format(time.RFC3339),
			UserAgent: v.UserAgent,
			IPAddress: v.IPAddress,
			Referer:   referer,
		}
	}

	response := AnalyticsResponse{
		ShortCode: shortURL,
		Visits:    responseVisits,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
