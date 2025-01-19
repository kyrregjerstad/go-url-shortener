package handler

// Common response types
type ErrorResponse struct {
	Error string `json:"error"`
}

type WelcomeResponse struct {
	Message string `json:"message"`
}

// Shorten endpoint types
type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"shortUrl"`
}

// Stats endpoint types
type StatsResponse struct {
	LongUrl   string  `json:"longUrl"`
	CreatedAt string  `json:"createdAt"`
	Visits    int     `json:"visits"`
	LastVisit *string `json:"lastVisit,omitempty"`
}

// Analytics endpoint types
type Visit struct {
	Timestamp string  `json:"timestamp"`
	UserAgent string  `json:"userAgent"`
	IPAddress string  `json:"ipAddress"`
	Referer   *string `json:"referer,omitempty"`
}

type AnalyticsResponse struct {
	ShortCode string  `json:"short_code"`
	Visits    []Visit `json:"visits"`
}
