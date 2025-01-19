package model

import (
	"time"
)

type VisitData struct {
	ID        int64
	ShortCode string
	Timestamp time.Time
	IPAddress string
	UserAgent string
	Referer   string
}
