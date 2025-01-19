package validation

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	ErrInvalidURL      = errors.New("invalid URL")
	ErrURLNotReachable = errors.New("URL not reachable")
)

func ValidateAndNormalizeURL(rawURL string) (string, error) {
	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		rawURL = "https://" + rawURL
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", ErrInvalidURL
	}

	// Ensure scheme and host are present
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return "", ErrInvalidURL
	}

	// Check if URL is reachable with a timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Head(parsedURL.String())
	if err != nil {
		return "", ErrURLNotReachable
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", ErrURLNotReachable
	}

	return parsedURL.String(), nil
}
