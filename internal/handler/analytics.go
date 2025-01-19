// internal/handler/analytics.go
package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetAnalytics(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[len("/analytics/"):]

	if shortURL == "" {
		http.Error(w, "Short URL is required", http.StatusBadRequest)
		return
	}

	visits, err := h.store.GetURLAnalytics(shortURL)
	if err != nil {
		http.Error(w, "Error fetching analytics", http.StatusInternalServerError)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"short_code": shortURL,
		"visits":     visits,
	})
}
