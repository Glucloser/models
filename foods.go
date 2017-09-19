package models

import (
	"github.com/jinzhu/gorm"
)

// Food represents a food eaten
type Food struct {
	gorm.Model
	Occurred
	Name    string `gorm:"size:2048"`
	Carbs   int
	Insulin float64
}
