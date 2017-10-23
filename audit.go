package models

import (
	"github.com/jinzhu/gorm"
	"math"
	"strconv"
	"time"
)

type AuditItem struct {
	gorm.Model
	Occurred
	NewDeviceTime                        time.Time
	BGReading                            int
	LinkedBGMeterID                      string `gorm:"size:256"`
	TempBasalAmount                      float64
	TempBasalType                        string `gorm:"size:128"`
	TempBasalDuration                    int
	BolusType                            string `gorm:"size:128"`
	BolusVolumeSelected                  float64
	BolusVolumeDelivered                 float64
	BolusDuration                        int
	PrimeType                            string `gorm:"size:128"`
	PrimeVolumeDelivered                 float64
	Suspend                              string `gorm:"size:128"`
	Rewind                               string `gorm:"size:128"`
	BolusWizardInsulinEstimate           float64
	BolusWizardTargetHighBG              int
	BolusWizardTargetLowBG               int
	BolusWizardCurrentCarbRatio          int
	BolusWizardCurrentInsulinSensitivity int
	BolusWizardCarbInput                 int
	BolusWizardBGInput                   int
	BolusWizardCorrectionEstimate        float64
	BolusWizardFoodEstimate              float64
	BolusWizardActiveInsulin             float64
	Alarm                                string `gorm:"size:256"`
	SensorCalibrationBG                  int
	SensorGluclose                       int
	ISIGValue                            float64
	DailyInsulinTotal                    float64
	RawType                              string `gorm:"size:512"`
	RawValues                            string `gorm:"size:5096"`
	RawDeviceType                        string `gorm:"size:1024"`
	RawSequenceNumber                    int    `gorm:"unique_index"`
}

func (item *AuditItem) SetRaw(key, value string) {
	switch key {
	case "Timestamp":
		parsedTime, err := time.Parse("1/02/06 15:04:05", value)
		if err != nil {
			return
		}
		item.OccurredAt = parsedTime
	case "New Device Time":
		parsedTime, err := time.Parse("1/02/06 15:04:05", value)
		if err != nil {
			return
		}
		item.NewDeviceTime = parsedTime
	case "BG Reading (mg/dL)":
		reading, err := strconv.Atoi(value)
		if err != nil {
			return
		}
		item.BGReading = reading
	case "Linked BG Meter ID":
		item.LinkedBGMeterID = value
	case "Temp Basal Amount (U/h)":
		amount, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		item.TempBasalAmount = amount
	case "Temp Basal Type":
		amount, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		item.TempBasalAmount = amount
	case "Temp Basal Duration (hh:mm:ss)":
		duration, err := time.ParseDuration(value)
		if err != nil {
			return
		}
		item.TempBasalDuration = int(math.Ceil(duration.Minutes()))
	case "Bolus Type":
		item.BolusType = value
	case "Bolus Volume Selected (U)":
		amount, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		item.BolusVolumeSelected = amount
	case "Bolus Volume Delivered (U)":
		amount, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		item.BolusVolumeDelivered = amount
	case "Bolus Duration (hh:mm:ss)":
		duration, err := time.ParseDuration(value)
		if err != nil {
			return
		}
		item.BolusDuration = int(math.Ceil(duration.Minutes()))
	case "Prime Type":
		item.PrimeType = value
	case "Prime Volume Delivered (U)":
		amount, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		item.PrimeVolumeDelivered = amount
	case "Suspend":
		item.Suspend = value
	case "Rewind":
		item.Rewind = value
	case "BWZ Estimate (U)":
		amount, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		item.BolusWizardInsulinEstimate = amount
	case "BWZ Target High BG (mg/dL)":
		amount, err := strconv.Atoi(value)
		if err != nil {
			return
		}
		item.BolusWizardTargetHighBG = amount
	case "BWZ Target Low BG (mg/dL)":
		amount, err := strconv.Atoi(value)
		if err != nil {
			return
		}
		item.BolusWizardTargetLowBG = amount
	case "BWZ Carb Ratio (grams)":
		amount, err := strconv.Atoi(value)
		if err != nil {
			return
		}
		item.BolusWizardCurrentCarbRatio = amount
	case "BWZ Insulin Sensitivity (mg/dL)":
		amount, err := strconv.Atoi(value)
		if err != nil {
			return
		}
		item.BolusWizardCurrentInsulinSensitivity = amount
	case "BWZ Carb Input (grams)":
		amount, err := strconv.Atoi(value)
		if err != nil {
			return
		}
		item.BolusWizardCarbInput = amount
	case "BWZ BG Input (mg/dL)":
		amount, err := strconv.Atoi(value)
		if err != nil {
			return
		}
		item.BolusWizardBGInput = amount
	case "BWZ Correction Estimate (U)":
		amount, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		item.BolusWizardCorrectionEstimate = amount
	case "BWZ Food Estimate (U)":
		amount, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		item.BolusWizardFoodEstimate = amount
	case "BWZ Active Insulin (U)":
		amount, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		item.BolusWizardActiveInsulin = amount
	case "Alarm":
		item.Alarm = value
	case "Sensor Calibration BG (mg/dL)":
		amount, err := strconv.Atoi(value)
		if err != nil {
			return
		}
		item.SensorCalibrationBG = amount
	case "Sensor Glucose (mg/dL)":
		amount, err := strconv.Atoi(value)
		if err != nil {
			return
		}
		item.SensorGluclose = amount
	case "ISIG Value":
		amount, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		item.ISIGValue = amount
	case "Daily Insulin Total (U)":
		amount, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		item.DailyInsulinTotal = amount
	case "Raw-Type":
		item.RawType = value
	case "Raw-Values":
		item.RawValues = value
	case "Raw-Device Type":
		item.RawDeviceType = value
	case "Raw-Seq Num":
		num, err := strconv.Atoi(value)
		if err != nil {
			return
		}
		item.RawSequenceNumber = num
	}
}
