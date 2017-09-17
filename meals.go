package models

import (
	"github.com/jinzhu/gorm"
)

// Meal represents a set of foods eaten at a place
type Meal struct {
	gorm.Model
	Occurred
	Place Place
	Foods []Food
}
