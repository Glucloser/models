package models

import "github.com/jinzhu/gorm"

// Place represents a place a meal was eaten
type Place struct {
	gorm.Model
	Name       string `gorm:"size:4096"`
	SearchName string `gorm:"size:1024"`
}

// PlaceVisit represents an visit to a place
type PlaceVisit struct {
	gorm.Model
	Occurred
	PlaceID uint
}
