package models

import (
	"github.com/jinzhu/gorm"
)

// User represents a user
type User struct {
	gorm.Model
	Name                string `gorm:"size:1024"`
	FoursquareID        string `gorm:"size:1024"`
	FacebookMessengerID string `gorm:"size:1024"`
	AlexaID             string `gorm:"size:2048"`
	GoogleID            string `gorm:"size:2048"`
}
