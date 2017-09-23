package models

import (
	"time"
)

// Occured represents the time when an event happend.
type Occurred struct {
	OccurredAt time.Time `gorm:"unique_index"`
}
