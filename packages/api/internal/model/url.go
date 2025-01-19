package model

import (
	"database/sql"
	"time"
)

type URLData struct {
	LongURL   string
	ShortCode string
	CreatedAt time.Time
	Visits    int
	LastVisit sql.NullTime
}
