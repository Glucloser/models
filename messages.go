package models

import (
	"github.com/jinzhu/gorm"
)

// Message stores messages for later study
type Message struct {
	gorm.Model
	Message string `gorm:"size:4096"`
}
