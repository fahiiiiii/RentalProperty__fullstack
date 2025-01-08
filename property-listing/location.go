package models

import (
	"time"
)

// Location represents a location and its hotels in the database
type Location struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CityName    string    `gorm:"not null" json:"city_name"`
	Country     string    `gorm:"not null" json:"country"`
	Hotels      []Hotel   `gorm:"foreignKey:LocationID" json:"hotels"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}