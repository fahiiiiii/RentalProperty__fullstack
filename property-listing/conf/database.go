package conf

import (
    "fmt"
    "log"
    "property-listing/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    // Hard-coded database connection parameters
    DB_HOST := "localhost" // Change this to your DB host if needed
    DB_USER := "fahimah"
    DB_PASSWORD := "fahimah123"
    DB_NAME := "rental_db"
    DB_PORT := "5432"

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        DB_HOST,
        DB_USER,
        DB_PASSWORD,
        DB_NAME,
        DB_PORT,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Auto-migrate the Location, RentalProperty, and PropertyDetails tables
    err = db.AutoMigrate(&models.Location{}, &models.RentalProperty{}, &models.PropertyDetails{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    DB = db
}
