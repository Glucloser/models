package models

import (
	"github.com/jinzhu/gorm"
)

// Food represents a food eaten
type Food struct {
	gorm.Model
	Occurred
	Name      string `gorm:"size:2048"`
	FoodStems []FoodStem
	Carbs     int
	Insulin   float64
}

// FoodStem represents the stem of a food name
type FoodStem struct {
	gorm.Model
	Stem   string `gorm:"size:2048"`
	FoodID uint
}
