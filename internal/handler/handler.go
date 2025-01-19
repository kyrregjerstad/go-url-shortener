package handler

import (
	"crypto/rand"
	"encoding/base64"
	"url-shortener/internal/storage"
)

type Handler struct {
	db *storage.Store
}

func NewHandler(db *storage.Store) *Handler {
	return &Handler{db: db}
}

func (h *Handler) generateShortCode() (string, error) {

	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	shortCode := base64.URLEncoding.EncodeToString(b)
	return shortCode[:8], nil
}
