package models

import "gorm.io/gorm"

type WeatherLog struct {
	gorm.Model
	City        string
	Temperature float64
	UpdatedAt   int64
}
