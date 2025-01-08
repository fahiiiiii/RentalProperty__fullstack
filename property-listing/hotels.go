package models

import (
	"time"
)

// Hotel represents a hotel in the database
type Hotel struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Name          string    `gorm:"not null" json:"name"`
	LocationID    uint      `json:"location_id"`
	ReviewScore   string    `json:"review_score"`
	PropertyClass float64   `json:"property_class"`
	IsPreferred   bool      `json:"is_preferred"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}