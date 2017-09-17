package models

import (
	"github.com/jinzhu/gorm"
)

type SugarSource string

const (
	SourceMeter = "meter"
	SourceCGM   = "cgm"
)

// Sugar represents a measured blood sugar
type Sugar struct {
	gorm.Model
	Occurred
	Value  int
	Source SugarSource `gorm:"size:64"`
}
